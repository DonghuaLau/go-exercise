
all: 
	@echo "Usage: make <target>"


cidr-mask:
	go build -o $@.out $@.go

simple-client:
	go build -o $@.out $@.go

server-client:
	go build -o server.out server.go
	go build -o client.out client.go

http:
	go build -o $@.out $@.go
