// Package config
// Created by RTT.
// Author: teocci@yandex.com on 2025-1월-14
package config

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type APIServer struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol,omitempty"`
}

type WebServer struct {
	Port int `json:"port"`
}

type ProfileData struct {
	API APIServer `json:"api"`
}

type ServerSetup struct {
	Profile  string                 `json:"profile"`
	Web      WebServer              `json:"web"`
	Profiles map[string]ProfileData `json:"profiles"`
	Config   string                 `json:"-"`
}

const (
	defaultConfigPath = "./config.json"
	defaultProfile    = "prod"
	defaultWebPort    = 3000
)

var (
	configInstance *ServerSetup
	once           sync.Once
	mutex          sync.RWMutex

	Config  string
	Profile string
	Port    int
)

// AddFlags adds the command-line flags for configuration
func AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&Config, "config", "c", defaultConfigPath, "Path to the configuration file")
	cmd.Flags().StringVarP(&Profile, "profile", "r", defaultProfile, "Set the configuration profile (e.g., dev, prod)")
	cmd.Flags().IntVarP(&Port, "port", "p", defaultWebPort, "Override the default server port")
}

// LoadConfigFile loads the configuration from a specified file
func LoadConfigFile(file string) error {
	v := viper.New()
	v.SetConfigFile(file)

	// Set default values
	v.SetDefault("profile", defaultProfile)
	v.SetDefault("web.port", defaultWebPort)

	// Read in the configuration file
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Override defaults with environment variables
	v.AutomaticEnv()

	// Unmarshal into ServerSetup struct
	var config ServerSetup
	if err := v.Unmarshal(&config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Lock and update the global config instance
	mutex.Lock()
	defer mutex.Unlock()

	configInstance = &config

	return nil
}

func Load() error {
	if Config == "" {
		return errors.New("no configuration file specified")
	}

	if err := LoadConfigFile(Config); err != nil {
		return fmt.Errorf("failed to load configuration file: %w", err)
	}

	// Merge CLI params (if set)
	mutex.Lock()
	defer mutex.Unlock()
	if configInstance != nil {
		if configInstance.Web.Port > 0 {
			configInstance.Web.Port = Port
		}
		if configInstance.Profile != "" {
			configInstance.Profile = Profile
		}
	}

	return nil
}

// Fetch retrieves the configuration, loading it if necessary
func Fetch() (*ServerSetup, error) {
	var err error

	once.Do(func() {
		err = Load()
	})

	return configInstance, err
}

// Watch watches the config file for changes and reloads it dynamically
func Watch() {
	v := viper.New()
	v.SetConfigFile(Config)

	// Watch for changes
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		if err := Load(); err != nil {
			log.Printf("Failed to reload configuration: %v", err)
		} else {
			log.Println("Configuration successfully reloaded")
		}
	})

	// Load initial config
	if err := Load(); err != nil {
		log.Printf("Error reading initial config: %v", err)
	}
}

// Get safely retrieves the current configuration instance
func Get() *ServerSetup {
	mutex.RLock()
	defer mutex.RUnlock()

	return configInstance
}
