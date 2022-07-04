package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/routers"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options []Option

// Include : Register routers
func include(opts ...Option) {
	options = append(options, opts...)
}

func ginInit(middleware ...gin.HandlerFunc) *gin.Engine {

	if !config.Config.DevelopmentMode {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	for _, mid := range middleware {
		engine.Use(mid)
	}

	for _, opt := range options {
		opt(engine)
	}

	return engine
}

func Routes() *gin.Engine {
	include(routers.R)

	return ginInit()
}
