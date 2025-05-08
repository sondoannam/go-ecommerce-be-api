package user

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public routers
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}

	// private routers
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Authen())
	{
		userRouterPrivate.GET("/get_info")
	}
}