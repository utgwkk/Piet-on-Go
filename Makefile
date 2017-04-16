pietongo: src/*/*.go
	GOPATH=$(PWD) go build -o $@ main

.PHONY: test

test:
	GOPATH=$(PWD) go test piet
