protofiles:
	protoc --go_out=gen --go_opt=paths=source_relative \
		--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
		proto/courses.proto

dependencies: protofiles
	go mod tidy

serve-dev-debug: dependencies
	dlv debug --headless --listen=localhost:2345 --log --api-version=2

serve-dev: dependencies
	go run main.go serve

build: dependencies
	go build -o main.bin main.go 

serve: build
	./main.bin serve