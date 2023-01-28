package main

import (
	"context"
	"iam/cmd/iam/webservice"
	"iam/configs"
	pkgneo4j "iam/pkg/neo4j"
	"log"
	"os"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func initializeDependencies(ctx context.Context) (*configs.Config, neo4j.DriverWithContext) {
	c, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Error encountered when loading the configuration: %v", err)
	}

	driver := pkgneo4j.GetDriver(ctx, c.DBUri, c.DBUser, c.DBPass)

	if err != nil {
		log.Fatalln("Couldn't connect to the db: ", err)
	}

	return c, driver
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Must provide program argument.")
	}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	conf, db := initializeDependencies(ctxWithTimeout)

	switch os.Args[1] {
	case "webservice":
		err := webservice.RunWebservice(conf.GinMode, db)
		if err != nil {
			pkgneo4j.CloseDriver(ctxWithTimeout, db)
			log.Fatalln("Webservice stopped: ", err)
		}
	default:
		pkgneo4j.CloseDriver(ctxWithTimeout, db)
		log.Fatalln("Mistakes were made")
	}
}
