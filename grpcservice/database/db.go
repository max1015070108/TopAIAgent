package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	// DB 全局数据库实例
	DB   *gorm.DB
	once sync.Once
)

// Config 数据库配置
type Config struct {
	DBPath string
	Debug  bool
}

// InitDB 初始化数据库连接
func InitDB(cfg *Config) error {
	var err error

	once.Do(func() {
		gormConfig := &gorm.Config{}
		if cfg.Debug {
			gormConfig.Logger = logger.Default.LogMode(logger.Info)
		}

		DB, err = gorm.Open(sqlite.Open(cfg.DBPath), gormConfig)
		if err != nil {
			err = fmt.Errorf("failed to connect database: %w", err)
			return
		}

		// 自动迁移表结构
		if err = DB.AutoMigrate(&Model{}, &NetworkModel{}); err != nil {
			err = fmt.Errorf("failed to migrate database: %w", err)
			return
		}
	})

	return err
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// GetDB 获取数据库实例（可选的辅助方法）
func GetDB() *gorm.DB {
	return DB
}
