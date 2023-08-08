package initialize

import (
	"github.com/nico612/newbee-goo/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
