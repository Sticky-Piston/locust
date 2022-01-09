generate:
	protoc -I=. --go_out=./protocols ./protocols/locust.proto

run:
	go run main.go

build:
	go build