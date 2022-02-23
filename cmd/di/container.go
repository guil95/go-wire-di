package di

import (
	"github.com/google/wire"

	"github.com/guil95/go-wire-di/internal/repository"
	"github.com/guil95/go-wire-di/internal/usecase"
	"github.com/guil95/go-wire-di/pkg/database"
)

var Container = wire.NewSet(
	database.Provider,
	repository.Provider,
	usecase.Provider,
	wire.Struct(new(App), "*"),
)
