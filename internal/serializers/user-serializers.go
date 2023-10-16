package serializers

type User struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsActive     int    `json:"is_active"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	Name         string `json:"name"`
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
