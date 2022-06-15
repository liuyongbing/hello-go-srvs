package handler

import (
	"context"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/global"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/model"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
)

// // Brand
// BrandList(context.Context, *BrandFilterRequest) (*BrandListResponse, error)
// CreateBrand(context.Context, *BrandRequest) (*BrandInfoResponse, error)
// UpdateBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
// DeleteBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)

/*
BrandList
*/
func (s *GoodsServer) BrandList(ctx context.Context, request *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	// 返回列表
	// var response *proto.BrandListResponse
	response := proto.BrandListResponse{}

	// 查询结果
	var list []model.Brands
	result := global.DB.Scopes(Paginate(int(request.Pages), int(request.PagePerNums))).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	// 记录总数
	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	response.Total = int32(total)

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
