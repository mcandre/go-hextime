all: test

test:
	hextime

gofmt:
	gofmt -s -w .

lint: gofmt
