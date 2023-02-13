//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"storage/config"
	"storage/internal/biz"
	"storage/internal/data"
	"storage/internal/server"
	"storage/internal/service"
)

func initApp(localStorageConfig config.LocalStorage, serverConfig config.Server) *gin.Engine {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet))
}
