OS := $(shell uname -s 2>/dev/null || echo Windows)

run: 
	GIN_MODE=release ./bin/api
build:
	go build -o ./bin/api ./cmd/api
test:
	go test -v ./...
watch:
ifeq (${OS},Windows)
	@air -c .air.win.toml
else
	@air -c .air.toml
endif
