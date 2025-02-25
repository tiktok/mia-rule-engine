.PHONY: fmt lint antlr test bench

fmt:
	gofmt -w ./
	go mod tidy

lint:
	golangci-lint run ./... \
		--timeout=5m \
		--skip-files=.*_test\.go$ \
    	--issues-exit-code=1 \
      	--max-issues-per-linter=0 \
      	--max-same-issues=0 \
      	--enable-all \
      	--exclude-use-default=false \
      	--concurrency=5

antlr:
	antlr -Dlanguage=Go parser/cmpl.g4

test:
	go test ./... \
		--timeout=5m \
		-gcflags="all=-N -l" \
		-count=1 \
		-covermode=count \
		-coverprofile=.coverage.out

bench:
	go test -bench=. -benchmem -run=^$$ ./benchmark > .benchmark.out


