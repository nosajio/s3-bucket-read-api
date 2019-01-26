package main

import (
	"os"
)

// EnvDefault allows getting environment variables with a fallback
// name: The name of the env var
// fallback: The default value when the variable is undefined
func EnvDefault(name string, fallback interface{}) interface{} {
	value, exists := os.LookupEnv(name)
	if !exists {
		return fallback
	}
	return value
}
