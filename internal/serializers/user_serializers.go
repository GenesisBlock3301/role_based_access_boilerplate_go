package serializers

type RegisterSerializer struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsActive     int    `json:"is_active"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type OTPSerializer struct {
	IsOTP bool   `json:"is_otp"`
	Code  string `json:"code"`
}

type VerifyOTPSerializer struct {
	Code string `json:"code"`
}

type LoginSerializer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Name         string `json:"name""`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *User) UserResponse() *LoginResponse {
	return &LoginResponse{
		Name:         u.Name,
		Email:        u.Email,
		Token:        u.Token,
		RefreshToken: u.RefreshToken,
	}
}

type LoginUserSerializer struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive int    `json:"is_active"`
}
