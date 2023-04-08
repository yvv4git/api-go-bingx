format_default:
	gofmt -w -s .
	goimports -l -w .
	gofumpt -l -w .
	gci write .

format_by_linter:
	golangci-lint run . -v --fix

lint:
	golangci-lint run . -v

sec:
	semgrep --config=auto .
