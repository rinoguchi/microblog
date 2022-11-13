.PHONY: go-install
go-install:
	go get
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	go install github.com/google/wire/cmd/wire@latest
	go install xorm.io/reverse@v0.1.2
	go install golang.org/x/tools/cmd/goimports@latest
	asdf reshim golang

.PHONY: gen-openapi
gen-openapi:
	oapi-codegen -config _docs/models.config.yaml _docs/api-schema.yaml
	oapi-codegen -old-config-style -templates _docs/oapi_codegen_templates/ -generate chi-server,spec,skip-prune -package controllers -o adapters/controllers/server.gen.go _docs/api-schema.yaml

.PHONY: gen-wire
gen-wire:
	wire

.PHONY: gen-db-model
gen-db-model:
	cat _docs/xorm_reverse.config.yaml | envsubst > _docs/xorm_reverse.config.gen.yaml && reverse -f _docs/xorm_reverse.config.gen.yaml

.PHONY: gen-db-model-mapper
gen-db-model-mapper:
	cd adapters/repositories/models/gen && go generate

.PHONY: gen-controller-model-mapper
gen-controller-model-mapper:
	cd adapters/controllers/models/gen && go generate

.PHONY: gen-repository-model-mapper
gen-repository-model-mapper:
	cd adapters/repositories/models/gen && go generate

.PHONY: gen-usecase-models
gen-usecase-models:
	cd usecases/models/gen && go generate

.PHONY: initialize
initialize:
	make gen-openapi
	make gen-wire
	make gen-controller-model-mapper
	make gen-usecase-models
	make gen-db-model
	make gen-db-model-mapper

.PHONY: serve
serve:
	go run .


.PHONY: deploy
deploy:
	make initialize
	gcloud app deploy

