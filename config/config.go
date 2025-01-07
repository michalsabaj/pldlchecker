package config

import (
	"fmt"
	"sync"
)

type Config struct {
	NameCharLimit                int
	DriverLicenseNumberCharLimit int
	Debug                        bool
	APIAddress                   string
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig initializes or retrieves the singleton instance of the configuration
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			NameCharLimit:                80,
			DriverLicenseNumberCharLimit: 8,
			Debug:                        false,
			APIAddress:                   "https://moj.gov.pl/nforms/api/UprawnieniaKierowcow/2.0.10/data/driver-permissions?hashDanychWyszukiwania=",
		}
	})
	return instance
}

// In case of a need to change the configuration, the following methods are available:
func (c *Config) SetNameCharLimit(limit int) error {
	if limit <= 0 {
		return fmt.Errorf("name character limit must be positive")
	}
	c.NameCharLimit = limit
	return nil
}

func (c *Config) SetDriverLicenseNumberCharLimit(limit int) {
	c.DriverLicenseNumberCharLimit = limit
}

func (c *Config) SetDebug(debug bool) {
	c.Debug = debug
}

func (c *Config) SetAPIAddress(address string) error {
	if address == "" {
		return fmt.Errorf("API address cannot be empty")
	}
	c.APIAddress = address
	return nil
}
