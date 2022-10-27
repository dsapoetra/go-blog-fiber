.PHONY: all
all: clean build

APP=gbf
APP_EXECUTABLE="./bin/$(APP)"
define DOCKERCOMMAND
docker run -d --name gbf \
-p 8080:5005 \
gbf:latest
endef

clean:
	rm -f $(APP_EXECUTABLE)
	go mod tidy
	go mod vendor

compile:
	mkdir -p bin/
	go build -o $(APP_EXECUTABLE)

build: clean compile

build-docker:
	docker build -f Dockerfile -t $(APP):latest .

run-docker:
	$(DOCKERCOMMAND)

clean-docker:
	docker stop $(APP) && docker rm -f $(APP)

lint:
	golangci-lint run --out-format checkstyle > lint.xml

mock-prepare:
	go install github.com/golang/mock/mockgen@v1.6.0

mock:
	@./scripts/mockgen.sh service

swag:
	swag init -o ./docs

run:
	go run main.go

create-migration :
	goose -dir platform/migrations create $(input) sql