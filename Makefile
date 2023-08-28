ARCH := amd64
BINARY_NAME := invalidator

linux:
	GOOS=linux GOARCH=$(ARCH) CGO_ENABLED=0 go build -v -o ./bin/$(BINARY_NAME)_$(ARCH)_linux ./...
zip-linux:
	cp ./bin/$(BINARY_NAME)_$(ARCH)_linux ./bin/$(BINARY_NAME) 
	zip ./bin/$(BINARY_NAME)_$(ARCH)_linux.zip ./bin/$(BINARY_NAME)
	rm -rf ./bin/$(BINARY_NAME)


windows:
	GOOS=windows GOARCH=$(ARCH) CGO_ENABLED=0 go build -v -o ./bin/$(BINARY_NAME)_$(ARCH)_windows ./...
zip-windows:
	cp ./bin/$(BINARY_NAME)_$(ARCH)_windows ./bin/$(BINARY_NAME).exe
	zip ./bin/$(BINARY_NAME)_$(ARCH)_windows.zip ./bin/$(BINARY_NAME).exe
	rm -rf ./bin/$(BINARY_NAME).exe

macos:
	GOOS=darwin GOARCH=$(ARCH) CGO_ENABLED=0 go build -v -o ./bin/$(BINARY_NAME)_$(ARCH)_macos ./...
zip-macos:
	cp ./bin/$(BINARY_NAME)_$(ARCH)_macos ./bin/$(BINARY_NAME)
	zip ./bin/$(BINARY_NAME)_$(ARCH)_macos.zip ./bin/$(BINARY_NAME)
	rm -rf ./bin/$(BINARY_NAME)

.PHONY: all
all: clean linux windows macos

.PHONY: zip-all
zip-all: zip-linux zip-windows zip-macos

.PHONY: release
release: clean all zip-all

clean:
	go mod tidy
	rm -rf ./bin/*
