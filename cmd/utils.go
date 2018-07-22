package cmd

import (
	"os"
)

// GetEnvString defines a environment variable with a specified name, fallback value.
// The return is a string value.
func GetEnvString(key, fallback string) string {
	if s := os.Getenv(key); s != "" {
		return s
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
