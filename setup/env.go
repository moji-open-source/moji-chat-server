package setup

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Address     string
	DatabaseDsn string `mapstructure:"DATABASE_DSN"`
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
