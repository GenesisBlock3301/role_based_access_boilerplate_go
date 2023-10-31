package routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RoleRouter(roleRouter *gin.RouterGroup) {
	roleHandler := controllers.RoleHandler{}
	roleRouter.POST("/create", roleHandler.CreateRoleController)
	roleRouter.GET("/list", roleHandler.GetALLRoleController)
}
