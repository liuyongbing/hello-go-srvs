package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/hashicorp/consul/api"

	"github.com/liuyongbing/hello-go-srvs/user_srv/global"
	"github.com/liuyongbing/hello-go-srvs/user_srv/handler"
	"github.com/liuyongbing/hello-go-srvs/user_srv/initialize"
	"github.com/liuyongbing/hello-go-srvs/user_srv/proto"
)

/*
Register
服务注册
*/
func Register(addr string, port int, name string, tags []string, id string) {
	consulInfo := global.ServerConfig.ConsulInfo

	cfg := api.DefaultConfig()
	// cfg.Address = "127.0.0.1:8500"
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 服务健康检查对象
	check := api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", addr, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 服务注册对象
	registration := api.AgentServiceRegistration{
		Name:    name,
		ID:      id,
		Tags:    tags,
		Port:    port,
		Address: addr,
		Check:   &check,
	}

	// 注册服务
	err = client.Agent().ServiceRegister(&registration)
	if err != nil {
		panic(err)
	}
}

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

	// 服务健康检查注册
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	addr := "192.168.31.141"
	port := *Port
	name := global.ServerConfig.Name
	id := global.ServerConfig.Name
	tags := []string{
		"user-srv",
		"gosrv-register",
		"consul",
	}
	Register(addr, port, name, tags, id)

	err = server.Serve(lis)
	if err != nil {
		panic("Failed to start grpc: " + err.Error())
	}
}
