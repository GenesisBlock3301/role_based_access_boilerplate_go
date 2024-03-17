package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go_user_role/backend/configurations"
	"github.com/go_user_role/backend/configurations/db"
	"github.com/go_user_role/backend/middlewares"
	"github.com/go_user_role/backend/routes"
	"github.com/go_user_role/backend/schemas"
	_ "github.com/go_user_role/docs"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
	"net/http"
)

func init() {
	initEnv()
	db.ConnectionWithDB()
}

func initEnv() {
	log.Println("Loading env setting....")
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("No local env file. Using global OS environment variables")
	}
	configurations.SetEnvVariable()
	schemas.SetTableName()
}

// @title           Role Based Authentication API
// @version         1.0
// @description     This is golang full-featured project.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Initialize `gin` router
	router := gin.Default()

	router.Use(middlewares.JWTAuthMiddleware())

	// Create a new custom rate limiter with a limit of 5 request per second per IP
	//limiter := middlewares.NewCustomRateLimiter(5, time.Second)

	//router.Use(limiter.CustomRateLimiterMiddleware())
	// for root -> redirect to -> /swagger/index.html
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.HandleMethodNotAllowed = true
	routes.RootRouter(router)
	// Define the server address
	addr := ":8080"

	// Print the server address
	fmt.Printf("Server is running on %s\n", addr)
	fmt.Printf("Navigate to http://localhost%s\n", addr)

	// Run the Gin application
	router.Run(addr)
}
