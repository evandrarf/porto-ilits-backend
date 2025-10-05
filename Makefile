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

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run       Run the application"
	@echo "  build     Build the application"
	@echo "  watch     Run the application with hot reload"
	@echo "  help      Display this help message"

.DEFAULT_GOAL := help