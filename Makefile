.PHONY: tests
tests:
	go mod tidy
	go test ./...
