.PHONY: gen-openapi
gen-openapi:
	oapi-codegen -config _docs/models.config.yaml _docs/api-schema.yaml
	oapi-codegen -config _docs/chi-server.config.yaml _docs/api-schema.yaml

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

