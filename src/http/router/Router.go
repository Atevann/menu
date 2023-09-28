package router

import (
	"github.com/gin-gonic/gin"
	"menu/src/di"
	"menu/src/http/handlers/api"
	"menu/src/http/handlers/status"
	"menu/src/http/router/middleware"
)

func Router(container di.Di) *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api")

	configureInfoRoutes(router)
	configureApiRoutes(apiGroup, container)

	return router
}

// configureInfoRoutes Конфиг информационных роутов
func configureInfoRoutes(r gin.IRouter) {
	r.GET("/health", status.HealthHandler)
}

// configureApiRoutes Конфиг роутов API
func configureApiRoutes(r gin.IRouter, container di.Di) {
	r.GET("/dishes", middleware.ProvideDependency(api.GetAllDishes, container))
}
