all: clean generate build

generate:
	protoc -I=. --go_out=./protocols ./protocols/locust.proto

build:
	go build

clean:
	rm -rf ./protocols/generated/*
