package main

import "github.com/gin-gonic/gin"

type Engine interface {
	/* TODO: add methods */
	Register(engine *gin.Engine, relativePath string)
}
