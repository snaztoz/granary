OUT := gran

build:
	go build -o $(OUT) cmd/gran/*.go

test:
	go test ./...

clean:
	rm $(OUT) *.gran

.PHONY: build test clean
