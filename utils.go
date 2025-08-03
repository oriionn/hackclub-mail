package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/ini.v1"
)

type Config struct {
	api_key string
}

func GetConfigPath() string {
	configFileName := "config.ini"
	var configDir string

	switch runtime.GOOS {
	case "windows":
		// Windows: %APPDATA%\hackclub-mail\
		configDir = filepath.Join(os.Getenv("APPDATA"), "hackclub-mail")
	case "darwin":
		// macOS: ~/Library/Application Support/hackclub-mail/
		homeDir, _ := os.UserHomeDir()
		configDir = filepath.Join(homeDir, "Library", "Application Support", "hackclub-mail")
	default:
		// Linux/Unix: ~/.config/hackclub-mail/
		homeDir, _ := os.UserHomeDir()
		configDir = filepath.Join(homeDir, ".config", "hackclub-mail")
	}

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}

	return filepath.Join(configDir, configFileName)
}

func ReadConfig() (Config, error) {
	var config Config
	configPath := GetConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return config, errors.New("Please create and edit 'config.ini' in this directory : '" + filepath.Dir(configPath) + "'")
	}

	cfg, err := ini.Load(configPath)
	if err != nil {
		return config, err
	}

	config.api_key = cfg.Section("general").Key("api_key").String()

	if (config.api_key == "") {
		return config, errors.New("Please edit '" + configPath + "' with your api key...")
	}

	return config, nil
}

func WriteConfig(config Config) error {
	configPath := GetConfigPath()

	cfg := ini.Empty()
	generalSection, err := cfg.NewSection("general")
	if err != nil {
		return err
	}

	_, err = generalSection.NewKey("api_key", config.api_key)
	if err != nil {
		return err
	}

	return cfg.SaveTo(configPath)
}
