// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/rinoguchi/microblog/adapters/controllers"
	"github.com/rinoguchi/microblog/adapters/repositories"
	"github.com/rinoguchi/microblog/usecases"
)

// Injectors from wire.go:

func InitializeServer() *controllers.Server {
	db := repositories.NewDB()
	commentRepository := repositories.NewCommentRepositoryImpl()
	commentUsecase := usecases.NewCommentUsecase(commentRepository)
	server := controllers.NewServer(db, commentUsecase)
	return server
}
