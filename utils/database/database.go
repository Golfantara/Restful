package database

import (
	"fmt"
	"test/configs"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config configs.ProgramConfig) *gorm.DB {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		logrus.Error("model : cannot connect to database, ", err.Error())
		return nil
	}
	return db
}