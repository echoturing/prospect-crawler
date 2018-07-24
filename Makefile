main:
	GOBIN=$(shell pwd)/bin vgo install -v ./cmd/...

clean:
	go clean -cache
