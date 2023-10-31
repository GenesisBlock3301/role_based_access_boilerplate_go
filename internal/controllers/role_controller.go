package controllers

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/serializers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleHandler struct {
	RoleService services.IRoleService
}

func (r *RoleHandler) CreateRoleController(ctx *gin.Context) {
	var roleInput serializers.Role
	if err := ctx.ShouldBindJSON(&roleInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := r.RoleService.CreateRoleService(&roleInput)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Successfully role created!"})
}

func (r *RoleHandler) GetALLRoleController(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	roleCount, roles := r.RoleService.GetAllRolesService(limit, offset)
	ctx.JSON(200, gin.H{
		"totalRole": roleCount,
		"roles":     roles,
	})
}
