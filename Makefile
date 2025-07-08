.PHONY: test
test:
	go test

.PHONY: readme
readme:
	gomarkdoc --output README.md

.PHONY: version
version:
	@go version

.PHONY: generate
generate: version
	go generate ./...

.PHONY: lint
lint: version
	staticcheck -checks all,-SA1019,-ST1000 ./...
	go vet ./...

.PHONY: tools
tools:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest

all: generate test tools lint readme
