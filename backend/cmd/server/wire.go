//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	apiv1 "buf-protoc/api/v1"
	"buf-protoc/backend/server"

	"github.com/google/wire"
)

func wireApp() (*server.Server, func(), error) {
	panic(wire.Build(apiv1.WireSet, newApp))
}
