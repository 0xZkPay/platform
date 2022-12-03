package status

import (
	"github.com/0xZkPay/platform/api/types"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/status")
	{
		g.GET("", status)
	}
}

func status(c *gin.Context) {
	status := types.ApiResponse{
		Status:  200,
		Message: "ONLINE",
	}
	c.JSON(200, status)
}
