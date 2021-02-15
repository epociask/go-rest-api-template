BIN_NAME = rest-api-template

wire:
	@echo building wire....
	@wire

build:
	@echo building binary...
	@GOPRIVATE=github.com/epociask CGO_ENABLED=0 go build -a -tags netgo -o bin/$(BIN_NAME);

run: 
	@./bin/${BIN_NAME}

test:
	@ go test ./...
