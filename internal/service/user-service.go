package service

import (
	"errors"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations/db"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/serializers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(user serializers.User) (serializers.User, error) {
	if err := db.DB.Table("role_based_access.users").Where("email = ?", user.Email).First(&user).Error; err == nil {
		return user, errors.New("user already exits")
	}
	user.Password = utils.HashAndSalt([]byte(user.Password))
	db.DB.Table("role_based_access.users").Create(&user)
	return user, nil
}

func VerifyCredentialService(email string, password string) (bool, uint) {
	user, err := FindByEmail(email)
	if err != nil {
		return false, 0
	}
	return comparePassword([]byte(user.Password), []byte(password)), user.ID
}

// FindByEmail Find user by email
func FindByEmail(email string) (serializers.LoginUserSerializer, error) {
	user := serializers.LoginUserSerializer{}
	if err := db.DB.Table("role_based_access.users").Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Compare User password and input password
func comparePassword(hashedPass []byte, inputPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPass, inputPass)
	if err != nil {
		return false
	}
	return true
}

func GetUserById(id uint) (serializers.LoginUserSerializer, error) {
	var user serializers.LoginUserSerializer
	if err := db.DB.Table("role_based_access.users").First(&user, id).Error; err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}
