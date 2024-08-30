OUT := gran
INSTALL_DIR := $(shell go env GOBIN)

build:
	go build -o $(OUT) ./cmd/gran

install: build
	mv $(OUT) $(INSTALL_DIR)

test:
	go test ./...

clean:
	rm -f $(OUT) *.gran

.PHONY: build install test clean
