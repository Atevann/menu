package status

import "github.com/gin-gonic/gin"

func HealthHandler(c *gin.Context) {
	c.JSONP(200, struct {
		Status string
	}{Status: "OK"})
}
