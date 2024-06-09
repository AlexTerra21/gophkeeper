package commands

import (
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/AlexTerra21/gophkeeper/cmd/client/types"
	"github.com/AlexTerra21/gophkeeper/pb"
)

func SaveCard(cond *types.Condition) {
	name := getStringData("Enter brief description card")
	number := getStringData("Enter card number")
	holderName := getStringData("Enter card holder name")
	date := getStringData("Enter card date")
	ccv := getStringData("Enter card ccv")

	ctx := metadata.AppendToOutgoingContext(cond.Ctx, "authorization", cond.AuthToken)

	in := &pb.SaveCardRequest{
		CardName:   name,
		Number:     number,
		HolderName: holderName,
		Date:       date,
		Ccv:        ccv,
	}

	_, err := cond.Client.SaveCard(ctx, in)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Save card successful")
}

func GetCard(cond *types.Condition) {
	name := getStringData("Enter brief description card")

	ctx := metadata.AppendToOutgoingContext(cond.Ctx, "authorization", cond.AuthToken)

	in := &pb.GetSecretRequest{
		Name: name,
	}
	resp, err := cond.Client.GetCard(ctx, in)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("%+v\n", resp)
}
