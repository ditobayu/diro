package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "diro-be/docs" // Import generated docs

	"diro-be/internal/config"
	"diro-be/internal/database"
	"diro-be/internal/repositories"
	"diro-be/internal/routes"
)

// @title Diro API
// @version 1.0
// @description API untuk sistem reservasi lapangan olahraga Diro
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Initialize Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.Default())

	// Initialize repositories
	reservationRepo := repositories.NewReservationRepository(database.DB)

	// Setup routes
	routes.SetupRoutes(router, cfg, reservationRepo)

	// Swagger routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	log.Println("Server starting on port 8080...")
	log.Println("Swagger documentation available at: http://localhost:8080/swagger/index.html")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
