package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	goodsProto "github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
	"github.com/liuyongbing/hello-go-srvs/stock_srv/proto"
)

var (
	conn       *grpc.ClientConn
	grpcClient proto.StockClient

	goodsGrpcClient goodsProto.GoodsClient
)

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	grpcClient = proto.NewStockClient(conn)
}

func InitGoodsGrpcClient() {
	var err error
	goodsGrpcConn, err := grpc.Dial("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	// defer goodsGrpcConn.Close()

	goodsGrpcClient = goodsProto.NewGoodsClient(goodsGrpcConn)
}

func main() {
	Init()

	InitGoodsGrpcClient()

	// Testing: Hello Stock
	RunTestCase(TestHelloStock, "TestHelloStock")

	// Testing: Banner list
	RunTestCase(TestBannerList, "BannerList")

	goodsId := 421
	// goodsNum := 20

	// Testing: SetInv
	// TestSetInv(GoodsId, GoodsNum)

	TestInvDetail(int32(goodsId))

	conn.Close()
}

func RunTestCase(TestCase func(), FuncName string) {
	fmt.Printf("Testing of func %s \n", FuncName)
	TestCase()
	fmt.Println()
}

func TestHelloStock() {
	rsp, err := grpcClient.HelloStock(context.Background(), &proto.HelloStockRequest{
		Name: "Test Say Hello.",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Message)
}

/*
TestBannerList
*/
func TestBannerList() {
	rsp, err := goodsGrpcClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Total:", rsp.Total)
	for _, v := range rsp.Data {
		fmt.Println("ID:", v.Id)
	}
}
