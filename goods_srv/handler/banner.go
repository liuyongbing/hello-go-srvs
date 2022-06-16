package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/global"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/model"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
)

/*
BannerList
*/
func (s *GoodsServer) BannerList(ctx context.Context, request *emptypb.Empty) (*proto.BannerListResponse, error) {
	response := proto.BannerListResponse{}

	var list []model.Banner
	result := global.DB.Find(&list)
	response.Total = int32(result.RowsAffected)

	var res []*proto.BannerResponse
	for _, v := range list {
		item := proto.BannerResponse{
			Id:    v.ID,
			Image: v.Image,
			Index: v.Index,
			Url:   v.Url,
		}
		res = append(res, &item)
	}
	response.Data = res

	return &response, nil
}

/*
CreateBanner
*/
func (s *GoodsServer) CreateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	// 新增
	item := model.Banner{
		Image: request.Image,
		Index: request.Index,
		Url:   request.Url,
	}
	global.DB.Save(&item)

	return &proto.BannerResponse{Id: item.ID}, nil
}

/*
UpdateBanner
*/
func (s *GoodsServer) UpdateBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	item := model.Banner{}

	if result := global.DB.First(&item, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "不存在")
	}

	// 检测空值
	if request.Image != "" {
		item.Image = request.Image
	}
	if request.Index != 0 {
		item.Index = request.Index
	}
	if request.Url != "" {
		item.Url = request.Url
	}

	// 修改
	global.DB.Save(&item)

	return &emptypb.Empty{}, nil
}

/*
DeleteBanner
*/
func (s *GoodsServer) DeleteBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Banner{}, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "不存在")
	}

	return &emptypb.Empty{}, nil
}
