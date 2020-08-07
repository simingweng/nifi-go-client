all: generate fmt vet

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v4.3.1 \
	generate \
	-i /local/api/swagger.json \
	-g go \
	--package-name nifi \
	-p packageVersion=1.11.4,withGoCodegenComment=true,structPrefix=true \
	--skip-validate-spec \
	--git-host github.com \
	--git-user-id simingweng \
	--git-repo-id nifi-go-client \
	-o /local
