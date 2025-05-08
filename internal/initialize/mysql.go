package initialize

import (
	"fmt"
	"time"

	"github.com/sondoannam/go-ecommerce-backend-api/global"
	"github.com/sondoannam/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.String("error", err.Error()))
		panic(err)
	}
}

func checkErrorPrint(err error, errString string) {
	if err != nil {
		fmt.Printf(errString, err.Error())
	}
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "Failed to connect to database!")

	global.Logger.Info("Connected to database successfully!")
	global.Mdb = db

	// set connection pool settings
	SetPool()

	migrateTables()
}

func SetPool() {
	sqlDb, err := global.Mdb.DB()

	checkErrorPrint(err, "Failed to get database connection pool!")

	m := global.Config.Mysql
	// sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err :=global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	checkErrorPrint(err, "Failed to migrate tables!")
}
