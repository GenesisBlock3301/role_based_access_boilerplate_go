package controller

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/service"
	"github.com/gin-gonic/gin"
)

func GetALLRoleController(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	roleCount, roles := service.GetAllRolesService(limit, offset)
	ctx.JSON(200, gin.H{
		"totalRole": roleCount,
		"roles":     roles,
	})
}
