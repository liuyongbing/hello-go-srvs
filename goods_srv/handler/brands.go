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
BrandList
*/
func (s *GoodsServer) BrandList(ctx context.Context, request *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	// 返回列表
	response := proto.BrandListResponse{}

	// 查询结果
	var list []model.Brands
	result := global.DB.Scopes(Paginate(int(request.Pages), int(request.PagePerNums))).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	// 记录总数
	var count int64
	global.DB.Model(&model.Brands{}).Count(&count)
	response.Total = int32(count)

	// 返回实例
	var res []*proto.BrandInfoResponse
	for _, v := range list {
		r := proto.BrandInfoResponse{
			Id:   v.ID,
			Name: v.Name,
			Logo: v.Logo,
		}
		res = append(res, &r)
	}
	response.Data = res

	return &response, nil
}

/*
CreateBrand
*/
func (s *GoodsServer) CreateBrand(ctx context.Context, request *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	// 检查存在
	result := global.DB.Where("name=?", request.Name).First(&model.Brands{})
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "已存在")
	}

	// 新增记录
	item := &model.Brands{
		Name: request.Name,
		Logo: request.Logo,
	}
	global.DB.Save(item)

	return &proto.BrandInfoResponse{
		Id: item.ID,
	}, nil
}

/*
UpdateBrand
*/
func (s *GoodsServer) UpdateBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	item := model.Brands{}
	// 是否存在
	result := global.DB.First(&item, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "不存在")
	}

	// 检测空值
	if request.Name != "" {
		item.Name = request.Name
	}
	if request.Logo != "" {
		item.Logo = request.Logo
	}

	// 修改记录
	global.DB.Save(&item)

	return &emptypb.Empty{}, nil
}

/*
DeleteBrand
*/
func (s *GoodsServer) DeleteBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	// 删除记录
	if result := global.DB.Delete(&model.Brands{}, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "不存在")
	}

	return &emptypb.Empty{}, nil
}
