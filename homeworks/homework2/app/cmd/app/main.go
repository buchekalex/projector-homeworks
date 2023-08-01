package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"go-mongo-crud-rest-api/internal/http"
	repository2 "go-mongo-crud-rest-api/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
	"time"

	nethttp "net/http"
)

func main() {
	const defaultMongoURI = "mongodb://127.0.0.1:27017"

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Printf("MONGO_URI is not set, using default value: %s", defaultMongoURI)
		mongoURI = "mongodb://127.0.0.1:27017"
	}

	var httpPort int
	var err error

	httpPortEnv := os.Getenv("API_HTTP_PORT")
	if httpPortEnv == "" {
		log.Printf("HTTP_PORT is not set, using default value: 8080")
		httpPort = 8080
	} else {
		httpPort, err = strconv.Atoi(httpPortEnv)
		if err != nil {
			log.Fatalf("HTTP_PORT is invalid: %s", err.Error())
		}
	}

	// create a database connection
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// create a repository
	repo := repository2.NewRepository(client.Database("users"))
	elasticRepository, err := NewElasticsearchRepository()
	if err != nil {
		log.Fatal(err)
	}

	// create an http server
	server := http.NewServer(repo, elasticRepository)

	// create a gin router
	router := gin.Default()
	{
		router.GET("/users/:email", server.GetUser)
		router.POST("/users", server.CreateUser)
		router.PUT("/users/:email", server.UpdateUser)
		router.DELETE("/users/:email", server.DeleteUser)
		router.GET("/health", server.HealthCheck)
	}

	// start the router
	router.Run(fmt.Sprintf(":%d", httpPort))
}

// InitializeElasticsearchClient initializes and returns an Elasticsearch client.
func InitializeElasticsearchClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
		Transport: &nethttp.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DisableCompression:    true,
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Elasticsearch client: %v", err)
	}

	// Check if the client is up and running
	res, err := esClient.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to ping Elasticsearch: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("Elasticsearch returned an error response: %s", res.String())
	}

	log.Printf("Elasticsearch cluster info: %s", res.String())

	return esClient, nil
}

// NewElasticsearchRepository creates a new instance of ElasticsearchRepository with its dependencies.
func NewElasticsearchRepository() (repository2.Repository, error) {
	esClient, err := InitializeElasticsearchClient()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Elasticsearch client: %v", err)
	}

	index := "users"

	return repository2.NewElasticsearchRepository(esClient, index), nil
}
