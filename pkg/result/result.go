package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	OK                  = OfResult(http.StatusOK, "OK")
	BadRequest          = OfResult(http.StatusBadRequest, "Bad Request")
	Unauthorized        = OfResult(http.StatusUnauthorized, "Unauthorized")
	Forbidden           = OfResult(http.StatusForbidden, "Forbidden")
	NotFound            = OfResult(http.StatusNotFound, "Not Found")
	MethodNotAllowed    = OfResult(http.StatusMethodNotAllowed, "Method Not Allowed")
	RequestTimeout      = OfResult(http.StatusRequestTimeout, "Request Timeout")
	InternalServerError = OfResult(http.StatusInternalServerError, "Internal Server Error")
	ServiceUnavailable  = OfResult(http.StatusServiceUnavailable, "Service Unavailable")
	ServiceError        = OfResult(10001, "Service Error")
)

// OfResult
func OfResult(code int, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
	}
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// WithCode
func (r *Result) WithCode(code int) *Result {
	r.Code = code
	return r
}

// WithMessage
func (r *Result) WithMessage(message string) *Result {
	r.Message = message
	return r
}

// WithData
func (r *Result) WithData(data any) *Result {
	r.Data = data
	return r
}

// JSON
func (r *Result) JSON(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r)
}
