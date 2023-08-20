build:
	go build -o bin/awscfn main.go

run:
	go run main.go

clean:
	go mod tidy
	rm -rf ./bin/*
