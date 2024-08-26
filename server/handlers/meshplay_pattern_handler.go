package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gofrs/uuid"
	guid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/khulnasoft/meshplay/server/meshes"
	"github.com/khulnasoft/meshplay/server/models"
	pCore "github.com/khulnasoft/meshplay/server/models/pattern/core"
	"github.com/khulnasoft/meshplay/server/models/pattern/stages"
	"github.com/khulnasoft/meshkit/errors"
	"github.com/khulnasoft/meshkit/models/events"
	meshmodel "github.com/khulnasoft/meshkit/models/meshmodel/registry"
	"github.com/khulnasoft/meshkit/utils/kubernetes"
	"github.com/khulnasoft/meshkit/utils/kubernetes/kompose"
	"github.com/khulnasoft/meshkit/utils/walker"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// MeshplayPatternRequestBody refers to the type of request body that
// SaveMeshplayPattern would receive
type MeshplayPatternRequestBody struct {
	Name          string                 `json:"name,omitempty"`
	URL           string                 `json:"url,omitempty"`
	Path          string                 `json:"path,omitempty"`
	Save          bool                   `json:"save,omitempty"`
	PatternData   *models.MeshplayPattern `json:"pattern_data,omitempty"`
	CytoscapeJSON string                 `json:"cytoscape_json,omitempty"`
}

// PatternFileRequestHandler will handle requests of both type GET and POST
// on the route /api/pattern
func (h *Handler) PatternFileRequestHandler(
	rw http.ResponseWriter,
	r *http.Request,
	prefObj *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	if r.Method == http.MethodGet {
		h.GetMeshplayPatternsHandler(rw, r, prefObj, user, provider)
		return
	}

	if r.Method == http.MethodPost {
		h.handlePatternPOST(rw, r, prefObj, user, provider)
		return
	}
}

// swagger:route POST /api/pattern PatternsAPI idPostPatternFile
// Handle POST requests for patterns
//
// Edit/update a meshplay pattern
// responses:
// 	200: meshplayPatternResponseWrapper

func (h *Handler) handlePatternPOST(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	var err error
	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("create").ActedUpon(userID).WithSeverity(events.Informational)

	res := meshes.EventsResponse{
		Component:     "core",
		ComponentName: "Design",
		OperationId:   uuid.Nil.String(), // to be removed
		EventType:     meshes.EventType_INFO,
	}
	sourcetype := mux.Vars(r)["sourcetype"]
	var parsedBody *MeshplayPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrRequestBody(err),
		}).WithDescription("Unable to parse uploaded pattern.").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		return
	}

	actedUpon := &userID
	if parsedBody.PatternData != nil && parsedBody.PatternData.ID != nil {
		actedUpon = parsedBody.PatternData.ID
	}

	eventBuilder.ActedUpon(*actedUpon)

	token, err := provider.GetProviderToken(r)
	if err != nil {
		h.log.Error(ErrRetrieveUserToken(err))
		http.Error(rw, ErrRetrieveUserToken(err).Error(), http.StatusInternalServerError)
		event := eventBuilder.WithSeverity(events.Critical).WithMetadata(map[string]interface{}{
			"error": ErrRetrieveUserToken(err),
		}).WithDescription("No auth token provided in the request.").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)

		return
	}

	format := r.URL.Query().Get("output")
	var meshplayPattern *models.MeshplayPattern

	if parsedBody.CytoscapeJSON != "" {
		pf, err := pCore.NewPatternFileFromCytoscapeJSJSON(parsedBody.Name, []byte(parsedBody.CytoscapeJSON))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(rw, "%s", err)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrSavePattern(err),
			}).WithDescription("Pattern save failed, cytoJSON could be malformed.").Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}

		pfByt, err := pf.ToYAML()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, "%s", err)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrSavePattern(err),
			}).WithDescription(ErrSavePattern(err).Error()).Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}

		patternName, err := models.GetPatternName(string(pfByt))
		if err != nil {
			h.log.Error(ErrSavePattern(err))
			http.Error(rw, ErrSavePattern(err).Error(), http.StatusBadRequest)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrSavePattern(err),
			}).WithDescription("unable to get \"name\" from the pattern.").Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}

		meshplayPattern := &models.MeshplayPattern{
			Name:        patternName,
			PatternFile: string(pfByt),
			Location: map[string]interface{}{
				"host": "",
				"path": "",
				"type": "local",
			},
			CatalogData: parsedBody.PatternData.CatalogData,
			Type: sql.NullString{
				String: string(models.Design),
				Valid:  true,
			},
		}
		if parsedBody.PatternData != nil {
			meshplayPattern.ID = parsedBody.PatternData.ID
		}

		if parsedBody.Save {
			resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
			if err != nil {
				h.log.Error(ErrSavePattern(err))
				http.Error(rw, ErrSavePattern(err).Error(), http.StatusInternalServerError)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrSavePattern(err),
				}).WithDescription(ErrSavePattern(err).Error()).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}

			h.formatPatternOutput(rw, resp, format, &res, eventBuilder)
			event := eventBuilder.Build()
			_ = provider.PersistEvent(event)
			// Do not send pattern save event if pattern is in cyto format as user is on meshmap and every node move will result in save request flooding user's screen.
			// go h.config.EventBroadcaster.Publish(userID, event)
			go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
			return
		}

		byt, err := json.Marshal([]models.MeshplayPattern{*meshplayPattern})
		if err != nil {
			h.log.Error(ErrEncodePattern(err))
			http.Error(rw, ErrEncodePattern(err).Error(), http.StatusInternalServerError)
			addMeshkitErr(&res, ErrEncodePattern(err))
			go h.EventsBuffer.Publish(&res)
			return
		}

		h.formatPatternOutput(rw, byt, format, &res, eventBuilder)

		return
	}
	// If Content is not empty then assume it's a local upload
	if parsedBody.PatternData != nil {
		// Assign a location if no location is specified
		if parsedBody.PatternData.Location == nil {
			parsedBody.PatternData.Location = map[string]interface{}{
				"host":   "",
				"path":   "",
				"type":   "local",
				"branch": "",
			}
		}

		meshplayPattern = parsedBody.PatternData
		bytPattern := []byte(meshplayPattern.PatternFile)
		meshplayPattern.SourceContent = bytPattern
		if sourcetype == string(models.DockerCompose) || sourcetype == string(models.K8sManifest) {
			var k8sres string
			if sourcetype == string(models.DockerCompose) {
				k8sres, err = kompose.Convert(bytPattern) // convert the docker compose file into kubernetes manifest
				if err != nil {
					obj := "convert"
					conversionErr := ErrApplicationFailure(err, obj)
					h.log.Error(conversionErr)

					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": conversionErr,
					}).WithDescription(fmt.Sprintf("Failed converting Docker Compose application %s", meshplayPattern.Name)).Build()

					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)

					http.Error(rw, conversionErr.Error(), http.StatusInternalServerError)
					addMeshkitErr(&res, ErrApplicationFailure(err, obj))
					go h.EventsBuffer.Publish(&res)
					return
				}
				meshplayPattern.Type = sql.NullString{
					String: string(models.DockerCompose),
					Valid:  true,
				}
			} else if sourcetype == string(models.K8sManifest) {
				k8sres = string(bytPattern)
				meshplayPattern.Type = sql.NullString{
					String: string(models.K8sManifest),
					Valid:  true,
				}
			}
			pattern, err := pCore.NewPatternFileFromK8sManifest(k8sres, false, h.registryManager)
			if err != nil {
				obj := "convert"
				conversionErr := ErrApplicationFailure(err, obj)
				h.log.Error(conversionErr)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": conversionErr,
				}).WithDescription(fmt.Sprintf("Failed converting K8s Manifest %s to design file format.", meshplayPattern.Name)).Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				http.Error(rw, conversionErr.Error(), http.StatusInternalServerError)
				addMeshkitErr(&res, err) //this error is already a meshkit error so no further wrapping required
				go h.EventsBuffer.Publish(&res)
				return
			}
			response, err := yaml.Marshal(pattern)
			if err != nil {
				obj := "convert"
				conversionErr := ErrApplicationFailure(err, obj)
				h.log.Error(conversionErr)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": conversionErr,
				}).WithDescription(fmt.Sprintf("Failed converting design %s to YAML format.", meshplayPattern.Name)).Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				http.Error(rw, conversionErr.Error(), http.StatusInternalServerError)
				addMeshkitErr(&res, ErrApplicationFailure(err, obj))
				go h.EventsBuffer.Publish(&res)
				return
			}
			meshplayPattern.PatternFile = string(response)
		} else {
			parsedBody.PatternData.Type = sql.NullString{
				String: string(models.Design),
				Valid:  true,
			}
			// Check if the pattern is valid
			err := pCore.IsValidPattern(parsedBody.PatternData.PatternFile)
			if err != nil {
				h.log.Error(ErrInvalidPattern(err))
				http.Error(rw, ErrInvalidPattern(err).Error(), http.StatusBadRequest)
				addMeshkitErr(&res, ErrInvalidPattern(err))
				go h.EventsBuffer.Publish(&res)
				return
			}
			// Assign a name if no name is provided
			if parsedBody.PatternData.Name == "" {
				patternName, err := models.GetPatternName(parsedBody.PatternData.PatternFile)
				if err != nil {
					h.log.Error(ErrSavePattern(err))
					http.Error(rw, ErrSavePattern(err).Error(), http.StatusBadRequest)
					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": ErrSavePattern(err),
					}).WithDescription("unable to get \"name\" from the pattern.").Build()

					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)
					return
				}
				parsedBody.PatternData.Name = patternName
			}

			meshplayPattern := parsedBody.PatternData

			if parsedBody.Save {
				resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
				if err != nil {
					h.log.Error(ErrSavePattern(err))
					http.Error(rw, ErrSavePattern(err).Error(), http.StatusInternalServerError)

					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": ErrSavePattern(err),
					}).WithDescription(ErrSavePattern(err).Error()).Build()

					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)
					return
				}

				h.formatPatternOutput(rw, resp, format, &res, eventBuilder)
				event := eventBuilder.Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
				return
			}

			byt, err := json.Marshal([]models.MeshplayPattern{*meshplayPattern})
			if err != nil {
				h.log.Error(ErrEncodePattern(err))
				http.Error(rw, ErrEncodePattern(err).Error(), http.StatusInternalServerError)
				addMeshkitErr(&res, ErrSavePattern(err))
				go h.EventsBuffer.Publish(&res)
				return
			}

			h.formatPatternOutput(rw, byt, format, &res, eventBuilder)
			event := eventBuilder.Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}
	}

	if parsedBody.URL != "" {
		if sourcetype == string(models.HelmChart) {
			helmSourceResp, err := http.Get(parsedBody.URL)
			defer func() {
				_ = helmSourceResp.Body.Close()
			}()
			if err != nil {
				obj := "import"
				importErr := ErrApplicationFailure(err, obj)
				h.log.Error(importErr)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": importErr,
				}).WithDescription(fmt.Sprintf("Failed importing design from URL %s.", parsedBody.URL)).Build()
				_ = provider.PersistEvent(event)

				go h.config.EventBroadcaster.Publish(userID, event)
				http.Error(rw, importErr.Error(), http.StatusInternalServerError)
				addMeshkitErr(&res, ErrApplicationFailure(err, obj))
				go h.EventsBuffer.Publish(&res)
				return
			}
			sourceContent, err := io.ReadAll(helmSourceResp.Body)
			if err != nil {
				http.Error(rw, "error read body", http.StatusInternalServerError)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrSaveApplication(err),
				}).WithDescription(fmt.Sprintf("error reading design from the remote URL %s, URL is malformed or not reachable.", parsedBody.URL)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				addMeshkitErr(&res, ErrSaveApplication(fmt.Errorf("error reading body")))
				go h.EventsBuffer.Publish(&res)
				return
			}

			resp, err := kubernetes.ConvertHelmChartToK8sManifest(kubernetes.ApplyHelmChartConfig{
				URL: parsedBody.URL,
			})
			if err != nil {
				obj := "import"
				importErr := ErrApplicationFailure(err, obj)
				h.log.Error(importErr)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": importErr,
				}).WithDescription(fmt.Sprintf("error converting helm chart %s to Kubernetes manifest, URL might be malformed or not reachable.", parsedBody.URL)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				http.Error(rw, importErr.Error(), http.StatusInternalServerError)
				addMeshkitErr(&res, ErrApplicationFailure(err, obj))
				go h.EventsBuffer.Publish(&res)
				return
			}

			result := string(resp)
			pattern, err := pCore.NewPatternFileFromK8sManifest(result, false, h.registryManager)
			if err != nil {
				obj := "convert"
				convertErr := ErrApplicationFailure(err, obj)
				h.log.Error(convertErr)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": convertErr,
				}).WithDescription(fmt.Sprintf("Failed converting Helm Chart %s to design file format", parsedBody.URL)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				addMeshkitErr(&res, err)
				go h.EventsBuffer.Publish(&res)
				return
			}

			response, err := yaml.Marshal(pattern)
			if err != nil {
				obj := "convert"
				convertErr := ErrApplicationFailure(err, obj)
				h.log.Error(convertErr)
				http.Error(rw, convertErr.Error(), http.StatusInternalServerError) // sending a 500 when we cannot convert the file into kuberentes manifest

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": convertErr,
				}).WithDescription(fmt.Sprintf("Failed converting Helm Chart %s to design file format", parsedBody.URL)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				addMeshkitErr(&res, ErrApplicationFailure(err, obj))
				go h.EventsBuffer.Publish(&res)
				return
			}

			url := strings.Split(parsedBody.URL, "/")
			meshplayPattern = &models.MeshplayPattern{
				Name:        strings.TrimSuffix(url[len(url)-1], ".tgz"),
				PatternFile: string(response),
				Type: sql.NullString{
					String: string(models.HelmChart),
					Valid:  true,
				},
				Location: map[string]interface{}{
					"type":   "http",
					"host":   parsedBody.URL,
					"path":   "",
					"branch": "",
				},
				SourceContent: sourceContent,
			}
		} else if sourcetype == string(models.DockerCompose) || sourcetype == string(models.K8sManifest) {
			parsedURL, err := url.Parse(parsedBody.URL)
			if err != nil {
				err := ErrSaveApplication(fmt.Errorf("error parsing URL"))
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": err,
				}).WithDescription(fmt.Sprintf("Invalid URL provided %s", parsedBody.URL)).Build()
				addMeshkitErr(&res, ErrSaveApplication(fmt.Errorf("error parsing URL")))
				go h.EventsBuffer.Publish(&res)
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}

			// Check if hostname is github
			if parsedURL.Host == "github.com" {
				parsedPath := strings.Split(parsedURL.Path, "/")
				if parsedPath[3] == "tree" {
					parsedPath = append(parsedPath[0:3], parsedPath[4:]...)
				}
				if len(parsedPath) < 3 {
					http.Error(rw, "malformed URL: url should be of type github.com/<owner>/<repo>/[branch]", http.StatusNotAcceptable)
				}

				owner := parsedPath[1]
				repo := parsedPath[2]
				branch := "master"
				path := parsedBody.Path
				if len(parsedPath) == 4 {
					branch = parsedPath[3]
				}
				if path == "" && len(parsedPath) > 4 {
					path = strings.Join(parsedPath[4:], "/")
				}

				pfs, err := githubRepoDesignScan(owner, repo, path, branch, sourcetype, h.registryManager)
				if err != nil {
					remoteApplicationErr := ErrRemoteApplication(err)
					http.Error(rw, remoteApplicationErr.Error(), http.StatusInternalServerError)

					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": err,
					}).WithDescription(fmt.Sprintf("Failed to retrieve remote design at %s", parsedBody.URL)).Build()

					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)
					addMeshkitErr(&res, err) //error guaranteed to be meshkit error
					go h.EventsBuffer.Publish(&res)
					return
				}

				meshplayPattern = &pfs[0]
			} else {
				// Fallback to generic HTTP import
				pfs, err := genericHTTPDesignFile(parsedBody.URL, sourcetype, h.registryManager)
				if err != nil {
					remoteApplicationErr := ErrRemoteApplication(err)
					http.Error(rw, remoteApplicationErr.Error(), http.StatusInternalServerError)

					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": err,
					}).WithDescription(fmt.Sprintf("Failed to retrieve remote design at %s", parsedBody.URL)).Build()
					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)
					addMeshkitErr(&res, err) //error guaranteed to be meshkit error
					go h.EventsBuffer.Publish(&res)
					return
				}
				meshplayPattern = &pfs[0]
			}
		} else {
			parsedBody.PatternData.Type = sql.NullString{
				String: string(models.Design),
				Valid:  true,
			}
			resp, err := provider.RemotePatternFile(r, parsedBody.URL, parsedBody.Path, parsedBody.Save)

			if err != nil {
				h.log.Error(ErrImportPattern(err))
				http.Error(rw, ErrImportPattern(err).Error(), http.StatusInternalServerError)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrImportPattern(err),
				}).WithDescription(ErrImportPattern(err).Error()).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}

			h.formatPatternOutput(rw, resp, format, &res, eventBuilder)
			event := eventBuilder.Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}
	}

	if sourcetype == string(models.DockerCompose) || sourcetype == string(models.K8sManifest) || sourcetype == string(models.HelmChart) {
		var savedPatternID *uuid.UUID

		if parsedBody.Save {
			resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
			if err != nil {
				obj := "save"

				saveErr := ErrApplicationFailure(err, obj)
				h.log.Error(saveErr)
				http.Error(rw, saveErr.Error(), http.StatusInternalServerError)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": saveErr,
				}).WithDescription(fmt.Sprintf("Failed persisting design %s", parsedBody.Name)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				addMeshkitErr(&res, ErrApplicationFailure(err, obj))
				go h.EventsBuffer.Publish(&res)
				return
			}

			h.formatPatternOutput(rw, resp, format, &res, eventBuilder)

			eventBuilder.WithSeverity(events.Informational)
			event := eventBuilder.Build()
			go h.config.EventBroadcaster.Publish(userID, event)
			_ = provider.PersistEvent(event)

			var meshplayPatternContent []models.MeshplayPattern
			err = json.Unmarshal(resp, &meshplayPatternContent)
			if err != nil {
				obj := "pattern"
				h.log.Error(models.ErrEncoding(err, obj))
				http.Error(rw, models.ErrEncoding(err, obj).Error(), http.StatusInternalServerError)
				return
			}
			savedPatternID = meshplayPatternContent[0].ID
			err = provider.SaveMeshplayPatternSourceContent(token, (savedPatternID).String(), meshplayPattern.SourceContent)

			if err != nil {
				obj := "upload"
				uploadSourceContentErr := ErrApplicationSourceContent(err, obj)

				h.log.Error(uploadSourceContentErr)
				http.Error(rw, uploadSourceContentErr.Error(), http.StatusInternalServerError)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": uploadSourceContentErr,
				}).WithDescription("Failed uploading original design content to remote provider.").Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				addMeshkitErr(&res, ErrApplicationSourceContent(err, obj))
				go h.EventsBuffer.Publish(&res)
				return
			}
			go h.config.ApplicationChannel.Publish(userID, struct{}{})
			eb := eventBuilder
			_ = provider.PersistEvent(eb.WithDescription(fmt.Sprintf("Design %s  Source content uploaded", meshplayPatternContent[0].Name)).Build())
			return
		}

		meshplayPattern.ID = savedPatternID
		byt, err := json.Marshal([]models.MeshplayPattern{*meshplayPattern})
		if err != nil {
			obj := "design"
			h.log.Error(models.ErrEncoding(err, obj))
			http.Error(rw, models.ErrEncoding(err, obj).Error(), http.StatusInternalServerError)
			return
		}

		h.formatPatternOutput(rw, byt, format, &res, eventBuilder)

		event := eventBuilder.Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
	}

}

func githubRepoDesignScan(
	owner,
	repo,
	path,
	branch,
	sourceType string,
	reg *meshmodel.RegistryManager,
) ([]models.MeshplayPattern, error) {
	var mu sync.Mutex
	ghWalker := walker.NewGit()
	result := make([]models.MeshplayPattern, 0)
	err := ghWalker.
		Owner(owner).
		Repo(repo).
		Branch(branch).
		Root(path).
		RegisterFileInterceptor(func(f walker.File) error {
			ext := filepath.Ext(f.Name)
			var k8sres string
			var err error
			k8sres = f.Content
			if ext == ".yml" || ext == ".yaml" {
				if sourceType == string(models.DockerCompose) {
					k8sres, err = kompose.Convert([]byte(f.Content))
					if err != nil {
						return ErrRemoteApplication(err)
					}
				}
				pattern, err := pCore.NewPatternFileFromK8sManifest(k8sres, false, reg)
				if err != nil {
					return err //always a meshkit error
				}
				response, err := yaml.Marshal(pattern)
				if err != nil {
					return models.ErrMarshal(err, string(response))
				}

				af := models.MeshplayPattern{
					Name:        strings.TrimSuffix(f.Name, ext),
					PatternFile: string(response),
					Location: map[string]interface{}{
						"type":   "github",
						"host":   fmt.Sprintf("github.com/%s/%s", owner, repo),
						"path":   f.Path,
						"branch": branch,
					},
					Type: sql.NullString{
						String: string(sourceType),
						Valid:  true,
					},
					SourceContent: []byte(f.Content),
				}

				mu.Lock()
				result = append(result, af)
				mu.Unlock()
			}

			return nil
		}).
		Walk()

	return result, ErrRemoteApplication(err)
}

func genericHTTPDesignFile(fileURL, sourceType string, reg *meshmodel.RegistryManager) ([]models.MeshplayPattern, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, ErrRemoteApplication(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrRemoteApplication(fmt.Errorf("file not found"))
	}

	defer models.SafeClose(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrRemoteApplication(err)
	}

	k8sres := string(body)

	if sourceType == string(models.DockerCompose) {
		k8sres, err = kompose.Convert(body)
		if err != nil {
			return nil, ErrRemoteApplication(err)
		}
	}

	pattern, err := pCore.NewPatternFileFromK8sManifest(k8sres, false, reg)
	if err != nil {
		return nil, err //This error is already a meshkit error
	}
	response, err := yaml.Marshal(pattern)

	if err != nil {
		return nil, models.ErrMarshal(err, string(response))
	}

	url := strings.Split(fileURL, "/")
	af := models.MeshplayPattern{
		Name:        url[len(url)-1],
		PatternFile: string(response),
		Location: map[string]interface{}{
			"type":   "http",
			"host":   fileURL,
			"path":   "",
			"branch": "",
		},
		Type: sql.NullString{
			String: string(sourceType),
			Valid:  true,
		},
		SourceContent: body,
	}
	return []models.MeshplayPattern{af}, nil
}

// swagger:route GET /api/pattern PatternsAPI idGetPatternFiles
// Handle GET request for patterns
//
// Returns the list of all the patterns saved by the current user
// This will return all the patterns with their details
//
// ```?order={field}``` orders on the passed field
//
// ```?search=<design name>``` A string matching is done on the specified design name
//
// ```?page={page-number}``` Default page number is 1
//
// ```?pagesize={pagesize}``` Default pagesize is 10
//
// ```?visibility={[visibility]}``` Default visibility is public + private; Mulitple visibility filters can be passed as an array
// Eg: ```?visibility=["public", "published"]``` will return public and published designs
//
// responses:
//
//	200: meshplayPatternsResponseWrapper
func (h *Handler) GetMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	q := r.URL.Query()
	tokenString := r.Context().Value(models.TokenCtxKey).(string)
	updateAfter := q.Get("updated_after")

	err := r.ParseForm() // necessary to get r.Form["visibility"], i.e, ?visibility=public&visbility=private
	if err != nil {
		h.log.Error(ErrFetchPattern(err))
		http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
		return
	}
	filter := struct {
		Visibility []string `json:"visibility"`
	}{}

	visibility := q.Get("visibility")
	if visibility != "" {
		err := json.Unmarshal([]byte(visibility), &filter.Visibility)
		if err != nil {
			h.log.Error(ErrFetchPattern(err))
			http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
			return
		}
	}

	resp, err := provider.GetMeshplayPatterns(tokenString, q.Get("page"), q.Get("pagesize"), q.Get("search"), q.Get("order"), updateAfter, filter.Visibility)
	if err != nil {
		h.log.Error(ErrFetchPattern(err))
		http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	// token, err := provider.GetProviderToken(r)
	if err != nil {
		http.Error(rw, "failed to get user token", http.StatusInternalServerError)
		return
	}
	// mc := NewContentModifier(token, provider, prefObj, user.UserID)
	// //acts like a middleware, modifying the bytes lazily just before sending them back
	// err = mc.AddMetadataForPatterns(r.Context(), &resp)
	if err != nil {
		fmt.Println("Could not add metadata about pattern's current support ", err.Error())
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/catalog PatternsAPI idGetCatalogMeshplayPatternsHandler
// Handle GET request for catalog patterns
//
// # Patterns can be further filtered through query parameter
//
// ```?order={field}``` orders on the passed field
//
// ```?page={page-number}``` Default page number is 0
//
// ```?pagesize={pagesize}``` Default pagesize is 10.
//
// ```?search={patternname}``` If search is non empty then a greedy search is performed
// responses:
//
//	200: meshplayPatternsResponseWrapper
func (h *Handler) GetCatalogMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	q := r.URL.Query()
	tokenString := r.Context().Value(models.TokenCtxKey).(string)

	resp, err := provider.GetCatalogMeshplayPatterns(tokenString, q.Get("page"), q.Get("pagesize"), q.Get("search"), q.Get("order"))
	if err != nil {
		h.log.Error(ErrFetchPattern(err))
		http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/pattern/{id} PatternsAPI idDeleteMeshplayPattern
// Handle Delete for a Meshplay Pattern
//
// Deletes a meshplay pattern with ID: id
// responses:
//
//	200: noContentWrapper
//
// DeleteMeshplayPatternHandler deletes a pattern with the given id
func (h *Handler) DeleteMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("delete").ActedUpon(uuid.FromStringOrNil(patternID))

	meshplayPattern := models.MeshplayPattern{}

	resp, err := provider.DeleteMeshplayPattern(r, patternID)
	if err != nil {
		errPatternDelete := ErrDeletePattern(err)

		h.log.Error(errPatternDelete)
		http.Error(rw, errPatternDelete.Error(), http.StatusInternalServerError)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": errPatternDelete,
		}).WithDescription("Error deleting pattern.").Build()
		http.Error(rw, errPatternDelete.Error(), http.StatusInternalServerError)
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		return
	}

	_ = json.Unmarshal(resp, &meshplayPattern)
	event := eventBuilder.WithSeverity(events.Informational).WithDescription(fmt.Sprintf("Pattern %s deleted.", meshplayPattern.Name)).Build()
	_ = provider.PersistEvent(event)
	go h.config.EventBroadcaster.Publish(userID, event)
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/{id} PatternsAPI idGetMeshplayPattern
// Handle GET request for Meshplay Pattern with the given id
//
// Get the pattern with the given id
// responses:
//  200:

// GetMeshplayPatternHandler returns the pattern file with the given id

func (h *Handler) DownloadMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	resp, err := provider.GetMeshplayPattern(r, patternID)
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		return
	}

	pattern := &models.MeshplayPattern{}

	err = json.Unmarshal(resp, &pattern)
	if err != nil {
		obj := "download pattern"
		h.log.Error(models.ErrUnmarshal(err, obj))
		http.Error(rw, models.ErrUnmarshal(err, obj).Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "Pattern/x-yaml")
	if _, err := io.Copy(rw, strings.NewReader(pattern.PatternFile)); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/pattern/clone/{id} PatternsAPI idCloneMeshplayPattern
// Handle Clone for a Meshplay Pattern
//
// Creates a local copy of a published pattern with id: id
// responses:
//
//	200 : noContentWrapper
//
// CloneMeshplayPatternHandler clones a pattern with the given id
func (h *Handler) CloneMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	var parsedBody *models.MeshplayClonePatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil || patternID == "" {
		h.log.Error(ErrRequestBody(err))
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}

	resp, err := provider.CloneMeshplayPattern(r, patternID, parsedBody)
	if err != nil {
		h.log.Error(ErrClonePattern(err))
		http.Error(rw, ErrClonePattern(err).Error(), http.StatusInternalServerError)
		return
	}
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route POST /api/pattern/catalog/publish PatternsAPI idPublishCatalogPatternHandler
// Handle Publish for a Meshplay Pattern
//
// Publishes pattern to Meshplay Catalog by setting visibility to published and setting catalog data
// responses:
//
//	202: noContentWrapper
//
// PublishCatalogPatternHandler sets visibility of pattern with given id as published
func (h *Handler) PublishCatalogPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().
		FromUser(userID).
		FromSystem(*h.SystemID).
		WithCategory("pattern").
		WithAction("publish").
		ActedUpon(userID)

	var parsedBody *models.MeshplayCatalogPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		e := eventBuilder.WithSeverity(events.Error).
			WithMetadata(map[string]interface{}{
				"error": ErrRequestBody(err),
			}).
			WithDescription("Error parsing design payload.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}
	resp, err := provider.PublishCatalogPattern(r, parsedBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		e := eventBuilder.WithSeverity(events.Error).
			WithMetadata(map[string]interface{}{
				"error": ErrPublishCatalogPattern(err),
			}).
			WithDescription("Error publishing design.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	var respBody *models.CatalogRequest
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		e := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrPublishCatalogPattern(err),
		}).WithDescription("Error parsing response.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
	}

	e := eventBuilder.WithSeverity(events.Informational).ActedUpon(parsedBody.ID).WithDescription(fmt.Sprintf("Request to publish '%s' design submitted with status: %s", respBody.ContentName, respBody.Status)).Build()
	_ = provider.PersistEvent(e)
	go h.config.EventBroadcaster.Publish(userID, e)

	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/pattern/catalog/unpublish PatternsAPI idUnPublishCatalogPatternHandler
// Handle Publish for a Meshplay Pattern
//
// Unpublishes pattern from Meshplay Catalog by setting visibility to private and removing catalog data from website
// responses:
//
//	200: noContentWrapper
//
// UnPublishCatalogPatternHandler sets visibility of pattern with given id as private
func (h *Handler) UnPublishCatalogPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().
		FromUser(userID).
		FromSystem(*h.SystemID).
		WithCategory("pattern").
		WithAction("unpublish_request").
		ActedUpon(userID)

	var parsedBody *models.MeshplayCatalogPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		e := eventBuilder.WithSeverity(events.Error).
			WithMetadata(map[string]interface{}{
				"error": ErrRequestBody(err),
			}).
			WithDescription("Error parsing design payload.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}
	resp, err := provider.UnPublishCatalogPattern(r, parsedBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		e := eventBuilder.WithSeverity(events.Error).
			WithMetadata(map[string]interface{}{
				"error": ErrPublishCatalogPattern(err),
			}).
			WithDescription("Error publishing design.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	var respBody *models.CatalogRequest
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		e := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrPublishCatalogPattern(err),
		}).WithDescription("Error parsing response.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
	}

	e := eventBuilder.WithSeverity(events.Informational).ActedUpon(parsedBody.ID).WithDescription(fmt.Sprintf("'%s' design unpublished", respBody.ContentName)).Build()
	_ = provider.PersistEvent(e)
	go h.config.EventBroadcaster.Publish(userID, e)

	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/patterns PatternsAPI idDeleteMeshplayPattern
// Handle Delete for multiple Meshplay Patterns
//
// DeleteMultiMeshplayPatternsHandler deletes patterns with the given ids
func (h *Handler) DeleteMultiMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error(rw, "err deleting pattern, converting bytes: ", err)
	}
	var patterns models.MeshplayPatternDeleteRequestBody
	err = json.Unmarshal([]byte(body), &patterns)
	if err != nil {
		logrus.Error("error marshaling patterns json: ", err)
	}

	logrus.Debugf("patterns to be deleted: %+v", patterns)

	resp, err := provider.DeleteMeshplayPatterns(r, patterns)

	if err != nil {
		http.Error(rw, fmt.Sprintf("failed to delete the pattern: %s", err), http.StatusInternalServerError)
		return
	}
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/{id} PatternsAPI idGetMeshplayPattern
// Handle GET for a Meshplay Pattern
//
// Fetches the pattern with the given id
// responses:
// 	200: meshplayPatternResponseWrapper

// GetMeshplayPatternHandler fetched the pattern with the given id
func (h *Handler) GetMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]

	resp, err := provider.GetMeshplayPattern(r, patternID)
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

func (h *Handler) formatPatternOutput(rw http.ResponseWriter, content []byte, format string, res *meshes.EventsResponse, eventBuilder *events.EventBuilder) {
	contentMeshplayPatternSlice := make([]models.MeshplayPattern, 0)

	if err := json.Unmarshal(content, &contentMeshplayPatternSlice); err != nil {
		http.Error(rw, ErrDecodePattern(err).Error(), http.StatusInternalServerError)
		addMeshkitErr(res, ErrDecodePattern(err))
		go h.EventsBuffer.Publish(res)
		return
	}

	result := []models.MeshplayPattern{}
	names := []string{}
	for _, content := range contentMeshplayPatternSlice {
		if content.ID != nil {
			eventBuilder.ActedUpon(*content.ID)
		}
		if format == "cytoscape" {
			patternFile, err := pCore.NewPatternFile([]byte(content.PatternFile))
			if err != nil {
				http.Error(rw, ErrParsePattern(err).Error(), http.StatusBadRequest)

				eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrParsePattern(err),
				}).WithDescription("Unable to parse pattern file, pattern could be malformed.").Build()
				return
			}

			//TODO: The below line has to go away once the client fully supports referencing variables  and pattern imports inside design
			newpatternfile := evalImportAndReferenceStage(&patternFile)

			cyjs, _ := newpatternfile.ToCytoscapeJS()

			bytes, err := json.Marshal(&cyjs)
			if err != nil {
				http.Error(rw, ErrConvertPattern(err).Error(), http.StatusInternalServerError)
				addMeshkitErr(res, ErrConvertPattern(err))
				go h.EventsBuffer.Publish(res)
				return
			}

			// Replace the patternfile with cytoscape type data
			content.PatternFile = string(bytes)
		}

		result = append(result, content)
		names = append(names, content.Name)
	}

	data, err := json.Marshal(&result)
	if err != nil {
		obj := "pattern file"
		http.Error(rw, models.ErrMarshal(err, obj).Error(), http.StatusInternalServerError)
		addMeshkitErr(res, models.ErrMarshal(err, obj))

		go h.EventsBuffer.Publish(res)
		return
	}
	eventBuilder.WithDescription(fmt.Sprintf("Design %s saved", strings.Join(names, ",")))
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(data))
	res.Details = "\"" + strings.Join(names, ",") + "\" design saved"
	res.Summary = "Changes to the \"" + strings.Join(names, ",") + "\" design have been saved."
	// go h.EventsBuffer.Publish(res)
}

// Since the client currently does not support pattern imports and externalized variables, the first(import) stage of pattern engine
// is evaluated here to simplify the pattern file such that it is valid when a deploy takes place
func evalImportAndReferenceStage(p *pCore.Pattern) (newp pCore.Pattern) {
	sap := &serviceActionProvider{}
	sip := &serviceInfoProvider{}
	chain := stages.CreateChain()
	chain.
		Add(stages.Import(sip, sap)).
		Add(stages.Filler(false)).
		Add(func(data *stages.Data, err error, next stages.ChainStageNextFunction) {
			data.Lock.Lock()
			newp = *data.Pattern
			data.Lock.Unlock()
		}).
		Process(&stages.Data{
			Pattern: p,
		})
	return newp
}

// Only pass Meshkit err here or there will be a panic
func addMeshkitErr(res *meshes.EventsResponse, err error) {
	if err != nil {
		res.EventType = meshes.EventType_ERROR
		res.ProbableCause = errors.GetCause(err)
		res.SuggestedRemediation = errors.GetRemedy(err)
		res.Details = err.Error()
		res.Summary = errors.GetSDescription(err)
		res.ErrorCode = errors.GetCode(err)
	}
}

// swagger:route PUT /api/pattern/{sourcetype} PatternsAPI idUpdateMeshplayPattern
// Handle PUT request for Meshplay Pattern with the given payload
//
// Updates the pattern with the given payload
// responses:
//
//	200: meshplayPatternResponseWrapper
func (h *Handler) handlePatternUpdate(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()
	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("update").ActedUpon(userID)

	res := meshes.EventsResponse{
		Component:     "core",
		ComponentName: "Design",
		OperationId:   guid.NewString(),
		EventType:     meshes.EventType_INFO,
	}

	sourcetype := mux.Vars(r)["sourcetype"]
	if sourcetype == "" {
		http.Error(rw, "missing route variable \"source-type\"", http.StatusBadRequest)
		addMeshkitErr(&res, ErrSaveApplication(fmt.Errorf("missing route \"source-type\"")))

		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrSaveApplication(fmt.Errorf("missing route variable \"source-type\" (one of %s, %s, %s)", models.K8sManifest, models.DockerCompose, models.HelmChart)),
		}).WithDescription("Please provide design source-type").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		go h.EventsBuffer.Publish(&res)
		return
	}

	var parsedBody *MeshplayPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		http.Error(rw, ErrRetrieveData(err).Error(), http.StatusBadRequest)
		return
	}

	actedUpon := &userID
	if parsedBody.PatternData != nil && parsedBody.PatternData.ID != nil {
		actedUpon = parsedBody.PatternData.ID
	}

	eventBuilder.ActedUpon(*actedUpon)

	token, err := provider.GetProviderToken(r)
	if err != nil {
		event := eventBuilder.WithSeverity(events.Critical).WithMetadata(map[string]interface{}{
			"error": ErrRetrieveUserToken(err),
		}).WithDescription("No auth token provided in the request.").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		http.Error(rw, ErrRetrieveUserToken(err).Error(), http.StatusInternalServerError)
		return
	}
	format := r.URL.Query().Get("output")

	if parsedBody.CytoscapeJSON != "" {
		pf, err := pCore.NewPatternFileFromCytoscapeJSJSON(parsedBody.Name, []byte(parsedBody.CytoscapeJSON))
		if err != nil {
			errAppSave := ErrSaveApplication(err)
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(rw, "%s", err)

			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": errAppSave,
			}).WithDescription(fmt.Sprintf("Error saving design %s", parsedBody.PatternData.Name)).Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			addMeshkitErr(&res, ErrSavePattern(err))
			go h.EventsBuffer.Publish(&res)
			return
		}

		pfByt, err := pf.ToYAML()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, "%s", err)
			addMeshkitErr(&res, ErrSavePattern(err))
			go h.EventsBuffer.Publish(&res)
			return
		}

		patternName, err := models.GetPatternName(string(pfByt))
		if err != nil {
			h.log.Error(ErrGetPattern(err))
			http.Error(rw, ErrGetPattern(err).Error(), http.StatusBadRequest)
			addMeshkitErr(&res, ErrGetPattern(err))
			go h.EventsBuffer.Publish(&res)
			return
		}

		meshplayPattern := &models.MeshplayPattern{
			Name:        patternName,
			PatternFile: string(pfByt),
			Location: map[string]interface{}{
				"host": "",
				"path": "",
				"type": "local",
			},
			Type: sql.NullString{
				String: sourcetype,
				Valid:  true,
			},
		}
		if parsedBody.PatternData != nil {
			meshplayPattern.ID = parsedBody.PatternData.ID
		}
		if parsedBody.Save {
			resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
			if err != nil {
				errAppSave := ErrSaveApplication(err)
				h.log.Error(errAppSave)

				rw.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(rw, "%s", err)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": errAppSave,
				}).WithDescription(fmt.Sprintf("Error saving design %s", parsedBody.PatternData.Name)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				addMeshkitErr(&res, ErrSavePattern(err))
				go h.EventsBuffer.Publish(&res)
				return
			}

			eventBuilder.WithSeverity(events.Informational)

			go h.config.ApplicationChannel.Publish(userID, struct{}{})
			h.formatPatternOutput(rw, resp, format, &res, eventBuilder)
			event := eventBuilder.Build()
			// go h.config.EventBroadcaster.Publish(userID, event)
			_ = provider.PersistEvent(event)

			return
		}

		byt, err := json.Marshal([]models.MeshplayPattern{*meshplayPattern})
		if err != nil {
			h.log.Error(ErrEncodePattern(err))
			http.Error(rw, ErrEncodePattern(err).Error(), http.StatusInternalServerError)
			return
		}

		h.formatPatternOutput(rw, byt, format, &res, eventBuilder)
		return
	}
	meshplayPattern := parsedBody.PatternData
	meshplayPattern.Type = sql.NullString{
		String: sourcetype,
		Valid:  true,
	}
	resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
	if err != nil {
		obj := "save"
		errAppSave := ErrSaveApplication(err)
		h.log.Error(errAppSave)

		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "%s", err)

		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": errAppSave,
		}).WithDescription(fmt.Sprintf("Error saving design %s", parsedBody.PatternData.Name)).Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		addMeshkitErr(&res, ErrApplicationFailure(err, obj))
		go h.EventsBuffer.Publish(&res)
		return
	}
	go h.config.ApplicationChannel.Publish(userID, struct{}{})

	eventBuilder.WithSeverity(events.Informational)
	h.formatPatternOutput(rw, resp, format, &res, eventBuilder)
	event := eventBuilder.Build()
	_ = provider.PersistEvent(event)
	go h.config.EventBroadcaster.Publish(userID, event)

}

// swagger:route POST /api/pattern/{sourcetype} PatternsAPI idPostPatternFileRequest
// Handle POST request for Pattern Files
//
// Creates a new Pattern with source-content
// responses:
//  200: meshplayPatternResponseWrapper

// PatternFileRequestHandler will handle requests of both type GET and POST
// on the route /api/pattern
func (h *Handler) DesignFileRequestHandlerWithSourceType(
	rw http.ResponseWriter,
	r *http.Request,
	prefObj *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	if r.Method == http.MethodPost {
		h.handlePatternPOST(rw, r, prefObj, user, provider)
		return
	}

	if r.Method == http.MethodPut {
		h.handlePatternUpdate(rw, r, prefObj, user, provider)
		return
	}
}

// swagger:route GET /api/pattern/types PatternsAPI typeGetMeshplayPatternTypesHandler
// Handle GET request for Meshplay Pattern types
//
// Get pattern file types
// responses:
//
//	200: meshplayApplicationTypesResponseWrapper
func (h *Handler) GetMeshplayDesignTypesHandler(
	rw http.ResponseWriter,
	_ *http.Request,
	_ *models.Preference,
	_ *models.User,
	_ models.Provider,
) {
	response := models.GetDesignsTypes()
	b, err := json.Marshal(response)
	if err != nil {
		obj := "available types"
		h.log.Error(models.ErrMarshal(err, obj))
		http.Error(rw, models.ErrMarshal(err, obj).Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(b))
}

// swagger:route GET /api/pattern/download/{id}/{sourcetype} PatternsAPI typeGetPatternSourceContent
// Handle GET request for Meshplay Patterns with of source content
//
// Get the pattern source-content
// responses:
//  200: meshplayPatternSourceContentResponseWrapper

// GetMeshplayPatternHandler fetched the design using the given id and sourcetype
func (h *Handler) GetMeshplayPatternSourceHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	designID := mux.Vars(r)["id"]
	resp, err := provider.GetDesignSourceContent(r, designID)
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		return
	}

	var mimeType string
	sourcetype := mux.Vars(r)["sourcetype"]

	if models.DesignType(sourcetype) == models.HelmChart { //serve the content in a tgz file
		mimeType = "application/x-tar"
	} else { // serve the content in yaml file
		mimeType = "application/x-yaml"
	}
	reader := bytes.NewReader(resp)
	rw.Header().Set("Content-Type", mimeType)
	_, err = io.Copy(rw, reader)
	if err != nil {
		h.log.Error(ErrApplicationSourceContent(err, "download"))
		http.Error(rw, ErrApplicationSourceContent(err, "download").Error(), http.StatusInternalServerError)
	}
}
