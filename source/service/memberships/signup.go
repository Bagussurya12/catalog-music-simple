package memberships

import (
	"errors"

	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Signup(request memberships.SignUpRequest) error {

	userExist, err := s.repository.GetUser(request.Email, request.Username, 0)

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Failed to Get User")

		return err
	}

	if userExist != nil {
		return errors.New("email or username already exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error().Err(err).Msg("Failed to Hash Password")
		return err
	}
	model := memberships.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  string(pass),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}

	return s.repository.CreateUser(model)
}
