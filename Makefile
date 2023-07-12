
format:
		goimports -w -l .
		gofmt -w .
		gofumpt -w .

license-check:
	# go install github.com/lsm-dev/license-header-checker/cmd/license-header-checker@latest
	license-header-checker -v -a -r apache-license.txt . go

check: license-check
		golangci-lint run

test:
		go test -v ./... -coverprofile=coverage.txt -covermode=atomic

build: format test