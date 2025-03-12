package memberships

import (
	"database/sql"
	"testing"

	"github.com/Bagussurya12/catalog-music-simple/source/configs"
	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"go.uber.org/mock/gomock"
)

func Test_service_Signup(t *testing.T) {

	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.SignUpRequest
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
				request: memberships.SignUpRequest{
					Email:    "test@mail.com",
					Username: "testUser",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, sql.ErrNoRows)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.Signup(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
