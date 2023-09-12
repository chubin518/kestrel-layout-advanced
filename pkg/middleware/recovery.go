package middleware

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
	"strings"

	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/gin-gonic/gin"
)

// NewRecovery
func NewRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					logging.ErrorContext(ctx.Request.Context(),
						"[URL]: %s [ERROR]: %v [REQUEST]: %s",
						ctx.Request.URL.String(),
						err, string(httpRequest),
					)
					// If the connection is dead, we can't write a status to it.
					ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}
				_, file, line, _ := runtime.Caller(3)
				logging.ErrorContext(ctx.Request.Context(),
					"[Recovery from panic %s] [URL]: %s [ERROR]: %v [REQUEST]: %s",
					fmt.Sprintf("%s:%d", file, line),
					ctx.Request.URL.String(),
					err,
					string(httpRequest),
				)
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
