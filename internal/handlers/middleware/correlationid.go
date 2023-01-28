package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

var CorrelationIdHeader = "X-Request-Id"

func CorrelationId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		correlationID := ctx.Request.Header.Get(CorrelationIdHeader)

		if correlationID == "" {
			log.Println("Call didn't provide a correlation id, one will be generated.")
			correlationID = uuid.New().String()
		}

		ctx.Set("correlation-id", correlationID)

		//log.SetPrefix("correlation-id: " + correlationID + " ")
	}
}
