package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gemsorg/beehive/pkg/authorization"
	"github.com/gemsorg/beehive/pkg/database"
	"github.com/gemsorg/beehive/pkg/datastore"
	"github.com/gemsorg/beehive/pkg/service"
	"github.com/joho/godotenv"

	"github.com/gemsorg/beehive/pkg/server"
)

func main() {
	environment := flag.String("env", "local", "use compose in compose-dev")
	flag.Parse()

	if *environment == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Connect to db
	db, err := database.Connect()
	if err != nil {
		log.Fatal("mysql connection error", err)
	}
	defer db.Close()
	ds := datastore.NewDatastore(db)
	authorizer := authorization.NewAuthorizer()
	svc := service.New(ds, authorizer)
	s := server.New(db, svc)
	servicePort := os.Getenv("BEEHIVE_SERVICE_PORT")
	log.Println("info", fmt.Sprintf("Starting Beehive Service %s", servicePort))
	http.Handle("/", s)
	http.ListenAndServe(fmt.Sprintf(":%s", servicePort), nil)
}
