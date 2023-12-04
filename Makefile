FILES = $(shell find . -type f -name '*.go')

init:
	rm -rf bin/*
	cd tools && go generate -x -tags=tools

check:
	./bin/golangci-lint run ./...
gen:
	go generate ./...
	make format

format:                 ## Format source code
	bin/gofumpt -l -w $(FILES)
	bin/goimports -local storage -l -w $(FILES)
	bin/gci write --section Standard --section Default --section "Prefix(storage)" $(FILES)
test:
	go test -v ./...
