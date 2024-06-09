package commands

import (
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/AlexTerra21/gophkeeper/cmd/client/types"
	"github.com/AlexTerra21/gophkeeper/pb"
)

func SavePassword(cond *types.Condition) {
	name := getStringData("Enter brief description login/password")
	login := getStringData("Enter login")
	password := getPasswordData("Enter password")

	ctx := metadata.AppendToOutgoingContext(cond.Ctx, "authorization", cond.AuthToken)

	in := &pb.SavePasswordRequest{
		Name:     name,
		Login:    login,
		Password: password,
	}

	_, err := cond.Client.SavePassword(ctx, in)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Save password successful")
}

func GetPassword(cond *types.Condition) {
	name := getStringData("Enter brief description login/password")

	ctx := metadata.AppendToOutgoingContext(cond.Ctx, "authorization", cond.AuthToken)

	in := &pb.GetSecretRequest{
		Name: name,
	}
	resp, err := cond.Client.GetPassword(ctx, in)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("%+v\n", resp)
}
