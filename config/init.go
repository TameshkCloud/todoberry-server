package config

import "fmt"
import "github.com/spf13/viper"

var ConfFilesPath string

// Init configuration
func Init() {
	fmt.Println("loading configurations...")
	LoadConfigurations()
}

func LoadConfigurations() {
	InitRedisConfig()
}

func LoadConfig(configPath string, configStruct interface{}) {
	viperInstance := viper.New()
	fullConfigPath := ""
	
	fullConfigPath = ConfFilesPath + "config/conf/" + configPath + ".toml"
	viperInstance.SetConfigFile(fullConfigPath)
	err := viperInstance.ReadInConfig()
	if err != nil {
		fmt.Println("failed to read from file", "config_path", fullConfigPath, "error", err)
	}
	err = viperInstance.Unmarshal(configStruct)
	if err != nil {
		fmt.Println("failed to unmarshal", "config_path", fullConfigPath, "error", err)
	}
	fmt.Println("config loaded successfully", "config_path", fullConfigPath)
}