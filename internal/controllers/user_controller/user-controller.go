package user_controller

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/serializers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/services"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/services/user_services"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserService user_services.UserService
}

// CreateUserController UserRegistration controllers
func (u *UserHandler) CreateUserController(ctx *gin.Context) {
	var userInput serializers.User
	// Validate UserInput
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := u.UserService.CreateUserService(userInput)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Email sent successfully . Check your email & verify email"})
}

func (u *UserHandler) LoginController(ctx *gin.Context) {
	var user serializers.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValidCredential, userID, inActiveUser := u.UserService.VerifyCredentialService(user.Email, user.Password)
	if inActiveUser {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Inactive user try to active your account first."})
		return
	}
	if isValidCredential {
		token, refresh, err := services.GenerateTokenPair(userID)
		user, err = utils.AddingUserTokens(user, token, refresh)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credential"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user": user.UserResponse(),
		})
	}
}

func (u *UserHandler) GetCurrentUserController(ctx *gin.Context) {
	userId, err := services.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := utils.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})

}

func (u *UserHandler) VerifyEmailController(ctx *gin.Context) {
	token := ctx.Query("token")
	err := u.UserService.VerifyEmailService(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"email": "Successfully activated"})

}
