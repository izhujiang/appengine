package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine, relativePath string) {
	if relativePath == "" {
		relativePath = "/api"
	}
	apiGroup := engine.Group(relativePath)
	apiGroup.GET("/hello", helloworld)
	v1 := apiGroup.Group("/v1")
	v1.GET("/hello", v1_helloworld)

}

func helloworld(c *gin.Context) {
	hello := gin.H{
		"gin": "hello world",
	}
	c.JSON(http.StatusOK, hello)

}

func v1_helloworld(c *gin.Context) {
	hello := gin.H{
		"gin": "hello world v1",
	}
	c.JSON(http.StatusOK, hello)
}
