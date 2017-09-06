build:
	protoc --go_out=experiment/ experiment.proto

test: 
	go test -v ./experiment
