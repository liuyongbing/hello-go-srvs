package handler

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/global"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/model"
	"github.com/liuyongbing/hello-go-srvs/goods_srv/proto"
)

type GoodsServer struct {
	// proto.UnimplementedUserServer
	proto.UnimplementedGoodsServer
}

func (s *GoodsServer) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Time: " + time.Now().String() + "; Request context: " + request.Name,
	}, nil
}

/*
ModelToResponse
返回数据结构绑定
*/
func ModelToResponse(goods model.Goods) proto.GoodsInfoResponse {
	modelRes := proto.GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		ShipFree:        goods.ShipFree,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		DescImages:      goods.DescImages,
		Images:          goods.Images,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   goods.Category.ID,
			Name: goods.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   goods.Brands.ID,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
	}

	return modelRes
}

// 	// 现在用户提交订单有多个商品，你得批量查询商品的信息吧
// 	BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error)
// 	GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)
// 	CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error)
// 	UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
// 	DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)

/*
GetUserList
*/
func (s *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	// 关键词搜索、查询新品、查询热门商品、通过价格区间筛选， 通过商品分类筛选
	goodsListResponse := &proto.GoodsListResponse{}

	var goods []model.Goods
	localDB := global.DB.Model(model.Goods{})
	if req.KeyWords != "" {
		//搜索
		localDB = localDB.Where("name LIKE ?", "%"+req.KeyWords+"%")
	}
	if req.IsHot {
		localDB = localDB.Where(model.Goods{IsHot: true})
	}
	if req.IsNew {
		localDB = localDB.Where(model.Goods{IsNew: true})
	}

	if req.PriceMin > 0 {
		localDB = localDB.Where("shop_price>=?", req.PriceMin)
	}
	if req.PriceMax > 0 {
		localDB = localDB.Where("shop_price<=?", req.PriceMax)
	}

	if req.Brand > 0 {
		localDB = localDB.Where("brands_id=?", req.Brand)
	}

	//通过category去查询商品
	var subQuery string
	if req.TopCategory > 0 {
		var category model.Category
		if result := global.DB.First(&category, req.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", req.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category WHERE parent_category_id=%d", req.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category WHERE id=%d", req.TopCategory)
		}

		localDB = localDB.Where(fmt.Sprintf("category_id in (%s)", subQuery))
	}

	var count int64
	localDB.Count(&count)
	goodsListResponse.Total = int32(count)

	result := localDB.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}

	return goodsListResponse, nil
}

// /*
// GetUserByMobile
// 通过 Mobile 查询用户
// */
// func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
// 	var user model.User
// 	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
// 	if result.RowsAffected == 0 {
// 		return nil, status.Errorf(codes.NotFound, "用户不存在")
// 	}
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	userInfoRsp := ModelToResponse(user)

// 	return &userInfoRsp, nil
// }

// /*
// GetUserById
// 通过 Id 查询用户
// */
// func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
// 	var user model.User
// 	result := global.DB.First(&user, req.Id)
// 	if result.RowsAffected == 0 {
// 		return nil, status.Errorf(codes.NotFound, "用户不存在")
// 	}
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	userInfoRsp := ModelToResponse(user)

// 	return &userInfoRsp, nil
// }

// /*
// CreateUser
// 创建用户
// */
// func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
// 	var user model.User
// 	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
// 	if result.RowsAffected == 1 {
// 		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
// 	}

// 	user.Mobile = req.Mobile
// 	user.Nickname = req.NickName

// 	//密码加密
// 	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
// 	salt, encodedPwd := password.Encode(req.PassWord, options)
// 	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

// 	result = global.DB.Create(&user)
// 	if result.Error != nil {
// 		return nil, status.Errorf(codes.Internal, result.Error.Error())
// 	}

// 	userInfoRsp := ModelToResponse(user)

// 	return &userInfoRsp, nil
// }

// /*
// UpdateUser
// 修改用户信息
// */
// func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
// 	var user model.User
// 	result := global.DB.First(user, req.Id)
// 	if result.RowsAffected == 0 {
// 		return nil, status.Errorf(codes.NotFound, "用户不存在")
// 	}

// 	birthDay := time.Unix(int64(req.BirthDay), 0)
// 	user.Nickname = req.NickName
// 	user.Birthday = &birthDay
// 	user.Gender = req.Gender

// 	result = global.DB.Save(user)
// 	if result.Error != nil {
// 		return nil, status.Errorf(codes.Internal, result.Error.Error())
// 	}

// 	return &empty.Empty{}, nil
// }

// /*
// CheckPassword
// 检查用户密码
// */
// func (s *UserServer) CheckPassword(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
// 	//密码校验
// 	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
// 	passwordInfo := strings.Split(req.EncryptedPassword, "$")
// 	check := password.Verify(req.PassWord, passwordInfo[2], passwordInfo[3], options)

// 	return &proto.CheckResponse{Success: check}, nil
// }
