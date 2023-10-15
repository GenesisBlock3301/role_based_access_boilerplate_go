package serializers

type User struct {
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type LoginUserSerializer struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive int    `json:"is_active"`
}
