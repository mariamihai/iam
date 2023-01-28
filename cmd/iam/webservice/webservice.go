package webservice

import (
	"github.com/gin-gonic/gin"
	"iam/pkg/handlers/middleware"
	"iam/pkg/handlers/webservice"
	"log"
)

// Todo: add db
func RunWebservice() error {
	w := webservice.NewWebservice()

	router := gin.Default()

	router.Use(
		middleware.CorrelationId(),
		middleware.Logger(),
		//middleware.LoggerWithFormatter(),
		// Recovers from any panics and writes a 500 if there was one
		gin.Recovery(),
	)

	router.GET("/ping", w.GetPing)

	log.Println("Starting server on port 8080.")
	return router.Run()
}
