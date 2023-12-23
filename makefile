.PHONY: all
all: vet lint test

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -coverprofile=cov.out ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint:
	which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

.PHONY: coverage
coverage: test
	go tool cover -html=cov.out -o=cov.html
