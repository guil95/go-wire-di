//+build wireinject

package di

import (
	"github.com/google/wire"
)

func SetupApp() *App {
	panic(wire.Build(Container))
}
