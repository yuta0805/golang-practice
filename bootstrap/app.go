package bootstrap

import "gorm.io/gorm"

type App struct {
	Env *Env
	db *gorm.DB
}

func NewApp() *App {
	env := NewEnv()
	db := NewDB(
		env.DB_HOST,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_DATABASE,
		env.DB_PORT,
		env.DB_DRIVER,
	)

	return &App{
		Env: env,
		db: db,
	}
}
