package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()

		// before request

		ctx.Next()

		// after request

		latency := time.Since(now).Seconds()
		status := ctx.Writer.Status()

		if latency > time.Second.Seconds()*10 {
			// Warning
			log.Printf("Call took more than 10 seconds: status %d - latency %.3g.", status, latency)
		}
	}
}

func LoggerWithFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s - \"%s %s %s %d %s \"%s\" %s\"\n",
			params.TimeStamp.Format(time.RFC3339),
			params.ClientIP,
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	})
}
