package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/AlexTerra21/gophkeeper/pb"
)

// Тестирование пары Register-Login
func TestService_Register_Login(t *testing.T) {
	service, err := PrepareTestEnv()
	if err != nil {
		t.Log(err.Error())
		return
	}
	type reg_args struct {
		ctx context.Context
		req *pb.RegisterRequest
	}
	reg_tests := []struct {
		name    string
		s       *Service
		args    reg_args
		want    *pb.Empty
		wantErr bool
	}{
		{
			name: "reg_ok",
			s:    service,
			args: reg_args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					Username: "Username",
					Password: "Password",
				},
			},
			want:    &pb.Empty{},
			wantErr: false,
		},
		{
			name: "reg_conflict",
			s:    service,
			args: reg_args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					Username: "Username",
					Password: "Password",
				},
			},
			want:    &pb.Empty{},
			wantErr: true,
		},
	}
	for _, tt := range reg_tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Register() = %v, want %v", got, tt.want)
			}
		})
	}

	type login_args struct {
		ctx context.Context
		req *pb.LoginRequest
	}
	login_tests := []struct {
		name    string
		s       *Service
		args    login_args
		want    *pb.Empty
		wantErr bool
	}{
		{
			name: "login_ok",
			s:    service,
			args: login_args{
				ctx: context.Background(),
				req: &pb.LoginRequest{
					Username: "Username",
					Password: "Password",
				},
			},
			want:    &pb.Empty{},
			wantErr: false,
		},
	}
	for _, tt := range login_tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
