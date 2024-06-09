package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/AlexTerra21/gophkeeper/pb"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/AlexTerra21/gophkeeper/cmd/client/commands"
	"github.com/AlexTerra21/gophkeeper/cmd/client/types"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

// go build -o cmd/client/gophkeeper_client.exe -ldflags "-X main.buildVersion=v1.20.0 -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')' -X 'main.buildCommit=$(git log -1 | grep commit)'" cmd/client/*.go
// ./cmd/client/gophkeeper_client.exe
func main() {
	ctx := context.Background()

	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	address := "0.0.0.0:3200"
	cond := &types.Condition{
		Ctx:    ctx,
		ISAuth: false,
	}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		fmt.Println("cannot load TLS credentials")
	}

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		fmt.Println("No gRPC connection: " + err.Error())
	}
	defer conn.Close()

	c := pb.NewGophkeeperClient(conn)
	cond.Client = c

	mainMenu(cond)

}

func mainMenu(cond *types.Condition) {
	prompt := promptui.Select{
		Label: "GophKeepeer",
		Items: []string{"Login", "Save Secret", "Get Secret", "Exit"},
	}
	for {
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("prompt failed %v\n", err)
			return
		}
		switch idx {
		case 0: // Login
			clientAuth(cond)
		case 1: // Save
			saveSecret(cond)
		case 2: // Get
			getSecret(cond)
		case 3: // Exit
			return
		}
	}
}

func clientAuth(cond *types.Condition) {
	prompt := promptui.Select{
		Label: "Welcome",
		Items: []string{"..", "Sign in", "Register"},
	}
	for {
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("prompt failed %v\n", err)
			return
		}
		switch idx {
		case 0: // Exit
			return
		case 1: // Sign in
			commands.Login(cond)
		case 2: // Register
			commands.Register(cond)

		}
	}
}

func saveSecret(cond *types.Condition) {
	if !cond.ISAuth {
		fmt.Println("Unauthenticated")
		return
	}
	prompt := promptui.Select{
		Label: "Save secret",
		Items: []string{"..", "Save password", "Save card", "Save text", "Save file"},
	}
	for {
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("prompt failed %v\n", err)
			return
		}
		switch idx {
		case 0: // Exit
			return
		case 1: // Save password
			commands.SavePassword(cond)
		case 2: // Save card
			commands.SaveCard(cond)
		case 3: // Save text
			commands.SaveText(cond)
		case 4: // Save file
			commands.SaveFile(cond)
		}
	}

}

func getSecret(cond *types.Condition) {
	if !cond.ISAuth {
		fmt.Println("Unauthenticated")
		return
	}
	prompt := promptui.Select{
		Label: "Get secret",
		Items: []string{"..", "Get password", "Get card", "Get text", "Get file"},
	}
	for {
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("prompt failed %v\n", err)
			return
		}
		switch idx {
		case 0: // Exit
			return
		case 1: // Get password
			commands.GetPassword(cond)
		case 2: // Get card
			commands.GetCard(cond)
		case 3: // Get text
			commands.GetText(cond)
		case 4: // Get file
			commands.GetFile(cond)
		}
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
