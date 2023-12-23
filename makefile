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

lint:
	which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run --skip-dirs scripts

coverage: test
	go tool cover -html=cov.out -o=cov.html
