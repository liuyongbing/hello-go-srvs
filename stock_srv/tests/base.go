package main

import (
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
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	grpcClient = proto.NewStockClient(conn)
}

// func main() {
// 	Init()

// 	// Testing: Say Hello
// 	RunTestCase(TestSayHello, "TestSayHello")

// 	conn.Close()
// }

func RunTestCase(TestCase func(), FuncName string) {
	fmt.Printf("Testing of func %s \n", FuncName)
	TestCase()
	fmt.Println()
}
