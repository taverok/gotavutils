package osutl

import (
	"fmt"
	"os"
)

func EnvOrDefault(envPrefix, env, defaultValue string) string {
	val, ok := os.LookupEnv(fmt.Sprintf("%s_%s", envPrefix, env))
	if ok {
		return val
	}

	return defaultValue
}
