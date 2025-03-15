package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"learn-elastic-search/internal/dtos"
	"learn-elastic-search/internal/models"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

var presentIndices []string

func SetPresentIndices (alreadyPresentIndices []string) {
	presentIndices = alreadyPresentIndices
}

func CreateIndex(elasticsearchClient *elasticsearch.Client, c *gin.Context) {
	indexName := c.Request.Header["Index-Name"][0]
	var indexConfig models.IndexConfig

	if err := c.BindJSON(&indexConfig); err != nil {
		panic(err)
	}

	jsonBody, err := json.Marshal(indexConfig)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create index config JSON"})
		return
	}

	res, err := elasticsearchClient.Indices.Create(indexName, 
		elasticsearchClient.Indices.Create.WithBody(bytes.NewReader(jsonBody)))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	presentIndices = append(presentIndices, indexName)

	c.IndentedJSON(http.StatusCreated, dtos.BaseResponseDto {
		Message: "Index added successfully",
		Code: 201,
	})
}

func IngestDocument(elasticsearchClient *elasticsearch.Client, c *gin.Context) {
	indexName := c.Request.Header["Index-Name"][0]

	res, err := elasticsearchClient.Index(indexName, bytes.NewReader(getDocumentBytes(indexName, c)))

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	c.IndentedJSON(http.StatusCreated, dtos.BaseResponseDto {
		Message: "Document ingested in Index successfully",
		Code: 201,
	})
}

func Search(elasticsearchClient *elasticsearch.Client, c *gin.Context) {
	query := c.Request.URL.Query()["query"][0]

	queryTemplate,_ := buildSearchQuery(query)
	
	res, err := elasticsearchClient.Search(
		elasticsearchClient.Search.WithIndex(presentIndices[0]),
		elasticsearchClient.Search.WithBody(strings.NewReader(queryTemplate)),
	)
	defer res.Body.Close()

	if err != nil {
		fmt.Print("error occured while fetching search results")
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}


	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.IndentedJSON(http.StatusOK, dtos.SuccessResponseDto {
		BaseResponseDto:  dtos.BaseResponseDto{
			Message: "SUCCESS",
			Code: 200,
		},
		Data: result,
	})

}

func getDocumentBytes(indexName string, c *gin.Context) ([]byte) {
	switch(indexName) {
	case "restaurant_index":
		var document models.RestaurantElasticsearchDocument
		if err := c.BindJSON(&document); err != nil {
			panic(err)
		}
		jsonBody, err := json.Marshal(document)
		if err != nil {
			fmt.Printf("Failed to marshal JSON: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create index config JSON"})
			return nil
		}
		return jsonBody
	case "food_index":
		var document models.FoodElasticsearchDocument
		if err := c.BindJSON(&document); err != nil {
			panic(err)
		}
		jsonBody, err := json.Marshal(document)
		if err != nil {
			fmt.Printf("Failed to marshal JSON: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create index config JSON"})
			return nil
		}
		return jsonBody
	}

	return nil
}

func buildSearchQuery(queryName string) (string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"restaurant_name":  queryName,
			},
		},
	}

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return "", err
	}
	return string(jsonQuery), nil
}