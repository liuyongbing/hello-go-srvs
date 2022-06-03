package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/liuyongbing/hello-go-srvs/user_srv/proto"
)

var (
	conn       *grpc.ClientConn
	userClient proto.UserClient
)

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	userClient = proto.NewUserClient(conn)
}

//GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error)
//GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
//GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
//CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error)
//UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
//CheckPassword(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error)

/*
TestGetUserList
测试：获取用户列表
*/
func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 4,
	})
	if err != nil {
		panic(err)
	}

	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)

		checkRsp, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			PassWord:          "Aa123456",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

/*
TestCreateUser
测试：创建用户
*/
func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("client_test_user_%d", i),
			Mobile:   fmt.Sprintf("1881234567%d", i),
			PassWord: "Aa123456",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()

	// 测试：创建用户
	TestCreateUser()

	// 测试：获取用户列表
	TestGetUserList()

	conn.Close()
}
