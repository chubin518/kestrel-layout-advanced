package middleware

import (
	"net/http"

	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const xRequestIDHeaderKey = "X-Request-Id"

// NewRequestId Adds an indentifier to the response using the X-Request-ID header. Passes the X-Request-ID value back to the caller if it's sent in the request headers.
func NewRequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}
		requestid := ctx.GetHeader(xRequestIDHeaderKey)
		if requestid == "" {
			uid, err := uuid.NewRandom()
			if err != nil {
				ctx.Next()
				return
			}
			requestid = uid.String()
			ctx.Request.Header.Set(xRequestIDHeaderKey, requestid)
		}
		ctx.Header(xRequestIDHeaderKey, requestid)
		traceCtx := logging.WithContext(ctx.Request.Context(), map[string]any{
			"traceid": requestid,
		})
		ctx.Request = ctx.Request.WithContext(traceCtx)
		ctx.Next()
	}
}
