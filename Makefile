PKG_LIST := $(shell go list ./... | grep -v /vendor/)

default: build

all: test build

tidy:  ## Execute tidy comand
	@go mod tidy

build: tidy ## Build the binary file
	go build -i -v $(PKG_LIST)

docker-build:
	docker build -t operation-quasar-fire .

docker-run:
	docker run -d -p 5000:5000 operation-quasar-fire

start-infra:
	docker compose -f resources/docker-compose/docker-compose.yml --env-file .env up -d

stop-infra:
	docker compose -f resources/docker-compose/docker-compose.yml stop

fmt: ## Formmat src code files
	@go fmt ${PKG_LIST}

test: ## Execute test
	@echo "go test"
	go test ./... --coverprofile=cover.out

cover: test
	@echo "generating cover..."
	@go tool cover -html=cover.out

race: ## Run data race detector
	@go test -race -short ${PKG_LIST}
