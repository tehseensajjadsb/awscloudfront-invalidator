
build:
	CGO_ENABLED=0 go build -o bin/awscfn ./src/

run:
	go run ./src/

clean:
	go mod tidy
	rm -rf ./bin/*
