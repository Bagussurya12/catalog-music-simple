package memberships

import (
	models "github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships

type Service interface {
	Signup(request models.SignUpRequest) error
	Login(request models.LoginRequest) (string, error)
}
type Handler struct {
	*gin.Engine
	service Service
}

func NewHandler(api *gin.Engine, service Service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)
}
