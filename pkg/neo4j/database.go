package neo4j

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

func GetDriver(ctx context.Context, DBURI, DBUser, DBPass string) neo4j.DriverWithContext {
	driver, err := newDriver(ctx, DBURI, DBUser, DBPass)

	if err != nil {
		log.Fatalln("Couldn't connect to the db: ", err)
	}

	return driver
}

func newDriver(ctx context.Context, uri, username, password string) (neo4j.DriverWithContext, error) {
	// Create Driver
	driverWithContext, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))

	// Handle any driver creation errors
	if err != nil {
		return nil, err
	}

	// Verify Connectivity
	err = driverWithContext.VerifyConnectivity(ctx)

	// If connectivity fails, handle the error
	if err != nil {
		return nil, err
	}

	log.Println("Obtained new driver with context.")
	return driverWithContext, nil
}

// CloseDriver call on application exit
func CloseDriver(ctx context.Context, driver neo4j.DriverWithContext) {
	log.Println("Closing the driver")
	err := driver.Close(ctx)

	if err != nil {
		log.Fatalln("Couldn't close the db: ", err)
	}
}
