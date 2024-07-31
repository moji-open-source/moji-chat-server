package setup

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Address     string
	DatabaseDsn string `mapstructure:"DATABASE_DSN"`

	RedisAddr       string `mapstructure:"RDB_ADDR"`
	RedisPassword   string `mapstructure:"RDB_PASSWORD"`
	RedisDatabase   int    `mapstructure:"RDB_DB"`
	RedisUsername   string `mapstructure:"RDB_USERNAME"`
	RedisPollSize   int    `mapstructure:"RDB_POOL_SIZE"`
	RedisClientName string `mapstructure:"RDB_CLIENT_NAME"`
}

func LoadEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
