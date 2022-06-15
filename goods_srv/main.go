package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	uuid "github.com/satori/go.uuid"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/global"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/handler"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/initialize"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/utils"
)

func main() {
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	// flag 解析命令行参数
	IP := flag.String("ip", "0.0.0.0", "IP地址")
	Port := flag.Int("port", 0, "端口号")
	flag.Parse()

	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("Failed to listen: " + err.Error())
	}

	// 服务健康检查注册
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	// addr := "192.168.31.141"
	addr := "10.8.19.134"
	port := *Port
	name := global.ServerConfig.Name
	// id := global.ServerConfig.Name
	// 负载均衡：通过终端开启多个服务
	id := uuid.NewV4().String()
	tags := []string{
		"user-srv",
		"gosrv-register",
		"consul",
	}
	utils.Register(addr, port, name, tags, id)

	fmt.Printf("服务启动中:[Name:%s][IP:%s][Port:%d]", global.ServerConfig.Name, *IP, *Port)

	// go func() {
	err = server.Serve(lis)
	if err != nil {
		panic("Failed to start grpc: " + err.Error())
	}
	// }()

	fmt.Println("服务启动成功")

}