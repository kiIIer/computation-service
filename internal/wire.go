package internal

import (
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.Struct(new(App), "*"),
)

func InitializeApp() *App {
	return &App{}
}

type App struct {
}
