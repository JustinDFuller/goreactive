.PHONY: test
test: 
	@gofmt -l -w -s *.go ./**/*.go;
	@go vet ./...;
	@go test -v -race -vet=off ./...;

.PHONY: bench
bench: test
	@go test -v -bench=. ./...;

