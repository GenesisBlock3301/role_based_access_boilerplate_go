package request

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string `json:"password"`
}
