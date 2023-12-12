package routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/controllers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/services"
	"github.com/gin-gonic/gin"
)

func RoleRouter(roleRouter *gin.RouterGroup) {
	roleService := &services.RoleService{}
	roleController := controllers.NewRoleController(*roleService)
	roleRouter.POST("/create", roleController.CreateRoleController)
	roleRouter.GET("/list", roleController.GetALLRoleController)
}
