package database

import (
	"fmt"

	"github.com/smochii/go-clean-architecture/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Conf.Db.Host, config.Conf.Db.User, config.Conf.Db.Password, config.Conf.Db.Name, config.Conf.Db.Port)
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if error != nil {
		panic(error)
	}
	return db
}
