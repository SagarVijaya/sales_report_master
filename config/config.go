package config

import (
	"log"
	"salesproject/apps/models"

	"github.com/BurntSushi/toml"
)

var configData *models.Config

// LoadGlobalConfig loads the config from a TOML file if not already loaded
func LoadGlobalConfig(filename string) {
	if configData != nil {
		log.Println("Config already loaded. Skipping reload.")
		return
	}

	var cfg models.Config
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}
	configData = &cfg
	log.Println("TOML config loaded successfully")
}

// GetConfig returns the already loaded config
func GetConfig() *models.Config {
	if configData == nil {
		log.Fatal("Config not loaded! Call LoadGlobalConfig first.")
	}
	return configData
}
