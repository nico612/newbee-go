package global

import (
	"github.com/nico612/newbee-goo/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Viper  *viper.Viper
	Logger *zap.Logger
	Config *config.Server
)
