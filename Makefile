
VERSION = $(shell printf "%s.%s" \
		$$(git rev-list --count HEAD) \
		$$(git rev-parse --short HEAD))

install:
	@echo :: getting dep
	go get -v -d

test: install
	@echo :: run tests
	go test -v ./tests

run: install
	@echo :: start the http server
	go run main.go

build:  $(OUTPUT)
	CGO_ENABLED=0 GOOS=linux go build -o bin/app \
		-ldflags "-X main.version=$(VERSION)" \
		-gcflags "-trimpath $(GOPATH)/src"

all: install build tests run


$(OUTPUT):
	mkdir -p $(OUTPUT)