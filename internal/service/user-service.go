package service

import (
	"errors"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations/db"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/serializers"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash password")
	}
	return string(hash)
}

func CreateUserService(user serializers.User) (serializers.User, error) {
	if err := db.DB.Table("role_based_access.users").Where("email = ?", user.Email).First(&user).Error; err == nil {
		return user, errors.New("user already exits")
	}
	user.Password = hashAndSalt([]byte(user.Password))
	db.DB.Table("role_based_access.users").Create(&user)
	return user, nil
}
