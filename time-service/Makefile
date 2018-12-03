build:
    #生成两个proto.go文件
	protoc --proto_path=$(GOPATH)/src:./proto/time --micro_out=./proto/time --go_out=./proto/time  proto/time/time.proto

	# dep 工具暂不可用，直接手动编译
	GOOS=linux GOARCH=amd64 go build
	docker build -t time-service .

run:
	docker run --name time-service -d -p 8082:8080 -e MICRO_SERVER_ADDRESS=:8081 -e MICRO_REGISTRY=mdns time-service

