package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func InitDB() {
	DB = openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func openDB(username, password, addr, name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	newLogger := logger.New(
		log.StandardLogger(),
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Error(err)
		log.Errorf("Database connection failed. Database name: %s", name)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error(err)
		log.Errorf("SetMaxIdleConns get an error.")
	} else {
		sqlDB.SetMaxOpenConns(20000)
		sqlDB.SetMaxIdleConns(100)
	}

	log.Info("db connected")

	return db
}
