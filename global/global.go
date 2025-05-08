package global

import (
	"github.com/sondoannam/go-ecommerce-backend-api/pkg/logger"
	"github.com/sondoannam/go-ecommerce-backend-api/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
)