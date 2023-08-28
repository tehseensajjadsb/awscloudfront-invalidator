# CloudFront Distribution Invalidation Tool

# Installation

- Build from source

```
# Choose Platform
# linux | windows | darwin
PLATFORM=linux
BINARY=invalidator_amd64_$PLATFORM
git clone https:/github.com/tehseensajjadsb/awscloudfront-invalidator.git
cd awscloudfront-invaliator
make $PLATFORM
chmod +x ./bin/$BINARY
mv ./bin/$BINARY ~/.local/bin/invalidator
```

# Features / TODO

- Main Tool Features:
    - [X] Add `aws sts get-caller-identity` feature
    - [X] Add Cache Invalidation Feature
    - [X] Detect Distribution by its ID
- QOL Features
    - [X] Detect Distribution by its first domain alias
    - [X] Detect Distribution by its S3 Origin
    - [ ] Optional wait for invalidation to complete
