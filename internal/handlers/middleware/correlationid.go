package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var CorrelationIDHeader = "X-Request-Id"

func CorrelationID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		correlationID := ctx.Request.Header.Get(CorrelationIDHeader)

		if correlationID == "" {
			log.Println("Call didn't provide a correlation id, one will be generated.")
			correlationID = uuid.New().String()
		}

		ctx.Set("correlation-id", correlationID)

		//log.SetPrefix("correlation-id: " + correlationID + " ")
	}
}
