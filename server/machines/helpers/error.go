package helpers

import (
	"github.com/khulnasoft/meshplay/server/models"
	"github.com/khulnasoft/meshkit/errors"
)

var (
	ErrAutoRegisterCode = "1556"
)

func ErrAutoRegister(err error, connType string) error {
	return errors.New(ErrAutoRegisterCode, errors.Alert, []string{}, []string{}, []string{}, []string{})
}

func IsConnectionUpdateErr(err error) bool {
	return errors.GetCode(err) == models.ErrUpdateConnectionStatusCode
}
