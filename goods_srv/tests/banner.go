package main

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

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
