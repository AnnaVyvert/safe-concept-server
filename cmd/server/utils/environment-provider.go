package utils

import (
	"fmt"
	"os"
)

func noEnvPanic(env string) {
	panic(fmt.Sprintf("environment_provider: %v env reading", env))
}

func GetEnvDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

func GetEnv(name string) string {
	env := os.Getenv(name)
	if env == "" {
		noEnvPanic(name)
	}

	return env
}
