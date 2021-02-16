package api

import (
	"github.com/andreposman/magic-number/internal/asset/controller"
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/config"
	"github.com/gin-gonic/gin"
)

// Init serves the endpoint that recieves the request and then output the calculation results
func Init() {
	request := new(assetModel.Request)
	request.AssetSymbol = config.DebugData().Asset
	request.DesiredMonthlyIncome = config.DebugData().DesiredMonthlyIncome

	asset := controller.GetAsset(request)

	r := gin.Default()
	r.GET("/asset", func(c *gin.Context) {
		c.JSON(200, gin.H{"asset": asset})
	})

	r.Run(":8080")
}
