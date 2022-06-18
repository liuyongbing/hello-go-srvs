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
CategoryBrandList
*/
func (s *GoodsServer) CategoryBrandList(ctx context.Context, request *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	// 返回列表
	response := proto.CategoryBrandListResponse{}

	// 查询结果
	var list []model.GoodsCategoryBrand
	result := global.DB.Preload("Category").Preload("Brands").Scopes(Paginate(int(request.Pages), int(request.PagePerNums))).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	// 记录总数
	var count int64
	global.DB.Model(&model.GoodsCategoryBrand{}).Count(&count)
	response.Total = int32(count)

	// 返回实例
	var res []*proto.CategoryBrandResponse
	for _, v := range list {
		r := proto.CategoryBrandResponse{
			Id: v.ID,
			Brand: &proto.BrandInfoResponse{
				Id:   v.Brands.ID,
				Name: v.Brands.Name,
				Logo: v.Brands.Logo,
			},
			Category: &proto.CategoryInfoResponse{
				Id:             v.Category.ID,
				Name:           v.Category.Name,
				ParentCategory: v.Category.ParentCategoryID,
				Level:          v.Category.Level,
				IsTab:          v.Category.IsTab,
			},
		}
		res = append(res, &r)
	}
	response.Data = res

	return &response, nil
}

/*
CategoryBrandList
*/
func (s *GoodsServer) GetCategoryBrandList(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	// 返回列表
	response := proto.BrandListResponse{}

	// 是否存在
	result := global.DB.First(&model.Category{}, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "不存在")
	}

	// 查询结果
	var list []model.GoodsCategoryBrand
	result = global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{CategoryID: request.Id}).Find(&list)
	if result.RowsAffected > 0 {
		response.Total = int32(result.RowsAffected)
	}

	// 返回实例
	var res []*proto.BrandInfoResponse
	for _, v := range list {
		r := proto.BrandInfoResponse{
			Id:   v.Brands.ID,
			Name: v.Brands.Name,
			Logo: v.Brands.Logo,
		}
		res = append(res, &r)
	}
	response.Data = res

	return &response, nil
}

/*
CreateCategoryBrand
*/
func (s *GoodsServer) CreateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	// 检查存在
	conditions := &model.GoodsCategoryBrand{
		CategoryID: request.CategoryId,
		BrandsID:   request.BrandId,
	}
	result := global.DB.Where(conditions).First(&model.GoodsCategoryBrand{})
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "已存在")
	}

	// 是否存在：分类
	result = global.DB.First(&model.Category{}, request.CategoryId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "分类不存在")
	}

	// 是否存在：品牌
	result = global.DB.First(&model.Brands{}, request.BrandId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	// 新增记录
	item := &model.GoodsCategoryBrand{
		CategoryID: request.CategoryId,
		BrandsID:   request.BrandId,
	}
	global.DB.Save(item)

	return &proto.CategoryBrandResponse{
		Id: item.ID,
	}, nil
}

/*
UpdateCategoryBrand
*/
func (s *GoodsServer) UpdateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	item := model.GoodsCategoryBrand{}
	// 是否存在
	result := global.DB.First(&item, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "主记录不存在")
	}

	// 是否存在：分类
	result = global.DB.First(&model.Category{}, request.CategoryId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "分类不存在")
	}

	// 是否存在：品牌
	result = global.DB.First(&model.Brands{}, request.BrandId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	item.CategoryID = request.CategoryId
	item.BrandsID = request.BrandId

	// 修改记录
	global.DB.Save(&item)

	return &emptypb.Empty{}, nil
}

/*
DeleteCategoryBrand
*/
func (s *GoodsServer) DeleteCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	// 删除记录
	if result := global.DB.Delete(&model.GoodsCategoryBrand{}, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "不存在")
	}

	return &emptypb.Empty{}, nil
}
