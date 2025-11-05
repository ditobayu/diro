package routes

import (
	"net/http"

	"diro-be/internal/config"
	"diro-be/internal/handlers"
	"diro-be/internal/repositories"
	"diro-be/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config, reservationRepo *repositories.ReservationRepository) {
	// CORS middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// Initialize services
	reservationService := services.NewReservationService(reservationRepo, cfg.XenditUsername, cfg.XenditPassword)

	// Initialize handlers
	reservationHandler := handlers.NewReservationHandler(reservationService)
	webhookHandler := handlers.NewWebhookHandler(reservationService)

	// API routes
	api := router.Group("/api/v1")
	{
		// Reservation routes
		reservations := api.Group("/reservations")
		{
			reservations.GET("/availability", reservationHandler.GetDayAvailability)
			reservations.POST("", reservationHandler.CreateReservation)
		}

		// Webhook routes
		webhooks := api.Group("/webhooks")
		{
			webhooks.POST("/xendit", webhookHandler.XenditWebhook)
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})
}
