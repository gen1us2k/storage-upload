init:
	rm -rf bin/*
	cd tools && go generate -x -tags=tools

check:
	./bin/golangci-lint run ./...
