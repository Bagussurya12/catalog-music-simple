package memberships

import (
	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"github.com/gin-gonic/gin"
)

type service interface {
	Signup(request memberships.SignUpRequest) error
}
type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}
