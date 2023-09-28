package middleware

import (
	"github.com/gin-gonic/gin"
	"menu/src/di"
)

// customHandler Кастомный хэндлер с пробросом di
type customHandler func(c *gin.Context, container *di.Di)

// ProvideDependency Функция для проброса di в хэндлеры
func ProvideDependency(handler customHandler, container di.Di) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, &container)
	}
}
