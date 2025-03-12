package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name              string
		mockFn            func()
		expectedStatuCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().Signup(memberships.SignUpRequest{
					Email:    "test@mail.com",
					Username: "testuser",
					Password: "password",
				}).Return(nil)
			},
			expectedStatuCode: 201,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endPoint := `/memberships/signup`
			data := memberships.SignUpRequest{
				Email:    "test@mail.com",
				Username: "testuser",
				Password: "password",
			}

			val, err := json.Marshal(data)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endPoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatuCode, w.Code)
		})
	}
}
