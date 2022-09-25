//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rinoguchi/microblog/adapters/controllers"
	"github.com/rinoguchi/microblog/adapters/repositories"
	"github.com/rinoguchi/microblog/usecases"
)

func InitializeServer() *controllers.Server {
	wire.Build(
		controllers.NewServer,
		repositories.NewDB,
		usecases.NewCommentUsecase,
		repositories.NewCommentRepositoryImpl,
	)
	return &controllers.Server{}
}
