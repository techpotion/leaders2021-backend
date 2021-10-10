package database

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/techpotion/leaders2021-backend/modules/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connectionString string

func Init() {
	fmt.Println("db init!")
	connectionString = makeConnectionString()
}

func New() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}

func makeConnectionString() string {
	var err error

	host := ""
	if host, err = utils.GetEnvStr("DB_HOST"); err != nil {
		fmt.Println("here 1")
		host = viper.GetString("database.host")
		fmt.Println("host: ", host)
	}

	port := 0
	if port, err = utils.GetEnvInt("DB_PORT"); err != nil {
		fmt.Println("here 2")
		port = viper.GetInt("database.port")
	}

	username := ""
	if username, err = utils.GetEnvStr("DB_USERNAME"); err != nil {
		username = viper.GetString("database.username")
	}

	password := ""
	if password, err = utils.GetEnvStr("DB_PASSWORD"); err != nil {
		password = viper.GetString("database.password")
	}

	db := ""
	if db, err = utils.GetEnvStr("DB_DB"); err != nil {
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
