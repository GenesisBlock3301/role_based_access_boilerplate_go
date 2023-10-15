package user_controller

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/serializers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/service"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/service/user_services"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUserController UserRegistration controller
func CreateUserController(ctx *gin.Context) {
	var userInput serializers.User
	// Validate UserInput
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := user_services.CreateUserService(userInput)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Email sent successfully . Check your email & verify email"})
}

func LoginController(ctx *gin.Context) {
	var user serializers.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValidCredential, userID, inActiveUser := user_services.VerifyCredentialService(user.Email, user.Password)
	if inActiveUser {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Inactive user_controller-services try to active your account first."})
		return
	}
	if isValidCredential {
		token, refresh, err := service.GenerateTokenPair(userID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credential"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token":         token,
			"refresh_token": refresh,
		})
	}
}

func GetCurrentUserController(ctx *gin.Context) {
	userId, err := service.ExtractTokenID(ctx)
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

func VerifyEmailController(ctx *gin.Context) {
	token := ctx.Query("token")
	err := user_services.VerifyEmailService(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"email": "Successfully activated"})

}
