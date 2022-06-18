package main

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
)

var (
	conn       *grpc.ClientConn
	grpcClient proto.GoodsClient
)

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	grpcClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()

	// Testing: Banner list
	RunTestCase(TestBannerList, "BannerList")

	// Testing: Brands list
	RunTestCase(TestBrandList, "BrandList")

	conn.Close()
}

func RunTestCase(TestCase func(), FuncName string) {
	fmt.Printf("Testing of func %s \n", FuncName)
	TestCase()
	fmt.Println()
}
