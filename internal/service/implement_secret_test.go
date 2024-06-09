package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/AlexTerra21/gophkeeper/pb"
	"google.golang.org/grpc/metadata"
)

func TestService_SaveGetPassword(t *testing.T) {
	service, err := PrepareTestEnv()
	if err != nil {
		t.Log(err.Error())
		return
	}
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("userID", "123"))
	type save_args struct {
		ctx context.Context
		req *pb.SavePasswordRequest
	}
	save_tests := []struct {
		name    string
		s       *Service
		args    save_args
		want    *pb.Empty
		wantErr bool
	}{
		{
			name: "save_ok",
			s:    service,
			args: save_args{
				ctx: ctx,
				req: &pb.SavePasswordRequest{
					Name:     "Name",
					Login:    "Login",
					Password: "Password",
				},
			},
			want:    &pb.Empty{},
			wantErr: false,
		},
		{
			name: "conflict",
			s:    service,
			args: save_args{
				ctx: ctx,
				req: &pb.SavePasswordRequest{
					Name:     "Name",
					Login:    "Login",
					Password: "Password",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range save_tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SavePassword(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.SavePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.SavePassword() = %v, want %v", got, tt.want)
			}

		})
	}
	type get_args struct {
		ctx context.Context
		req *pb.GetSecretRequest
	}
	get_tests := []struct {
		name    string
		s       *Service
		args    get_args
		want    *pb.PasswordResponse
		wantErr bool
	}{
		{
			name: "get_ok",
			s:    service,
			args: get_args{
				ctx: ctx,
				req: &pb.GetSecretRequest{
					Name: "Name",
				},
			},
			want: &pb.PasswordResponse{
				Login:    "Login",
				Password: "Password",
			},
			wantErr: false,
		},
	}
	for _, tt := range get_tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPassword(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetPassword() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestService_SaveGetCard(t *testing.T) {
	service, err := PrepareTestEnv()
	if err != nil {
		t.Log(err.Error())
		return
	}
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("userID", "123"))
	type save_args struct {
		ctx context.Context
		req *pb.SaveCardRequest
	}
	save_tests := []struct {
		name    string
		s       *Service
		args    save_args
		want    *pb.Empty
		wantErr bool
	}{
		{
			name: "save_ok",
			s:    service,
			args: save_args{
				ctx: ctx,
				req: &pb.SaveCardRequest{
					CardName:   "CardName",
					Number:     "Number",
					HolderName: "HolderName",
					Date:       "Date",
					Ccv:        "Ccv",
				},
			},
			want:    &pb.Empty{},
			wantErr: false,
		},
		{
			name: "conflict",
			s:    service,
			args: save_args{
				ctx: ctx,
				req: &pb.SaveCardRequest{
					CardName:   "CardName",
					Number:     "Number",
					HolderName: "HolderName",
					Date:       "Date",
					Ccv:        "Ccv",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range save_tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SaveCard(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.SaveCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.SaveCard() = %v, want %v", got, tt.want)
			}
		})
	}

	type get_args struct {
		ctx context.Context
		req *pb.GetSecretRequest
	}
	get_tests := []struct {
		name    string
		s       *Service
		args    get_args
		want    *pb.CardResponse
		wantErr bool
	}{
		{
			name: "get_ok",
			s:    service,
			args: get_args{
				ctx: ctx,
				req: &pb.GetSecretRequest{
					Name: "CardName",
				},
			},
			want: &pb.CardResponse{
				Number:     "Number",
				HolderName: "HolderName",
				Date:       "Date",
				Ccv:        "Ccv",
			},
		},
	}
	for _, tt := range get_tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetCard(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_SaveText(t *testing.T) {
	service, err := PrepareTestEnv()
	if err != nil {
		t.Log(err.Error())
		return
	}
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("userID", "123"))
	type save_args struct {
		ctx context.Context
		req *pb.SaveTextRequest
	}
	save_tests := []struct {
		name    string
		s       *Service
		args    save_args
		want    *pb.Empty
		wantErr bool
	}{
		{
			name: "save_ok",
			s:    service,
			args: save_args{
				ctx: ctx,
				req: &pb.SaveTextRequest{
					Name: "Name",
					Text: "Text",
				},
			},
			want:    &pb.Empty{},
			wantErr: false,
		},
		{
			name: "conflict",
			s:    service,
			args: save_args{
				ctx: ctx,
				req: &pb.SaveTextRequest{
					Name: "Name",
					Text: "Text",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range save_tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SaveText(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.SaveText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.SaveText() = %v, want %v", got, tt.want)
			}
		})
	}

	type get_args struct {
		ctx context.Context
		req *pb.GetSecretRequest
	}
	get_tests := []struct {
		name    string
		s       *Service
		args    get_args
		want    *pb.TextResponse
		wantErr bool
	}{
		{
			name: "get_ok",
			s:    service,
			args: get_args{
				ctx: ctx,
				req: &pb.GetSecretRequest{
					Name: "Name",
				},
			},
			want: &pb.TextResponse{
				Text: "Text",
			},
			wantErr: false,
		},
	}
	for _, tt := range get_tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetText(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetText() = %v, want %v", got, tt.want)
			}
		})
	}
}
