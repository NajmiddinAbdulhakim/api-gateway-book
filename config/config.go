package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	
	BookServiceHost string
	BookServicePort int

	HTTPPort string
}

func Load() Config {
	c := Config{}

	c.HTTPPort = cast.ToString(look(`HTTP_PORT`, `:9999`))

	c.BookServiceHost = cast.ToString(look(`BOOK_SERVICE_HOST`, `127.0.0.1`))
	c.BookServicePort = cast.ToInt(look(`BOOK_SERVICE_PORT`, 9000))


	return c
}

func look(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
