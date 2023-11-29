package routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/controllers"
	"github.com/gin-gonic/gin"
)

func RoleRouter(roleRouter *gin.RouterGroup) {
	roleRouter.POST("/create", controllers.CreateRoleController)
	roleRouter.GET("/list", controllers.GetALLRoleController)
}
