package commands

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/imroc/biu"
	"google.golang.org/grpc/metadata"

	"github.com/AlexTerra21/gophkeeper/cmd/client/types"
	"github.com/AlexTerra21/gophkeeper/pb"
)

// Команда для сохранения бинарного файла
func SaveFile(cond *types.Condition) {
	name := getStringData("Enter brief description file")
	path := getStringData("Enter file path")

	ctx := metadata.AppendToOutgoingContext(cond.Ctx, "authorization", cond.AuthToken)

	buf, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := biu.BytesToBinaryString(buf)

	fmt.Printf("Length file = %v\n", len(text))
	if len(text) > 4100000 {
		fmt.Println("File too big")
		return
	}
	in := &pb.SaveTextRequest{
		Name: name,
		Text: text,
	}

	_, err = cond.Client.SaveText(ctx, in)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Save file successful")
}

// Команда для загрузки бинарного файла
func GetFile(cond *types.Condition) {
	name := getStringData("Enter brief description file")
	path := getStringData("Enter path to save file")

	ctx := metadata.AppendToOutgoingContext(cond.Ctx, "authorization", cond.AuthToken)

	in := &pb.GetSecretRequest{
		Name: name,
	}
	resp, err := cond.Client.GetText(ctx, in)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := biu.BinaryStringToBytes(resp.Text)
	err = os.WriteFile(path, buf, fs.FileMode(os.O_CREATE))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("File %s saved to %s\n", name, path)
}
