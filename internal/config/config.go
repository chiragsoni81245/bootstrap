package config

import (
	"github.com/spf13/viper"
)

type Config map[string]any

func GetConfig(configPath string) (*Config, error){
	// Set the path to the configuration file
	viper.SetConfigFile(configPath)

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
        return nil, err
	}

    var config Config
    config = viper.AllSettings()

    return &config, nil
}
