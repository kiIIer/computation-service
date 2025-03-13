package main

import (
	"computation-service/internal"
	"github.com/google/wire"
)

func main() {
	wire.Build(internal.InitializeApp)
	select {} // Keep running
}
