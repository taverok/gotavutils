package osutl

import (
	"fmt"
	"os"
)

func EnvOrDefault(app, env, defaultValue string) string {
	val, ok := os.LookupEnv(fmt.Sprintf("%s_%s", app, env))
	if ok {
		return val
	}

	return defaultValue
}
