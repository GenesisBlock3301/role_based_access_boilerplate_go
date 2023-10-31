package services

import (
	"errors"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations/db"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/schemas"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/serializers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/utils"
	"github.com/dgrijalva/jwt-go"
)

type IUserServiceInterface interface {
	CreateUserService(user serializers.RegisterSerializer) (bool, error)
	VerifyCredentialService(email string, password string) (bool, uint, bool)
	VerifyEmailService(token string) error
}

type UserService struct{}

func (u *UserService) CreateUserService(user serializers.RegisterSerializer) (bool, error) {
	err := db.DB.Table(schemas.Users).Where("email = ?", user.Email).First(&user).Error
	if err == nil {
		return false, errors.New("user already exits")
	}
	user.Password = utils.HashAndSalt([]byte(user.Password))
	db.DB.Table(schemas.Users).Create(&user)
	token, _ := utils.GenerateEmailToken(user.Email)
	err = SendEmail(user.Email, "Verify Your Email", token, serializers.OTPSerializer{})
	if err != nil {
		return false, errors.New("sending email failed")
	}
	return true, nil
}

func (u *UserService) VerifyCredentialService(email string, password string) (bool, uint, bool) {
	user, err := utils.FindByEmail(email)
	if user.IsActive == 0 {
		return false, 0, true
	}
	if err != nil {
		return false, 0, false
	}
	return utils.ComparePassword([]byte(user.Password), []byte(password)), user.ID, false
}

func (u *UserService) VerifyEmailService(token string) error {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(configurations.EmailTokenSecret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return err
		}
		return err
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok || !parseToken.Valid {
		return err
	}
	email := claims["email"]
	err = utils.UpdateUserVerificationStatus(email)
	if err != nil {
		return err
	}
	return nil
}
