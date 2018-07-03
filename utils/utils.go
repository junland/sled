package utils

import (
	"os"
)

// GetEnvString defines a environment variable with a specified name, fallback value.
// The return is a string value.
func GetEnvString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetEnvBool defines a environment variable with a specified name, fallback value.
// The return is either a true or false.
func GetEnvBool(key string, fallback bool) bool {
	switch os.Getenv(key) {
	case "true":
		return true
	case "false":
		return false
	default:
		return fallback
	}
}
