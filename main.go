package main

import (
	"fmt"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rinoguchi/microblog/adapters/controllers"
)

func main() {
	swagger, err := controllers.GetSwagger() // APIスキーマ定義を取得
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	server := InitializeServer()
	router := chi.NewRouter()
	router.Use(middleware.OapiRequestValidator(swagger)) // validationを設定
	controllers.HandlerFromMux(server, router)           // chiのrouterと実装したserverを紐付け
	http.ListenAndServe(":8080", router)                 // 8080ポートをリッスン
}
