package routers

import (
	"learn-elastic-search/internal/services"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

func ElasticSearchRouter(elasticsearchClient *elasticsearch.Client, r *gin.RouterGroup) {
	r.POST("/index", func(c *gin.Context) {
		services.CreateIndex(elasticsearchClient, c)
	})

	r.POST("/index/document", func(c *gin.Context) {
		services.IngestDocument(elasticsearchClient, c)
	})

	r.GET("/search", func(c *gin.Context) {
		services.Search(elasticsearchClient, c)
	})

}