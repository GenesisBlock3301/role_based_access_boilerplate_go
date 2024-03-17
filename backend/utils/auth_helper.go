package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/go_user_role/backend/configurations"
	"github.com/go_user_role/backend/configurations/db"
	"github.com/go_user_role/backend/schemas"
	"github.com/go_user_role/backend/serializers"
	"golang.org/x/crypto/bcrypt"
)

// ComparePassword User password and input password
func ComparePassword(hashedPass []byte, inputPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPass, inputPass)
	if err != nil {
		return false
	}
	return true
}

func GetUserById(id uint) (serializers.LoginUserSerializer, error) {
	var user serializers.LoginUserSerializer
	if err := db.DB.Table(schemas.Users).First(&user, id).Error; err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

// FindByEmail Find user_controller-services by email
func FindByEmail(email string) (serializers.LoginUserSerializer, error) {
	user := serializers.LoginUserSerializer{}
	// Find user based on email
	if err := db.DB.Table(schemas.Users).Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GenerateEmailToken Generate Token for email verification, this token will part of query parameter.
func GenerateEmailToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, err := token.SignedString([]byte(configurations.EmailTokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// UpdateUserVerificationStatus After click on Emailed verification's link, then activate user's account.
func UpdateUserVerificationStatus(email interface{}) error {
	var user serializers.LoginUserSerializer
	err := db.DB.Table(schemas.Users).Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}
	if user.IsActive == 0 {
		user.IsActive = 1
	}
	err = db.DB.Table(schemas.Users).Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func AddingUserTokens(user serializers.User, token, refreshToken string) (serializers.User, error) {
	if err := db.DB.Table(schemas.Users).Where("email = ?", user.Email).First(&user).Error; err != nil {
		return user, err
	}

	user.Token = token
	user.RefreshToken = refreshToken
	err := db.DB.Table(schemas.Users).Where("email = ?", user.Email).Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
