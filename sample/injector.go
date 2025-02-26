//go:build wireinject
// +build wireinject

package sample

import (
	"github.com/google/wire"
)

func InitializedService() *SimpleService {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil
}
