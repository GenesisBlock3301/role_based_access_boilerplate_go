package response

import "time"

type RoleResponse struct {
	ID        uint       `json:"id"`
	RoleName  string     `json:"role_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
