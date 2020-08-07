NIFI_VERSION ?= 1.11.4

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
	-p packageVersion=${NIFI_VERSION},withGoCodegenComment=true,structPrefix=true \
	--skip-validate-spec \
	--git-host github.com \
	--git-user-id simingweng \
	--git-repo-id nifi-go-client \
	-o /local

test-with-nifi:
	docker run --rm -d --health-cmd "curl -f http://`hostname`:8080/nifi || exit 1" --name nifi -p 8080:8080 apache/nifi:${NIFI_VERSION}
	until [ `docker inspect -f {{.State.Health.Status}} nifi` == "healthy" ]; do sleep 5; done
	go test
	docker container stop nifi