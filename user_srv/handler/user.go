package handler

import (
	"context"

	"gorm.io/gorm"

	"github.com/liuyongbing/hello-go-srvs/user_srv/global"

	"github.com/liuyongbing/hello-go-srvs/user_srv/model"
	"github.com/liuyongbing/hello-go-srvs/user_srv/proto"
)

//type UserServer interface {
//	GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
//	GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
//	GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error)
//	CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
//	UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error)
//	CheckPassword(context.Context, *PasswordCheckInfo) (*CheckResponse, error)
//	mustEmbedUnimplementedUserServer()
//}
type UserServer struct {
}

// 返回数据结构绑定
func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		NickName: user.Nickname,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Mobile:   user.Mobile,
	}

	// 处理无默认值的情况
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}

	return userInfoRsp
}

// 分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	// 返回的数据体
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	// 分页
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	// 返回数据结构绑定
	for _, user := range users {
		userInfoRsp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}

	return rsp, nil
}
