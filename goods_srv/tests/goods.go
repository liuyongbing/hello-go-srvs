package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
)

var (
	conn       *grpc.ClientConn
	grpcClient proto.GoodsClient
)

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	grpcClient = proto.NewGoodsClient(conn)
}

/*
TestBannerList
*/
func TestBannerList() {
	rsp, err := grpcClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Total:", rsp.Total)
	for _, v := range rsp.Data {
		fmt.Println("ID:", v.Id)
	}
}

func main() {
	Init()

	// 测试：获取用户列表
	TestBannerList()

	conn.Close()
}
