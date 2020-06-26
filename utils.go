package f

import (
	"context"
	"github.com/J-Leg/player-count"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// Constants
const (
	STATSCOL = "population_stats"
	EXCCOL   = "exceptions"
)

func isLocal() bool {
	// Check if ENV variable reserved for GCP exists
	v, ok := os.LookupEnv("LOCAL")
	if !ok {
		return false
	}

	return (v == "Y")
}

func initDb(ctx context.Context) *pc.Collections {
	var newDb *mongo.Database
	var clientOptions *options.ClientOptions
	var dbURI string
	env := os.Getenv("ENV")

	if env == "prd" {
    log.Printf("Environment: PRD")
		clientOptions = options.Client().ApplyURI(os.Getenv("PRD_URI"))
	} else if env == "tst" || env == "dev" {
    log.Printf("Environment: DEV")
		clientOptions = options.Client().ApplyURI(os.Getenv("DEV_URI"))
	} else {
		log.Fatalf("[CRITICAL] Undefined phase!\n")
	}

	newClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("[CRITICAL] Error initialising client. URI: %s\n", dbURI)
	}

  newDb = newClient.Database("games_stats_app")

	err = newClient.Connect(ctx)
	if err != nil {
		log.Fatalf("[CRITICAL] error connecting client. %s\n", err)
	}

	newCollections := pc.Collections{
		Stats:      newDb.Collection(STATSCOL),
		Exceptions: newDb.Collection(EXCCOL),
	}

	return &newCollections
}
