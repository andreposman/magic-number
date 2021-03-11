package api

import (
	"net/http"
	"os"

	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init serves the endpoint that recieves the request and then output the calculation results
func Init(asset *assetModel.Asset) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/asset", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"asset": asset})
	})

	router.Run(":" + port)
}
