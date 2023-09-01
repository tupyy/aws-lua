vendor:
	go mod tidy
	go mod vendor

build:
	go build -o bin/aws-lua main.go

