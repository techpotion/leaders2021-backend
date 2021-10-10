package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/techpotion/leaders2021-backend/modules/server"
)

func main() {
	// viper init
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln()
	}
	logrus.Info("viper config initialized successfully")

	// server starting
	server.Run()
}
