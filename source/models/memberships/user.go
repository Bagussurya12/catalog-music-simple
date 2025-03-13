package memberships

import "gorm.io/gorm"

type (
	User struct {
		Email     string `gorm:"unique;not null"`
		Username  string `gorm:"unique"`
		Password  string `gorm:"not null"`
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
		gorm.Model
	}
)

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
