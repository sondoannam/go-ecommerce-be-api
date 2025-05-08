package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/sondoannam/go-ecommerce-backend-api/global"
	"github.com/sondoannam/go-ecommerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middleware
	// r.Use() // logging
	// r.Use() // cors
	// r.Use() // limit
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/checkStatus")
	}
	{
		userRouter.ProductRouter.InitProductRouter(mainGroup)
		userRouter.UserRouter.InitUserRouter(mainGroup)
	}
	{
		managerRouter.UserRouter.InitUserRouter(mainGroup)
		managerRouter.AdminRouter.InitAdminRouter(mainGroup)
	}

	return r
}

