GO         := go
GOFMT      := gofmt
GOPATH     := $(shell $(GO) env GOPATH)
GOHOSTOS   := $(shell $(GO) env GOHOSTOS)
GOHOSTARCH := $(shell $(GO) env GOHOSTARCH)

BINARY := server

STATICCHECK := $(GOPATH)/bin/staticcheck
$(STATICCHECK):
	$(GO) install honnef.co/go/tools/cmd/staticcheck@v0.4.6

SWAG := $(GOPATH)/bin/swag
$(SWAG):
	$(GO) install github.com/swaggo/swag/cmd/swag@latest

.PHONY: check
check: $(STATICCHECK) $(GOLANGCI)
	$(GOFMT) -s -w ./
	$(GO) vet ./...
	$(STATICCHECK) ./build/... ./cmd/... ./internal/... ./test/...
	$(GO) mod tidy

.PHONY: test
test:
	$(GO) test -v ./...

.PHONY: swag
swag: $(SWAG)
	$(SWAG) init -g cmd/main.go

.PHONY: build
build:
	$(GO) mod download
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) $(GO) build -o $(BINARY) ./cmd

.PHONY: run
run:
	./$(BINARY) --sqlite-path app.sqlite --docs
