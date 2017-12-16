build:
	go build -o cli *.go

run: build
	./cli

help:
	go run *.go help
