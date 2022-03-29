package utils

import (
	"fmt"
	cfg "github.com/dollarkillerx/common/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitPostgres ...
func InitPostgres(config cfg.PostgresConfiguration, debug bool) (db *gorm.DB, err error) {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	//	logger.Config{
	//		SlowThreshold: time.Second,   // 慢 SQL 阈值
	//		LogLevel:      logger.Silent, // 日志级别
	//		IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:      false,         // 禁用彩色打印
	//	},
	//)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.DBName, config.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetMaxOpenConns(100)

	if debug {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	return db, nil
}
