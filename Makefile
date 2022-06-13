.PHONY: gen-openapi
gen-openapi:
	oapi-codegen -config api-docs/models.config.yaml api-docs/api-schema.yaml
	oapi-codegen -config api-docs/chi-server.config.yaml api-docs/api-schema.yaml
