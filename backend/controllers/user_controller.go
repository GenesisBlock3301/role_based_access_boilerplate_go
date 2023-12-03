package controllers

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/configurations"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/serializers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/services"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/utils"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/validations"
	"github.com/diebietse/gotp/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// CreateUserController UserRegistration controllers
// @Summary Register User.
// @Schemes http https
// @Description User Registration
// @Tags User
// @Param user body serializers.RegisterSerializer true "User info"
// @Accept json
// @Produce json
// @Success 200 {string} successfully login
// @failure      400              {string}  string    "error"
// @Router /create [post]
func CreateUserController(ctx *gin.Context) {
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
	_, err = services.CreateUserService(userInput)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Email sent successfully . Check your email & verify email"})
}

// LoginController
// @Summary Login user.
// @Schemes http https
// @Description User login
// @Tags User
// @Param user body serializers.LoginSerializer true "User info"
// @Accept json
// @Produce json
// @Success 200 {string} successfully login.
// @failure      400              {string}  string    "error"
// @Router /login [post]
func LoginController(ctx *gin.Context) {
	var user serializers.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValidCredential, userID, inActiveUser := services.VerifyCredentialService(user.Email, user.Password)
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

// @BasePath /api/v1

// GetCurrentUserController
// @Summary Get authenticated user.
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} get user successfully.
// @Router / [get]
func GetCurrentUserController(ctx *gin.Context) {
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

// VerifyEmailController
// @BasePath /api/v1
// @Summary Verify email controller.
// @Schemes
// @Param token query string true "Email to be verified" Format(email)
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} successfully verify email.
// @Router /email-verify/ [get]
func VerifyEmailController(ctx *gin.Context) {
	token := ctx.Query("token")

	// Extract email and verify
	err := services.VerifyEmailService(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"email": "Successfully activated"})

}

// GenerateOTP
// @BasePath /api/v1
// @Summary Generate OTP.
// @Param user body serializers.LoginSerializer true "Generate OTP"
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} successfully generate OTP
// @Router /generate-otp/ [get]
func GenerateOTP(ctx *gin.Context) {
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

// VerifyOTP
// @BasePath /api/v1
// @Summary Verify OTP.
// @Schemes
// @Param user body serializers.VerifyOTPSerializer true "OTP info"
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} sent verify init.
// @Router /verify-otp/ [get]
func VerifyOTP(ctx *gin.Context) {
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
