package config

import (
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
			Debug:                        true,
			APIAddress:                   "https://moj.gov.pl/nforms/api/UprawnieniaKierowcow/2.0.10/data/driver-permissions?hashDanychWyszukiwania=",
		}
	})
	return instance
}
