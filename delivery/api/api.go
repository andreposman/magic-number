package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("v1/asset/:asset", func(c *gin.Context) {
		asset := c.Param("asset")

		c.String(http.StatusOK, "\nThe asset selected was: %s\n", strings.ToUpper(asset))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
