build:
	go build -o bin/awscfn .

run:
	go run .

clean:
	go mod tidy
	rm -rf ./bin/*
