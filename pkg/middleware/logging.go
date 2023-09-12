package middleware

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/gin-gonic/gin"
)

// NewLogging
func NewLogging(regexs ...string) gin.HandlerFunc {
	skips := make([]*regexp.Regexp, len(regexs))
	for i, p := range regexs {
		skips[i] = regexp.MustCompile(p)
	}

	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}

		path := ctx.Request.URL.Path

		for _, regex := range skips {
			if regex.MatchString(path) {
				ctx.Next()
				return
			}
		}

		start := time.Now()
		var requetBuf []byte
		if ctx.Request.Body != nil &&
			ctx.Request.Body != http.NoBody &&
			!strings.HasPrefix(ctx.GetHeader("Content-Type"), "multipart/form-data") {
			requetBuf, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requetBuf))
		}
		wrapper := &ResponseWrapper{bodyBuf: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = wrapper

		ctx.Next()

		cost := time.Since(start)

		if len(ctx.Errors) > 0 {
			logging.ErrorContext(ctx.Request.Context(),
				"ERROR [%v] REQUEST [%s %s] BODY [%s] RESPONSE [%d] BODY [%s] TIME [%d] ms",
				ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
				ctx.Request.Method,
				ctx.Request.URL.String(),
				string(requetBuf),
				ctx.Writer.Status(),
				wrapper.bodyBuf.String(),
				cost.Milliseconds(),
			)
		} else {
			logging.InfoContext(ctx.Request.Context(),
				"REQUEST [%s %s] BODY [%s] RESPONSE [%d] BODY [%s] TIME [%d] ms",
				ctx.Request.Method,
				ctx.Request.URL.String(),
				string(requetBuf),
				ctx.Writer.Status(),
				wrapper.bodyBuf.String(),
				cost.Milliseconds(),
			)
		}
	}
}

type ResponseWrapper struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (rw *ResponseWrapper) Write(buf []byte) (int, error) {
	if n, err := rw.bodyBuf.Write(buf); err != nil {
		return n, err
	}
	return rw.ResponseWriter.Write(buf)
}
