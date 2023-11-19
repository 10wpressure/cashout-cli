# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Output binary names
BINARY_NAME=sign
BINARY_MAC=$(BINARY_NAME)_mac
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_WIN=$(BINARY_NAME)_win.exe

all: test build
.PHONY: all

run: build
	@./bin/$(BINARY_NAME)
.PHONY: run

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v
.PHONY: build

test:
	$(GOTEST) -v ./...
.PHONY: test

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_MAC)
	rm -f $(BINARY_LINUX)
	rm -f $(BINARY_WIN)
.PHONY: clean

# Build for macOS
build-mac:
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_MAC) -v

# Build for Linux
build-linux:
	@GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_LINUX) -v

# Build for Windows
build-win:
	@GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_WIN) -v