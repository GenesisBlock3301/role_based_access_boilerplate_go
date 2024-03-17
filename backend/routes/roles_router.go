package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go_user_role/backend/controllers"
	"github.com/go_user_role/backend/services"
)

func RoleRouter(roleRouter *gin.RouterGroup) {
	roleService := &services.RoleService{}
	roleController := controllers.NewRoleController(*roleService)
	roleRouter.POST("/create", roleController.CreateRoleController)
	roleRouter.GET("/list", roleController.GetALLRoleController)
}
