package config

import "os"

type TConfig struct{}

func Config() TConfig {
	return TConfig{}
}

func get(key string) string {
	return os.Getenv(key)
}
