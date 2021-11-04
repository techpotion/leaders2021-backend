package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/techpotion/leaders2021-backend/modules/controllers/marks"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"github.com/techpotion/leaders2021-backend/modules/server"
)

func init() {
	// viper config initalization
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln()
	}
	logrus.Info("viper config initialized successfully")

	// using non build-it Init instead built-in
	// because viper config is required but
	// main.init is being called after the
	// database.init function
	database.Init()
	marks.Init()
}

func main() {
	server.Run()
}
