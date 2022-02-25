package errors

import (
	"fmt"
	"net/http"
)

// Module constants definition.
const (
	ModuleCommon = 00
)

// Common module error codes definition.
var (
	ErrInvalidRequest = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 1)
	ErrUnauthorized   = fmtErrorCode(http.StatusUnauthorized, ModuleCommon, 1)
	ErrInternalServer = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 1)
	ErrNoResponse     = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 2)
	ErrNotFound       = fmtErrorCode(http.StatusNotFound, ModuleCommon, 1)
)

// GetErrorMessage gets error message from errorMessageMap.
func GetErrorMessage(errCode ErrorCode, args ...interface{}) string {
	if len(args) > 0 {
		msg, ok := args[0].(string)
		if ok {
			return msg
		}
	}

	msg := http.StatusText(errCode.Status())
	return fmt.Sprintf(msg, args...)
}
