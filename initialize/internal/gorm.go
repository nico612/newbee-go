package internal

import (
	"github.com/nico612/newbee-goo/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config() *gorm.Config {

	// 用于在执行数据库迁移（migration）时是否禁用外键约束
	// 当数据库表之间存在外键关系时，正常情况下，在执行数据库迁移时会尝试应用外键约束
	// 但在某些情况下，您可能希望在迁移时禁用外键约束，以便更灵活地控制迁移流程。
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}

	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	// 配置数据库日志级别
	switch global.Config.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
