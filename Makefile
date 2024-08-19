.PHONY: all clean test build
GOLANGCI_VERSION=v1.54.2

lint:
	golangci-lint run
build:
	go build basic.go
# Install golangci-lint	
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GO_PATH}/bin ${GOLANGCI_VERSION}
test:
	scripts/test.sh