package service

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations/db"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/data/response"
	"strconv"
)

func GetAllRolesService(limit string, offset string) (int64, []response.RoleResponse) {
	var count int64
	db.DB.Model(&response.RoleResponse{}).Count(&count)
	roleLimit, err := strconv.Atoi(limit)
	if err != nil || roleLimit == 0 {
		roleLimit = -1
	}
	roleOffset, err := strconv.Atoi(offset)
	if err != nil || roleOffset == 0 {
		roleOffset = -1
	}
	var roles []response.RoleResponse
	db.DB.Table("role_based_access.roles").Limit(roleOffset).Offset(roleOffset).Find(&roles)
	return count, roles
}
