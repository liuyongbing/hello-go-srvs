package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/liuyongbing/hello-go-srvs/user_srv/handler"
	"github.com/liuyongbing/hello-go-srvs/user_srv/initialize"
	"github.com/liuyongbing/hello-go-srvs/user_srv/proto"
)

func main() {
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	IP := flag.String("ip", "0.0.0.0", "IP地址")
	Port := flag.Int("port", 50051, "端口号")
	flag.Parse()

	fmt.Println("IP: ", *IP)
	fmt.Println("Port: ", *Port)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("Failed to listen: " + err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic("Failed to start grpc: " + err.Error())
	}
}
