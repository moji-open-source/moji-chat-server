package setup

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Env      *Env
	Database *gorm.DB
	Redis    *redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = LoadEnv()
	app.Database = NewDatabase(app.Env)
	app.Redis = NewRedlisClient(app.Env)
	return *app
}
