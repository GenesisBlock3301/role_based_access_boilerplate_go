package controllers

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/serializers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/services"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/utils"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/validations"
	"github.com/diebietse/gotp/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// UserHandler Implement dependency injection here
type UserHandler struct {
	UserService services.IUserServiceInterface
}

// CreateUserController UserRegistration controllers
func (u *UserHandler) CreateUserController(ctx *gin.Context) {
	var userInput serializers.RegisterSerializer
	// Validate UserInput
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(userInput)
	isErrors := validations.RegistrationInputValidate(err)
	if len(isErrors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": isErrors})
		return
	}
	_, err = u.UserService.CreateUserService(userInput)
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

	// Extract email and verify
	err := u.UserService.VerifyEmailService(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"email": "Successfully activated"})

}

func (u *UserHandler) GenerateOTP(ctx *gin.Context) {
	var UserLogin serializers.LoginSerializer
	// Validate UserInput
	if err := ctx.ShouldBindJSON(&UserLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secret, err := gotp.DecodeBase32(configurations.OtpSecret)
	totp, err := gotp.NewTOTP(secret)
	otpString, err := totp.Now()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "OTP Generation Failed!"})
		return
	}
	data := serializers.OTPSerializer{
		IsOTP: true,
		Code:  otpString,
	}
	services.SendEmail(UserLogin.Email, "OTP Code:", "", data)
	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully Otp Generated", "OTP": otpString})

}

func (u *UserHandler) VerifyOTP(ctx *gin.Context) {
	var otp serializers.VerifyOTPSerializer
	// Validate OTP Input.
	if err := ctx.ShouldBindJSON(&otp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secret, err := gotp.DecodeBase32(configurations.OtpSecret)
	totp, _ := gotp.NewTOTP(secret)
	currentTime := time.Now().Unix()
	verify, err := totp.Verify(otp.Code, int(currentTime))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "OTP Generation Failed!"})
		return
	}
	if verify == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "OTP verification failed!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully Otp verified", "OTP": verify})
}
