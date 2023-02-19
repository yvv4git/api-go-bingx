format:
	gofmt -w -s .
	goimports -l -w .
	gofumpt -l -w .
	gci write .

lint:
	golangci-lint run . -v

sec:
	semgrep --config=auto .
