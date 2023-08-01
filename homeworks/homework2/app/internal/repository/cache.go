package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"go-mongo-crud-rest-api/internal/model"
	"strings"
)

// ElasticsearchRepository is the implementation of the Repository interface using Elasticsearch as the data store.
type ElasticsearchRepository struct {
	esClient *elasticsearch.Client
	index    string
}

// NewElasticsearchRepository creates a new instance of ElasticsearchRepository.
func NewElasticsearchRepository(esClient *elasticsearch.Client, index string) *ElasticsearchRepository {
	return &ElasticsearchRepository{
		esClient: esClient,
		index:    index,
	}
}

// GetUser retrieves a user from Elasticsearch based on the given email.
func (er *ElasticsearchRepository) GetUser(ctx context.Context, email string) (model.User, error) {
	var user model.User

	// Create a search query to find the user by email in Elasticsearch.
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]string{
				"email.keyword": email, // Use the keyword subfield for exact match search
			},
		},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return user, err
	}

	// Execute the search query against Elasticsearch.
	response, err := er.esClient.Search(
		er.esClient.Search.WithContext(ctx),
		er.esClient.Search.WithIndex(er.index),
		er.esClient.Search.WithBody(strings.NewReader(string(queryJSON))),
	)

	if err != nil {
		return user, err
	}
	defer response.Body.Close()

	if response.IsError() {
		return user, errors.New(response.Status())
	}

	// Parse the search response and retrieve the user data.
	var searchData map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&searchData); err != nil {
		return user, err
	}

	// Assuming the user data is stored in the "hits" field of the response.
	hits, found := searchData["hits"].(map[string]interface{})["hits"].([]interface{})
	if !found || len(hits) == 0 {
		return user, errors.New("user not found")
	}

	// Unmarshal the user data into the model.User struct.
	if err := json.Unmarshal([]byte(hits[0].(map[string]interface{})["_source"].(string)), &user); err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser creates a new user in Elasticsearch.
func (er *ElasticsearchRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	// Convert the user struct to JSON.
	userJSON, err := json.Marshal(user)
	if err != nil {
		return model.User{}, err
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
		return model.User{}, err
	}
	defer response.Body.Close()

	if response.IsError() {
		return model.User{}, errors.New(response.Status())
	}

	// Return the user data (not necessary, but can be useful to confirm)
	return user, nil
}

// UpdateUser updates an existing user in Elasticsearch.
func (er *ElasticsearchRepository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	// Convert the user struct to JSON.
	userJSON, err := json.Marshal(user)
	if err != nil {
		return model.User{}, err
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
		return model.User{}, err
	}
	defer response.Body.Close()

	if response.IsError() {
		return model.User{}, errors.New(response.Status())
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
	defer response.Body.Close()

	if response.IsError() {
		return errors.New(response.Status())
	}

	return nil
}
