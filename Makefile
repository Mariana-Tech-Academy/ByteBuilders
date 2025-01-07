run: |
	gofmt -w .
	go run main.go


lint:
	golangci-lint run ./... --timeout=2m -D staticcheck,govet