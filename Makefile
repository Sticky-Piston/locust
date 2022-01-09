all: clean generate build

generate:
	protoc -I=. --go_out=./protocols ./protocols/locust.proto

run:
	go run main.go

build:
	go build

clean:
	rm -rf ./protocols/generated/*
