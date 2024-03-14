install-wire:
ifeq (, $(shell which wire))
	@echo "wire not found, installing..."
	go install github.com/google/wire/cmd/wire@latest
endif
	@echo "wire has been installed!"

init: install-wire

wire: install-wire
	@wire gen ./internal/app

dev:
	APP_ENV=dev go run ./cmd/main.go

test:
	go test ./... -v -cover
