package memberships

import (
	"errors"
	"fmt"

	"github.com/Bagussurya12/catalog-music-simple/pkg/jwt"
	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(request memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(request.Email, "", uint(0))
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get user from database")
		return "", err
	}

	if userDetail == nil {
		return "", errors.New("email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password))

	if err != nil {
		fmt.Println("test case password not match")
		return "", errors.New("password or email not found")
	}

	accessToken, err := jwt.CreateToken(int64(userDetail.ID), userDetail.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed create jwt token")
		return "", err
	}

	return accessToken, nil
}
