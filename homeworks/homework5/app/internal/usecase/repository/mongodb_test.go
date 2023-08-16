//go:build integration
// +build integration

package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

func Test_GetUser(t *testing.T) {
	// create a database connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// create a repository
	repo := NewRepository(client.Database("users"))

	usr, err := repo.GetUser(context.Background(), "TpDMbHjZMl@example.com")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(usr.Email)
}
