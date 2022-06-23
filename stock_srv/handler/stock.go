package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/liuyongbing/hello-go-srvs/stock_srv/proto"
)

type StockServer struct {
	proto.UnimplementedStockServer
}

// SayHello(context.Context, *HelloRequest) (*HelloReply, error)

// Reback(context.Context, *SellInfo) (*emptypb.Empty, error)
/*
SetInv
*/
func (s *StockServer) SetInv(ctx context.Context, request *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	// TODO: Logic

	return &emptypb.Empty{}, nil
}

/*
InvDetail
*/
func (s *StockServer) InvDetail(ctx context.Context, request *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	// TODO: Logic

	return &proto.GoodsInvInfo{}, nil
}

/*
Sell
*/
func (s *StockServer) Sell(ctx context.Context, request *proto.SellInfo) (*emptypb.Empty, error) {
	// TODO: Logic

	return &emptypb.Empty{}, nil
}

/*
Reback
*/
func (s *StockServer) Reback(ctx context.Context, request *proto.SellInfo) (*emptypb.Empty, error) {
	// TODO: Logic

	return &emptypb.Empty{}, nil
}
