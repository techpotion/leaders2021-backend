package database

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/techpotion/leaders2021-backend/modules/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func getConnectionString() string {
	var err error

	host := ""
	if host, err = utils.GetEnvStr("DB_HOST"); err != nil {
		host = viper.GetString("database.host")
	}

	port := 5442
	if port, err = utils.GetEnvInt("DB_HOST"); err != nil {
		port = viper.GetInt("database.port")
	}

	username := ""
	if username, err = utils.GetEnvStr("DB_HOST"); err != nil {
		username = viper.GetString("database.username")
	}

	password := ""
	if password, err = utils.GetEnvStr("DB_HOST"); err != nil {
		password = viper.GetString("database.password")
	}

	db := ""
	if db, err = utils.GetEnvStr("DB_HOST"); err != nil {
		db = viper.GetString("database.db")
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		host,
		username,
		password,
		db,
		port,
	)
}
