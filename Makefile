build:
	go build -o bin/awscfn ./src/

run:
	go run ./src/

clean:
	go mod tidy
	rm -rf ./bin/*
