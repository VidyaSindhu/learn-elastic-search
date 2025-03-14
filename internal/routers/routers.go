package routers

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

var elasticsearchClient *elasticsearch.Client;

func SetupRouter(esClient *elasticsearch.Client) *gin.Engine {
	elasticsearchClient = esClient
	r := gin.Default()
	v1Apis := r.Group("/v1")
	ElasticSearchRouter(elasticsearchClient, v1Apis.Group("/elastic-search"))
	return r
}