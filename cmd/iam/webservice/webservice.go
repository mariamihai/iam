package webservice

import (
	"iam/internal/handlers/middleware"
	"iam/internal/handlers/webservice"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func RunWebservice(ginMode string, _ neo4j.DriverWithContext) error {
	w := webservice.NewWebservice()

	router := gin.Default()
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}
	gin.SetMode(ginMode)

	router.Use(
		middleware.CorrelationID(),
		middleware.Logger(),
		//middleware.LoggerWithFormatter(),
		// Recovers from any panics and writes a 500 if there was one
		gin.Recovery(),
	)

	router.GET("/ping", w.GetPing)

	log.Println("Starting server.")
	return router.Run()
}
