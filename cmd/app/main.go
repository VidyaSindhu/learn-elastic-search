package main

import (
	"encoding/json"
	"learn-elastic-search/internal/routers"
	"learn-elastic-search/internal/services"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	esClientConfig := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	elasticsearchClient, err := elasticsearch.NewClient(esClientConfig)

	if err != nil {
		panic(err)
	}
	
	var alreadyPresentIndices = fetchIndicesOnStartup(elasticsearchClient)
	services.SetPresentIndices(alreadyPresentIndices)
	
	routers := routers.SetupRouter(elasticsearchClient)
	routers.Run("localhost:8080")
}


func fetchIndicesOnStartup(esClient *elasticsearch.Client) []string {
	res, err := esClient.Cat.Indices(
		esClient.Cat.Indices.WithFormat("json"),
	)
	if err != nil {
		log.Fatalf("Error fetching indices: %v", err)
	}
	defer res.Body.Close()

	var indices []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&indices); err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}

	var addedIndices []string
	for _, index := range indices {
		if indexName, ok := index["index"].(string); ok {
			addedIndices = append(addedIndices, indexName)
		}
	}

	return addedIndices
}
