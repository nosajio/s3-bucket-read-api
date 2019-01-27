package main

import (
	"fmt"
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

// RelativeURLPath returns a path appended to the SELF_HOST env var
func RelativeURLPath(path string) string {
	selfURL := EnvDefault("SELF_HOST", "")
	if selfURL == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", selfURL, path)
}
