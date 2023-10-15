package role_controller

//
//import (
//	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/services"
//	"github.com/gin-gonic/gin"
//)
//
//func GetALLRoleController(ctx *gin.Context) {
//	limit := ctx.Query("limit")
//	offset := ctx.Query("offset")
//	roleCount, roles := services.GetAllRolesService(limit, offset)
//	ctx.JSON(200, gin.H{
//		"totalRole": roleCount,
//		"roles":     roles,
//	})
//}
