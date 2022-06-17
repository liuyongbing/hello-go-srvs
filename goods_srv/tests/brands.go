package main

import (
	"context"
	"fmt"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
)

/*
TestBrandList
*/
func TestBrandList() {
	rsp, err := grpcClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Total:", rsp.Total)
	for _, v := range rsp.Data {
		fmt.Println("ID:", v.Id)
	}
}
