package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	uuid "github.com/satori/go.uuid"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/global"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/handler"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/initialize"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/utils"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/utils/register/consul"
)

func main() {
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	// flag 解析命令行参数
	IP := flag.String("ip", "0.0.0.0", "IP地址")
	Port := flag.Int("port", 50052, "端口号")
	flag.Parse()

	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}

	server := grpc.NewServer()
	proto.RegisterGoodsServer(server, &handler.GoodsServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("Failed to listen: " + err.Error())
	}

	// 服务健康检查注册
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	addr := global.ServerConfig.Host
	port := *Port
	name := global.ServerConfig.Name
	// id := global.ServerConfig.Name
	// 负载均衡：通过终端开启多个服务
	id := uuid.NewV4().String()
	tags := global.ServerConfig.Tags
	// utils.Register(addr, port, name, tags, id)
	regClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	err = regClient.Register(addr, port, name, tags, id)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}

	fmt.Printf("服务启动中:[Name:%s][IP:%s][Port:%d]\n", global.ServerConfig.Name, addr, *Port)

	go func() {
		err = server.Serve(lis)
		if err != nil {
			zap.S().Panic("Failed to start grpc: " + err.Error())
		}
	}()

	fmt.Println("服务启动成功")

	// 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 服务注销
	if err := regClient.DeRegister(id); err != nil {
		zap.S().Info("注销失败", err.Error())
	} else {
		zap.S().Info("注销成功")
	}
}
