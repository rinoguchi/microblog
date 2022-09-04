// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/rinoguchi/microblog/adapters/controllers"
	"github.com/rinoguchi/microblog/adapters/repositories"
	"github.com/rinoguchi/microblog/usecases"
)

// Injectors from wire.go:

func InitializeServer() *controllers.Server {
	commentRepository := repositories.NewCommentRepositoryImpl()
	commentUsecase := usecases.NewCommentUsecase(commentRepository)
	server := controllers.NewServer(commentUsecase)
	return server
}
