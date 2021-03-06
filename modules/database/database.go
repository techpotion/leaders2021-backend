package database

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/techpotion/leaders2021-backend/modules/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connectionString string // connection string
var db *gorm.DB             // database singleton instance

// Init initializes connection string as a global module variable
// We use Init instead of self executed init as we need
// some other stuff (viper) to be configured
// before  database.init executiong
func Init() {
	connectionString = makeConnectionString()
	var dbErr error
	db, dbErr = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if dbErr != nil {
		logrus.WithFields(logrus.Fields{
			"error": dbErr.Error(),
		}).Fatalln("failed to connect to the database")
	}

}

// New returns database instance
func New() (*gorm.DB, error) {
	return db, nil
}

// makeConnectionString returns new connection string
func makeConnectionString() string {
	var err error

	host := ""
	if host, err = utils.GetEnvStr("DB_HOST"); err != nil {
		host = viper.GetString("database.host")
	}

	port := 0
	if port, err = utils.GetEnvInt("DB_PORT"); err != nil {
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
