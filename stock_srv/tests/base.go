package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/liuyongbing/hello-go-srvs/stock_srv/proto"
)

var (
	conn       *grpc.ClientConn
	grpcClient proto.StockClient
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

func main() {
	Init()

	// Testing: Hello Stock
	RunTestCase(TestHelloStock, "TestHelloStock")

	var goodsId, goodsNum int32
	goodsId = 421
	goodsNum = 100

	// Testing: SetInv
	TestSetInv(goodsId, goodsNum)

	// Testing: InvDetail
	TestInvDetail(goodsId)

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
