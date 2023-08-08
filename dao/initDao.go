package dao

import (
	"log"
	"os"
	"time"

	"github.com/abuziming/dousheng_demo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

// 连接数据库
func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 彩色打印
		},
	)

	// 连接数据库
	Db, err := gorm.Open(mysql.Open(config.DBconnect()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicln(err)
	}

	// 迁移对象到数据表上
	err = Db.AutoMigrate(&User{}, &Video{}, &Comment{}, &UserLogin{})
	if err != nil {
		log.Panicln(err)
	}
}
