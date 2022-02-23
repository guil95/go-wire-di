package usecase

import "github.com/google/wire"

var Provider = wire.NewSet(NewUseCase)
