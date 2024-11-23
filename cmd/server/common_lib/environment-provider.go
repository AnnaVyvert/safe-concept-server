package common_lib

import (
	"fmt"
	"os"
)

func noEnvPanic(env string) {
	panic(fmt.Sprintf("environment_provider: %v env reading", env))
}

func GetEnv(name string) string {
	env := os.Getenv(name)
	if env == "" {
		noEnvPanic(name)
		return ""
	}

	return env
}
