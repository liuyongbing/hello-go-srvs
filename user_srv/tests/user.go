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

//UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)

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

/*
TestGetUserByMobile
测试: 通过手机获取用户
*/
func TestGetUserByMobile() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
			Mobile: fmt.Sprintf("1881234567%d", i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("User Mobil: ", rsp.Mobile)
	}
}

/*
TestGetUserById
测试: 通过User Id获取用户
*/
func TestGetUserById() {
	for i := 1; i <= 10; i++ {
		rsp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
			Id: int32(i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("User ID: ", rsp.Id)
	}
}

func main() {
	Init()

	// 测试：创建用户
	//TestCreateUser()

	// 测试：获取用户列表
	TestGetUserList()

	// 测试: 通过手机获取用户
	TestGetUserByMobile()

	//测试: 通过User Id获取用户
	TestGetUserById()

	conn.Close()
}
