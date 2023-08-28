# CloudFront Distribution Invalidation Tool

## Installation

### Download Latest Release

Download the latest release and move it to your PATH

### Build from source

Clone Repository

```
git clone https:/github.com/tehseensajjadsb/awscloudfront-invalidator.git
cd awscloudfront-invaliator
```

Choose `GOOS` & `GOARCH`. Build the binary

```
GOOS=linux
GOARCH=amd64 
CGO_ENABLED=0 go build -v ./... -o ./bin/invalidator
```

```
chmod +x ./bin/invalidator
mv ./bin/invalidator /usr/local/bin/invalidator
```

## Usage

```
$ invalidator --help
```


# Features / TODO

- Main Tool Features:
    - [X] Add `aws sts get-caller-identity` feature
    - [X] Add Cache Invalidation Feature
    - [X] Detect Distribution by its ID
- QOL Features
    - [ ] Detect Distribution by its first domain alias
    - [X] Detect Distribution by its S3 Origin
    - [ ] Optional wait for invalidation to complete
