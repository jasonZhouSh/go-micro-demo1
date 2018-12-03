build:
    #生成两个proto.go文件
	protoc --proto_path=$(GOPATH)/src:./proto/user --micro_out=./proto/user --go_out=./proto/user  proto/user/user.proto

	# dep 工具暂不可用，直接手动编译
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service .

run:
	docker run --name user-service  -p 8083:8083 -e MICRO_SERVER_ADDRESS=:8084 -e MICRO_REGISTRY=mdns user-service