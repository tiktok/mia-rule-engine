format:
	gofmt -w ./
	go mod tidy

build:
	go build ./...

gen-antlr:
	antlr -Dlanguage=Go parser/cmpl.g4

test:
	IS_UNIT_TEST_ENV=1 go test ./... -gcflags="all=-N -l" -count=1 -covermode=count -coverprofile=.coverage.out
