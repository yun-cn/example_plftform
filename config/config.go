package config

import (
	"fmt"
	"os"

	mapset "github.com/deckarep/golang-set"

	"github.com/spf13/viper"
)

func assertAppEnv() {
	allowed := mapset.NewSetFromSlice([]interface{}{"", "test", "development", "production"})
	appEnv := os.Getenv("SPACEMARKET_CRAWLER_ENV")

	if !allowed.Contains(appEnv) {
		panic(fmt.Errorf("SPACEMARKET_CRAWLER_ENV must be one of (test | development | production) but was %s", appEnv))
	}
}

// Setup Sets SPACEMARKET_CRAWLER_ENV from environment or default to development
func Setup() {
	viper.SetDefault("env", "development")
	assertAppEnv()
	viper.SetEnvPrefix("spacemarket_crawler")
	viper.BindEnv("env")
	viper.AddConfigPath("./configs")
	viper.SetConfigName(viper.GetString("env"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error: %s", err))
	}
}
