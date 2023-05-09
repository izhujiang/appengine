package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/izhujiang/appengine/endpoint/api"
	"github.com/izhujiang/appengine/endpoint/middleware"
)

func Register(engine *gin.Engine, relativePath string) {
	authGroup := engine.Group("/auth")
	authGroup.POST("/register", signup)
	authGroup.POST("/login", login)

	entryGroup := engine.Group("/entry")
	entryGroup.Use(middleware.JWTAuthMiddleware())
	entryGroup.POST("", addEntry)
	entryGroup.GET("", getAllEntries)

	api.Register(engine, "/more")
}
