package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"go-mongo-crud-rest-api/internal/entitiy"
	"go-mongo-crud-rest-api/internal/usecase"
	"net/http"
	"strings"

	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// SearchResults wraps the Elasticsearch search response.
type SearchResults struct {
	Total int    `json:"total"`
	Hits  []*Hit `json:"hits"`
}

// Hit wraps the document returned in search response.
type Hit struct {
	URL        string        `json:"url"`
	Sort       []interface{} `json:"sort"`
	Highlights *struct {
		Title      []string `json:"title"`
		Alt        []string `json:"alt"`
		Transcript []string `json:"transcript"`
	} `json:"highlights,omitempty"`
}

// ElasticsearchRepository is the implementation of the Repository interface using Elasticsearch as the data store.
type ElasticsearchRepository struct {
	esClient *es8.Client
	index    string
}

// NewElasticsearchRepository creates a new instance of ElasticsearchRepository.
func NewElasticsearchRepository(esClient *es8.Client, index string) *ElasticsearchRepository {
	if err := createIndexIfNotExists(esClient, index); err != nil {
		log.Print(err)
	}

	return &ElasticsearchRepository{
		esClient: esClient,
		index:    index,
	}
}

func createIndexIfNotExists(esClient *es8.Client, index string) error {
	// First, check if the index exists.
	res, err := esClient.Indices.Exists([]string{index})
	if err != nil {
		return err
	}

	// If the index does not exist, create it.
	if res.StatusCode == http.StatusNotFound {
		req := esapi.IndicesCreateRequest{
			Index: index,
			Body: strings.NewReader(`
				{
					"mappings": {
						"properties": {
							"email": {
								"type": "keyword"
							}
							// Specify other fields as needed.
						}
					}
				}`),
		}

		res, err := req.Do(context.Background(), esClient)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return errors.New("Error creating index: " + res.String())
		}
	}

	return nil
}

func (er *ElasticsearchRepository) FindUser(ctx context.Context, email string) (*entitiy.User, error) {
	query := fmt.Sprintf(`{ "query": { "match": {"email" : "%s"} } }`, email)
	resp, err := er.esClient.Search(
		er.esClient.Search.WithIndex(er.index),
		er.esClient.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}

	log.Printf(resp.String())

	var result esSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Hits.Hits) == 0 {
		return nil, usecase.ErrUserNotFound
	}

	return &result.Hits.Hits[0].Source, nil
}

type esSearchResponse struct {
	Hits struct {
		Hits []struct {
			Source entitiy.User `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// GetUser retrieves a user from Elasticsearch based on the given email.
func (er *ElasticsearchRepository) GetUser(ctx context.Context, email string) (*entitiy.User, error) {
	query := fmt.Sprintf(`{ "query": { "term": {"email.keyword" : "%s"} } }`, email)
	resp, err := er.esClient.Search(
		er.esClient.Search.WithIndex(er.index),
		er.esClient.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}

	log.Printf(resp.String())

	var result esSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Hits.Hits) == 0 {
		return nil, usecase.ErrUserNotFound
	}

	return &result.Hits.Hits[0].Source, nil
}

// CreateUser creates a new user in Elasticsearch.
func (er *ElasticsearchRepository) CreateUser(ctx context.Context, user *entitiy.User) (*entitiy.User, error) {
	// Convert the user struct to JSON.
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	// Prepare the index request for Elasticsearch.
	req := esapi.IndexRequest{
		Index:      er.index,
		DocumentID: user.Email, // Assuming the email is used as the document ID.
		Body:       strings.NewReader(string(userJSON)),
	}

	// Execute the index request against Elasticsearch.
	response, err := req.Do(ctx, er.esClient)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = response.Body.Close()
	}()

	if response.IsError() {
		return nil, errors.New(response.Status())
	}

	// Return the user data (not necessary, but can be useful to confirm)
	return user, nil
}

// UpdateUser updates an existing user in Elasticsearch.
func (er *ElasticsearchRepository) UpdateUser(ctx context.Context, user *entitiy.User) (*entitiy.User, error) {
	// Convert the user struct to JSON.
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	// Prepare the update request for Elasticsearch.
	req := esapi.UpdateRequest{
		Index:      er.index,
		DocumentID: user.Email, // Assuming the email is used as the document ID.
		Body:       strings.NewReader(`{"doc":` + string(userJSON) + `}`),
	}

	// Execute the update request against Elasticsearch.
	response, err := req.Do(ctx, er.esClient)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = response.Body.Close()
	}()

	if response.IsError() {
		return nil, errors.New(response.Status())
	}

	return user, nil
}

// DeleteUser deletes a user from Elasticsearch based on the given email.
func (er *ElasticsearchRepository) DeleteUser(ctx context.Context, email string) error {
	// Prepare the delete request for Elasticsearch.
	req := esapi.DeleteRequest{
		Index:      er.index,
		DocumentID: email, // Assuming the email is used as the document ID.
	}

	// Execute the delete request against Elasticsearch.
	response, err := req.Do(ctx, er.esClient)
	if err != nil {
		return err
	}
	defer func() {
		err = response.Body.Close()
	}()

	if response.IsError() {
		return errors.New(response.Status())
	}

	return nil
}
