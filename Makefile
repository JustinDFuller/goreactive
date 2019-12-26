.PHONY: test
test: 
	@go vet ./...;
	@go test -v -race -vet=off ./...;
	@gofmt -l -w -s *.go ./**/*.go;

.PHONY: bench
bench: test
	@go test -v -bench=. ./...;

