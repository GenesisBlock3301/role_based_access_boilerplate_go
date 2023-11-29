package serializers

type Role struct {
	RoleName string `json:"role_name"`
}

type RoleResponse struct {
	ID       uint   `json:"id"`
	RoleName string `json:"role_name"`
}
