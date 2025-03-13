package memberships

import (
	"fmt"
	"testing"

	"github.com/Bagussurya12/catalog-music-simple/source/configs"
	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.LoginRequest{
					Email:    "baguss@mail.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "baguss@mail.com",
					Password: "$2a$10$56gUcY3yScGdrsjooEsHT.x/nCryB3NdHPfw6zgN0ppyd4WAwlSR2",
					Username: "bagussurya",
				}, nil)
			},
		},
		{
			name: "fail when password not match",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@mail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "password not match",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@mail.com",
					Password: "wrong password test",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "abgs",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				fmt.Printf("test case: %+v", tt.name)
				assert.NotEmpty(t, got)
			} else {
				fmt.Printf("\ntest case: %s \n", tt.name)
				assert.Empty(t, got)
			}
		})
	}
}
