package handler

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/global"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/model"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
)

// GetAllCategorysList
func (s *GoodsServer) GetAllCategorysList(ctx context.Context, request *emptypb.Empty) (*proto.CategoryListResponse, error) {
	response := proto.CategoryListResponse{}

	items := []model.Category{}
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&items)
	b, err := json.Marshal(&items)
	if err != nil {
		return nil, err
	}

	response.JsonData = string(b)

	return &response, nil
}

// GetSubCategory
func (s *GoodsServer) GetSubCategory(ctx context.Context, request *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	response := proto.SubCategoryListResponse{
		// Total:        0,
		// Info:         &proto.CategoryInfoResponse{},
		// SubCategorys: []*proto.CategoryInfoResponse{},
	}

	item := model.Category{}
	// global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&items)
	global.DB.Preload("SubCategory.SubCategory").First(&item, request.Id)
	// b, err := json.Marshal(&item)
	// if err != nil {
	// 	return nil, err
	// }

	response.Info = &proto.CategoryInfoResponse{
		Id:             item.ID,
		Name:           item.Name,
		ParentCategory: item.ParentCategoryID,
		Level:          item.Level,
		IsTab:          item.IsTab,
	}

	// children := []*proto.CategoryInfoResponse
	var children []*proto.CategoryInfoResponse
	for _, child := range item.SubCategory {
		children = append(children, &proto.CategoryInfoResponse{
			Id:             child.ID,
			Name:           child.Name,
			ParentCategory: child.ParentCategoryID,
			Level:          child.Level,
			IsTab:          child.IsTab,
		})
	}
	response.SubCategorys = children

	return &response, nil
}

// CreateCategory
func (s *GoodsServer) CreateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	response := proto.CategoryInfoResponse{}

	// 检查重复
	result := global.DB.Where(&model.Category{Name: request.Name}).First(&model.Category{})
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "已存在")
	}

	item := model.Category{
		Name:             request.Name,
		ParentCategoryID: request.ParentCategory,
		Level:            request.Level,
		IsTab:            request.IsTab,
	}
	global.DB.Save(&item)

	response.Id = item.ID

	return &response, nil
}

// UpdateCategory
func (s *GoodsServer) UpdateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	item := model.Category{}
	// 检查重复
	result := global.DB.First(&item, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "已存在")
	}

	// 判断空值
	if request.Name != "" {
		item.Name = request.Name
	}
	if request.ParentCategory != 0 {
		item.ParentCategoryID = request.ParentCategory
	}
	if request.Level != 0 {
		item.Level = request.Level
	}
	if request.IsTab {
		item.IsTab = request.IsTab
	}
	global.DB.Save(&item)

	return &emptypb.Empty{}, nil
}

// DeleteCategory
func (s *GoodsServer) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "不存在")
	}
	return &emptypb.Empty{}, nil
}
