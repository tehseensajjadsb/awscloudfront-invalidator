PLATFORMS := linux windows darwin
ARCH = amd64

$(PLATFORMS):
	GOOS=$@ GOARCH=$(ARCH) CGO_ENABLED=0 go build -v -o ./bin/invalidator_$(ARCH)_$@ ./...

release: $(PLATFORMS)

.PHONY: clean release $(PLATFORMS)

run:
	go run ./...

clean:
	go mod tidy
	rm -rf ./bin/*
