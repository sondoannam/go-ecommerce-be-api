package initialize

import (
	"fmt"

	"github.com/sondoannam/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	mysql := global.Config.Mysql
	fmt.Println("Loading database configuration... ", mysql.Username)
	InitLogger()
	global.Logger.Info("Config log ok!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8000")
}