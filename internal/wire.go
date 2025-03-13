package internal

import (
	"computation-service/internal/db"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	db.NewDB,
	wire.Struct(new(App), "*"),
)

func InitializeApp(db *db.DB) *App {
	return &App{DB: db}
}

type App struct {
	DB *db.DB
}
