MAKEFLAGS += --always-make

all: generate fmt lint test tidy

# run tests
test:
	go test -race -shuffle=on ./...

# format source code
fmt:
	golangci-lint fmt

# lint source code
lint:
	golangci-lint run --tests=false

# run code generation
generate:
	go generate ./...

# tidy up go.mod
tidy:
	go mod tidy -v
