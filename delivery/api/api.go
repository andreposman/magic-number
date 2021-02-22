package api

import (
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/gin-gonic/gin"
)

// Init serves the endpoint that recieves the request and then output the calculation results
func Init(asset *assetModel.Asset) {
	r := gin.Default()

	r.GET("/asset", func(c *gin.Context) {
		c.JSON(200, gin.H{"asset": asset})
	})

	r.Run(":8080")
}
