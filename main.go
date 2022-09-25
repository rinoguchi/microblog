package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rinoguchi/microblog/adapters/controllers"
)

func main() {
	println("main started")
	swagger, err := controllers.GetSwagger() // APIスキーマ定義を取得
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	server := InitializeServer()
	router := chi.NewRouter()

	allowOrigins := strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		MaxAge:         300,
	}))

	router.Use(middleware.OapiRequestValidator(swagger)) // validationを設定
	router.Use(server.SetTxMiddleware)                   // トランザクションをcontextに設定

	controllers.HandlerFromMux(server, router) // chiのrouterと実装したserverを紐付け
	port := ":8080"
	println("starting server port" + port)
	http.ListenAndServe(port, router)
}
