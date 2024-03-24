package env

import (
	"fmt"
	"os"
)

func GetOrPanicOnEmpty(env string) string {
	e := os.Getenv(env)
	if e == "" {
		panic(fmt.Sprintf("environment variable '%s' cannot be empty", env))
	}
	return e
}
