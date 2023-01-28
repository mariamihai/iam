package main

import (
	"iam/cmd/iam/webservice"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Must provide program argument.")
	}

	switch os.Args[1] {
	case "webservice":
		log.Fatal("Webservice stopped", webservice.RunWebservice())
	default:
		log.Fatal("Mistakes were made")
	}
}
