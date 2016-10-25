all: test

test:
	hextime

govet:
	go vet -v

gofmt:
	gofmt -s -w .

lint: govet gofmt
