.PHONY: fmt antlr test

fmt:
	gofmt -w ./
	go mod tidy

antlr:
	antlr -Dlanguage=Go parser/cmpl.g4

test:
	go test ./... -gcflags="all=-N -l" -count=1 -covermode=count -coverprofile=.coverage.out
