package setup

type Application struct {
	Env *Env
}

func App() Application {
	app := &Application{}
	app.Env = LoadEnv()
	return *app
}
