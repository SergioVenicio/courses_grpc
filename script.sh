protoc --go_out=gen --go_opt=paths=source_relative \
    --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
    proto/courses.proto

dlv debug --headless --listen=localhost:2345 --log --api-version=2

kill $(lsof -t -i:2345)
