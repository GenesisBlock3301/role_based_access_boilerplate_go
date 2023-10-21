package main

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations/db"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/routes"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/schemas"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	initEnv()
	db.ConnectionWithDB()
}

func initEnv() {
	log.Println("Loading env setting....")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No local env file. Using global OS environment variables")
	}
	configurations.SetEnvVariable()
	schemas.SetTableName()
}

func main() {
	// Initialize `gin` router
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	routes.RootRouter(router)
	router.Run("localhost:8080")
}
