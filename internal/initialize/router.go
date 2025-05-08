package initialize

import (
	"github.com/gin-gonic/gin"
	c "github.com/sondoannam/go-ecommerce-backend-api/internal/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", Pong)
		v1.GET("/user/1", c.NewUserController().GetUserById)
	}

	return r
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
