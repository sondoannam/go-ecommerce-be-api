package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/sondoannam/go-ecommerce-backend-api/pkg/logger"
	"github.com/sondoannam/go-ecommerce-backend-api/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Rdb *redis.Client
	Mdb *gorm.DB
)