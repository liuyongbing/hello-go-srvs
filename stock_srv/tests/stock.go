package main

import (
	"context"
	"fmt"

	"github.com/liuyongbing/hello-go-srvs/stock_srv/proto"
)

func TestSetInv(goodsId, Num int32) {
	_, err := grpcClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num:     Num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")
}

func TestInvDetail(goodsId int32) {
	rsp, err := grpcClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Num)
}

// func TestSell(wg *sync.WaitGroup) {
// 	/*
// 		1. 第一件扣减成功： 第二件： 1. 没有库存信息 2. 库存不足
// 		2. 两件都扣减成功
// 	*/
// 	defer wg.Done()
// 	_, err := grpcClient.Sell(context.Background(), &proto.SellInfo{
// 		GoodsInfo: []*proto.GoodsInvInfo{
// 			{GoodsId: 421, Num: 1},
// 			//{GoodsId: 422, Num: 30},
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("库存扣减成功")
// }

// func TestReback() {
// 	_, err := grpcClient.Reback(context.Background(), &proto.SellInfo{
// 		GoodsInfo: []*proto.GoodsInvInfo{
// 			{GoodsId: 421, Num: 10},
// 			{GoodsId: 422, Num: 30},
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("归还成功")
// }
