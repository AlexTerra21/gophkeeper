package commands

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/AlexTerra21/gophkeeper/cmd/client/types"
	"github.com/AlexTerra21/gophkeeper/pb"
)

// Команда для аутентификации пользователя
func Login(cond *types.Condition) {

	login := getStringData("Enter login")
	password := getPasswordData("Enter password")

	var header metadata.MD
	in := &pb.LoginRequest{
		Username: login,
		Password: password,
	}
	_, err := cond.Client.Login(cond.Ctx, in, grpc.Header(&header))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cond.ISAuth = true
	cond.AuthToken = header.Get("authorization")[0]
	fmt.Println("Authenticated successful")
}

// Команда для регистрации пользователя
func Register(cond *types.Condition) {
	login := getStringData("Enter login")
	password := getPasswordData("Enter password")

	var header metadata.MD
	in := &pb.RegisterRequest{
		Username: login,
		Password: password,
	}
	_, err := cond.Client.Register(cond.Ctx, in, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return
	}
	cond.ISAuth = true
	cond.AuthToken = header.Get("authorization")[0]
	fmt.Println("Authenticated successful")
}
