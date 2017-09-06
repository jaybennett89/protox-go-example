build:
	protoc --go_out=example/ example.proto

test: 
	go test -v ./example
