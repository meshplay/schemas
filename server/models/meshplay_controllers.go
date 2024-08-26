package models

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/khulnasoft/meshkit/broker/nats"
	"github.com/khulnasoft/meshkit/database"
	"github.com/khulnasoft/meshkit/logger"
	"github.com/khulnasoft/meshkit/models/controllers"
	"github.com/khulnasoft/meshkit/utils"
	meshplaykube "github.com/khulnasoft/meshkit/utils/kubernetes"
	"github.com/spf13/viper"
)

const (
	ChartRepo                     = "https://meshplay.github.io/khulnasoft.com/charts"
	MeshplayServerBrokerConnection = "meshplay-server"
)

type MeshplayController int

const (
	MeshplayBroker MeshplayController = iota
	Meshsync
	MeshplayOperator
)

type MeshplayControllersHelper struct {
	//  maps each context with the controller handlers
	// this map will be used as the source of truth
	ctxControllerHandlersMap map[string]map[MeshplayController]controllers.IMeshplayController
	// maps each context with it's operator status
	ctxOperatorStatusMap map[string]controllers.MeshplayControllerStatus
	// maps each context with a meshsync data handler
	ctxMeshsyncDataHandlerMap map[string]MeshsyncDataHandler

	mu                sync.Mutex

	log          logger.Handler
	oprDepConfig controllers.OperatorDeploymentConfig
	dbHandler    *database.Handler
}

func (mch *MeshplayControllersHelper) GetControllerHandlersForEachContext() map[string]map[MeshplayController]controllers.IMeshplayController {
	return mch.ctxControllerHandlersMap
}

func (mch *MeshplayControllersHelper) GetMeshSyncDataHandlersForEachContext() map[string]MeshsyncDataHandler {
	return mch.ctxMeshsyncDataHandlerMap
}

func (mch *MeshplayControllersHelper) GetOperatorsStatusMap() map[string]controllers.MeshplayControllerStatus {
	return mch.ctxOperatorStatusMap
}

func NewMeshplayControllersHelper(log logger.Handler, operatorDepConfig controllers.OperatorDeploymentConfig, dbHandler *database.Handler) *MeshplayControllersHelper {
	return &MeshplayControllersHelper{
		ctxControllerHandlersMap:  make(map[string]map[MeshplayController]controllers.IMeshplayController),
		log:                       log,
		oprDepConfig:              operatorDepConfig,
		ctxOperatorStatusMap:      make(map[string]controllers.MeshplayControllerStatus),
		ctxMeshsyncDataHandlerMap: make(map[string]MeshsyncDataHandler),
		dbHandler:                 dbHandler,
	}
}

// initializes Meshsync data handler for the contexts for whom it has not been
// initialized yet. Apart from updating the map, it also runs the handler after
// updating the map. The presence of a handler for a context in a map indicate that
// the meshsync data for that context is properly being handled
func (mch *MeshplayControllersHelper) UpdateMeshsynDataHandlers(ctx context.Context, connectionID, userID, meshplayInstanceID uuid.UUID, provider Provider) *MeshplayControllersHelper {
	// only checking those contexts whose MeshplayConrollers are active
	go func(mch *MeshplayControllersHelper) {
		mch.mu.Lock()
		defer mch.mu.Unlock()
		for ctxID, controllerHandlers := range mch.ctxControllerHandlersMap {
			if _, ok := mch.ctxMeshsyncDataHandlerMap[ctxID]; !ok {
				// brokerStatus := controllerHandlers[MeshplayBroker].GetStatus()
				// do something if broker is being deployed , maybe try again after sometime
				brokerEndpoint, err := controllerHandlers[MeshplayBroker].GetPublicEndpoint()
				if brokerEndpoint == "" {
					if err != nil {
						mch.log.Warn(err)
					}
					mch.log.Info(fmt.Sprintf("Meshplay Broker unreachable for Kubernetes context (%v)", ctxID))
					continue
				}
				mch.log.Info(fmt.Sprintf("Connected to Meshplay Broker (%s) for Kubernetes context (%s)", brokerEndpoint, ctxID))
				brokerHandler, err := nats.New(nats.Options{
					// URLS: []string{"localhost:4222"},
					URLS:           []string{brokerEndpoint},
					ConnectionName: MeshplayServerBrokerConnection,
					Username:       "",
					Password:       "",
					ReconnectWait:  2 * time.Second,
					MaxReconnect:   60,
				})
				if err != nil {
					mch.log.Warn(err)
					mch.log.Info(fmt.Sprintf("MeshSync not configured for Kubernetes context (%v) due to '%v'", ctxID, err.Error()))
					continue
				}
				mch.log.Info(fmt.Sprintf("Connected to Meshplay Broker (%v) for Kubernetes context (%v)", brokerEndpoint, ctxID))
				token, _ := ctx.Value(TokenCtxKey).(string)
				msDataHandler := NewMeshsyncDataHandler(brokerHandler, *mch.dbHandler, mch.log, provider, userID, connectionID, meshplayInstanceID, token)
				err = msDataHandler.Run()
				if err != nil {
					mch.log.Warn(err)
					mch.log.Info(fmt.Sprintf("Unable to connect MeshSync for Kubernetes context (%s) due to: %s", ctxID, err.Error()))
					continue
				}
				mch.ctxMeshsyncDataHandlerMap[ctxID] = *msDataHandler
				mch.log.Info(fmt.Sprintf("MeshSync connected for Kubernetes context (%s)", ctxID))
			}
		}
	}(mch)

	return mch
}

// attach a MeshplayController for each context if
// 1. the config is valid
// 2. if it is not already attached
func (mch *MeshplayControllersHelper) UpdateCtxControllerHandlers(ctxs []K8sContext) *MeshplayControllersHelper {
	go func(mch *MeshplayControllersHelper) {
		mch.mu.Lock()
		defer mch.mu.Unlock()
		
		// resetting this value as a specific controller handler instance does not have any significance opposed to
		// a MeshsyncDataHandler instance where it signifies whether or not a listener is attached
		mch.ctxControllerHandlersMap = make(map[string]map[MeshplayController]controllers.IMeshplayController)
		for _, ctx := range ctxs {
			
			ctxID := ctx.ID
			cfg, _ := ctx.GenerateKubeConfig()
			client, err := meshplaykube.New(cfg)
			// means that the config is invalid
			if err != nil {
				
				// invalid configs are not added to the map
				continue
			}
			mch.ctxControllerHandlersMap[ctxID] = map[MeshplayController]controllers.IMeshplayController{
				MeshplayBroker:   controllers.NewMeshplayBrokerHandler(client),
				MeshplayOperator: controllers.NewMeshplayOperatorHandler(client, mch.oprDepConfig),
				Meshsync:        controllers.NewMeshsyncHandler(client),
			}
		}
	}(mch)
	return mch
}

// update the status of MeshplayOperator in all the contexts
// for whom MeshplayControllers are attached
// should be called after UpdateCtxControllerHandlers
func (mch *MeshplayControllersHelper) UpdateOperatorsStatusMap(ot *OperatorTracker) *MeshplayControllersHelper {
	go func(mch *MeshplayControllersHelper) {
		mch.mu.Lock()
		defer mch.mu.Unlock()
		mch.ctxOperatorStatusMap = make(map[string]controllers.MeshplayControllerStatus)
		for ctxID, ctrlHandler := range mch.ctxControllerHandlersMap {
			if ot.IsUndeployed(ctxID) {
				mch.ctxOperatorStatusMap[ctxID] = controllers.Undeployed
			} else {
				mch.ctxOperatorStatusMap[ctxID] = ctrlHandler[MeshplayOperator].GetStatus()
			}
		}
	}(mch)
	return mch
}

type OperatorTracker struct {
	ctxIDtoDeploymentStatus map[string]bool
	mx                      sync.Mutex
	DisableOperator         bool
}

func NewOperatorTracker(disabled bool) *OperatorTracker {
	return &OperatorTracker{
		ctxIDtoDeploymentStatus: make(map[string]bool),
		mx:                      sync.Mutex{},
		DisableOperator:         disabled,
	}
}

func (ot *OperatorTracker) Undeployed(ctxID string, undeployed bool) {
	if ot.DisableOperator { //no-op when operator is disabled
		return
	}
	ot.mx.Lock()
	defer ot.mx.Unlock()
	if ot.ctxIDtoDeploymentStatus == nil {
		ot.ctxIDtoDeploymentStatus = make(map[string]bool)
	}
	ot.ctxIDtoDeploymentStatus[ctxID] = undeployed
}
func (ot *OperatorTracker) IsUndeployed(ctxID string) bool {
	if ot.DisableOperator { //Return true everytime so that operators stay in undeployed state across all contexts
		return true
	}
	ot.mx.Lock()
	defer ot.mx.Unlock()
	if ot.ctxIDtoDeploymentStatus == nil {
		ot.ctxIDtoDeploymentStatus = make(map[string]bool)
		return false
	}
	return ot.ctxIDtoDeploymentStatus[ctxID]
}

// looks at the status of Meshplay Operator for each cluster and takes necessary action.
// it will deploy the operator only when it is in NotDeployed state
func (mch *MeshplayControllersHelper) DeployUndeployedOperators(ot *OperatorTracker) *MeshplayControllersHelper {
	if ot.DisableOperator { //Return true everytime so that operators stay in undeployed state across all contexts
		return mch
	}
	go func(mch *MeshplayControllersHelper) {
		mch.mu.Lock()
		defer mch.mu.Unlock()
		for ctxID, ctrlHandler := range mch.ctxControllerHandlersMap {
			if oprStatus, ok := mch.ctxOperatorStatusMap[ctxID]; ok {
				
				if oprStatus == controllers.NotDeployed {
					
					err := ctrlHandler[MeshplayOperator].Deploy(false)
					if err != nil {
						mch.log.Error(err)
					}
				}
			}
		}
	}(mch)

	return mch
}

func (mch *MeshplayControllersHelper) UndeployDeployedOperators(ot *OperatorTracker) *MeshplayControllersHelper {
	go func(mch *MeshplayControllersHelper) {
		
		mch.mu.Lock()
		defer mch.mu.Unlock()
		for ctxID, ctrlHandler := range mch.ctxControllerHandlersMap {
					
			if oprStatus, ok := mch.ctxOperatorStatusMap[ctxID]; ok {
				
				if oprStatus != controllers.Undeployed {
					
					err := ctrlHandler[MeshplayOperator].Undeploy()
					
					if err != nil {
						mch.log.Error(err)
					}
				}
			}
		}
	}(mch)
	return mch
}

func NewOperatorDeploymentConfig(adapterTracker AdaptersTrackerInterface) controllers.OperatorDeploymentConfig {
	// get meshplay release version
	meshplayReleaseVersion := viper.GetString("BUILD")
	if meshplayReleaseVersion == "" || meshplayReleaseVersion == "Not Set" || meshplayReleaseVersion == "edge-latest" {
		_, latestRelease, err := CheckLatestVersion(meshplayReleaseVersion)
		// if unable to fetch latest release tag, meshkit helm functions handle
		// this automatically fetch the latest one
		if err != nil {
			// mch.log.Error(fmt.Errorf("Couldn't check release tag: %s. Will use latest version", err))
			meshplayReleaseVersion = ""
		} else {
			meshplayReleaseVersion = latestRelease
		}
	}

	return controllers.OperatorDeploymentConfig{
		MeshplayReleaseVersion: meshplayReleaseVersion,
		GetHelmOverrides: func(delete bool) map[string]interface{} {
			return setOverrideValues(delete, adapterTracker)
		},
		HelmChartRepo: ChartRepo,
	}
}

// checkLatestVersion takes in the current server version compares it with the target
// and returns the (isOutdated, latestVersion, error)
func CheckLatestVersion(serverVersion string) (*bool, string, error) {
	// Inform user of the latest release version
	versions, err := utils.GetLatestReleaseTagsSorted("meshplay", "meshplay")
	latestVersion := versions[len(versions)-1]
	isOutdated := false
	if err != nil {
		return nil, "", ErrCreateOperatorDeploymentConfig(err)
	}
	// Compare current running Meshplay server version to the latest available Meshplay release on GitHub.
	if latestVersion != serverVersion {
		isOutdated = true
		return &isOutdated, latestVersion, nil
	}

	return &isOutdated, latestVersion, nil
}

// setOverrideValues detects the currently insalled adapters and sets appropriate
// overrides so as to not uninstall them. It also sets override values for
// operator so that it can be enabled or disabled depending on the need
func setOverrideValues(delete bool, adapterTracker AdaptersTrackerInterface) map[string]interface{} {
	installedAdapters := make([]string, 0)
	adapters := adapterTracker.GetAdapters(context.TODO())

	for _, adapter := range adapters {
		if adapter.Name != "" {
			installedAdapters = append(installedAdapters, strings.Split(adapter.Location, ":")[0])
		}
	}

	overrideValues := map[string]interface{}{
		"fullnameOverride": "meshplay-operator",
		"meshplay": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-istio": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-cilium": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-linkerd": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-consul": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-kuma": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-nsm": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-nginx-sm": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-traefik-mesh": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-app-mesh": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-operator": map[string]interface{}{
			"enabled": true,
		},
	}

	for _, adapter := range installedAdapters {
		if _, ok := overrideValues[adapter]; ok {
			overrideValues[adapter] = map[string]interface{}{
				"enabled": true,
			}
		}
	}

	if delete {
		overrideValues["meshplay-operator"] = map[string]interface{}{
			"enabled": false,
		}
	}

	return overrideValues
}

// setOverrideValues detects the currently insalled adapters and sets appropriate
// overrides so as to not uninstall them.
func SetOverrideValuesForMeshplayDeploy(adapters []Adapter, adapter Adapter, install bool) map[string]interface{} {
	installedAdapters := make([]string, 0)
	for _, adapter := range adapters {
		if adapter.Name != "" {
			installedAdapters = append(installedAdapters, strings.Split(adapter.Location, ":")[0])
		}
	}

	overrideValues := map[string]interface{}{
		"meshplay-istio": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-cilium": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-linkerd": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-consul": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-kuma": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-nsm": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-nginx-sm": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-traefik-mesh": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-app-mesh": map[string]interface{}{
			"enabled": false,
		},
	}

	for _, adapter := range installedAdapters {
		if _, ok := overrideValues[adapter]; ok {
			overrideValues[adapter] = map[string]interface{}{
				"enabled": true,
			}
		}
	}

	// based on deploy/undeploy action change the status of adapter override
	if _, ok := overrideValues[strings.Split(adapter.Location, ":")[0]]; ok {
		overrideValues[strings.Split(adapter.Location, ":")[0]] = map[string]interface{}{
			"enabled": install,
		}
	}

	return overrideValues
}
