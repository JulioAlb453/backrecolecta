package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse estructura estandarizada para respuestas de error
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contiene los detalles del error
type ErrorDetail struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// ErrorCodes constantes para tipos de error
const (
	ErrCodeValidation      = "VALIDATION_ERROR"
	ErrCodeNotFound        = "NOT_FOUND"
	ErrCodeUnauthorized    = "UNAUTHORIZED"
	ErrCodeForbidden       = "FORBIDDEN"
	ErrCodeConflict        = "CONFLICT"
	ErrCodeInternalServer  = "INTERNAL_SERVER_ERROR"
	ErrCodeBadRequest      = "BAD_REQUEST"
	ErrCodeInvalidInput    = "INVALID_INPUT"
	ErrCodeDatabaseError   = "DATABASE_ERROR"
	ErrCodeOperationFailed = "OPERATION_FAILED"
)

func RespondError(c *gin.Context, statusCode int, errorCode string, message string, details interface{}) {
	response := ErrorResponse{
		Error: ErrorDetail{
			Code:    errorCode,
			Message: message,
			Details: details,
		},
	}
	c.JSON(statusCode, response)
}

func RespondValidationError(c *gin.Context, message string, details interface{}) {
	RespondError(c, http.StatusBadRequest, ErrCodeValidation, message, details)
}

func RespondBadRequest(c *gin.Context, message string, details interface{}) {
	RespondError(c, http.StatusBadRequest, ErrCodeBadRequest, message, details)
}

func RespondInvalidInput(c *gin.Context, message string) {
	RespondError(c, http.StatusBadRequest, ErrCodeInvalidInput, message, nil)
}

func RespondNotFound(c *gin.Context, resourceType string, identifier string) {
	message := resourceType + " no encontrado"
	details := map[string]string{"identifier": identifier}
	RespondError(c, http.StatusNotFound, ErrCodeNotFound, message, details)
}

func RespondConflict(c *gin.Context, message string, details interface{}) {
	RespondError(c, http.StatusConflict, ErrCodeConflict, message, details)
}

func RespondInternalServerError(c *gin.Context, message string, err error) {
	details := map[string]string{}
	if err != nil {
		details["error"] = err.Error()
	}
	RespondError(c, http.StatusInternalServerError, ErrCodeInternalServer, message, details)
}

func RespondDatabaseError(c *gin.Context, message string, err error) {
	details := map[string]string{}
	if err != nil {
		details["error"] = err.Error()
	}
	RespondError(c, http.StatusInternalServerError, ErrCodeDatabaseError, message, details)
}

func RespondSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func RespondCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func RespondOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func RespondNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
