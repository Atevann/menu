package api

import (
	"github.com/gin-gonic/gin"
	"menu/src/di"
)

func GetAllDishes(c *gin.Context, container *di.Di) {
	c.JSONP(200, struct {
		Status string
	}{Status: "OK"})
}
