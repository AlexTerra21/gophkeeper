package commands

import (
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/AlexTerra21/gophkeeper/cmd/client/types"
	"github.com/AlexTerra21/gophkeeper/pb"
)

func SaveText(cond *types.Condition) {
	name := getStringData("Enter brief description note")
	text := getStringData("Enter note")

	ctx := metadata.AppendToOutgoingContext(cond.Ctx, "authorization", cond.AuthToken)

	in := &pb.SaveTextRequest{
		Name: name,
		Text: text,
	}

	_, err := cond.Client.SaveText(ctx, in)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Save note successful")
}

func GetText(cond *types.Condition) {
	name := getStringData("Enter brief description note")

	ctx := metadata.AppendToOutgoingContext(cond.Ctx, "authorization", cond.AuthToken)

	in := &pb.GetSecretRequest{
		Name: name,
	}
	resp, err := cond.Client.GetText(ctx, in)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("%+v\n", resp)
}
