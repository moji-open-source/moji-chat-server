package setup

import "gorm.io/gorm"

type Application struct {
	Env      *Env
	Database *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = LoadEnv()
	app.Database = NewDatabase(app.Env)
	return *app
}
