package config

import (
	"fmt"
	// "strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type CadenceConfig struct {
	Domain   string
	Service  string
	HostPort string
}

type AppConfig struct {
	Env            string
	WorkerTaskList string
	Cadence        CadenceConfig
	Logger         *zap.Logger
}

func (h *AppConfig) Setup() {
	viper.SetConfigName("application")
	viper.AddConfigPath("app/resources")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error in Reading file, %s", err)
	}
	err := viper.Unmarshal(&h)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	h.Logger = logger

	logger.Debug("Finished loading Configuration!")
}
