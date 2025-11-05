package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"diro-be/internal/models"
	"diro-be/internal/services"
)

// WebhookHandler handles webhook HTTP requests
type WebhookHandler struct {
	reservationService *services.ReservationService
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(reservationService *services.ReservationService) *WebhookHandler {
	return &WebhookHandler{
		reservationService: reservationService,
	}
}

// XenditWebhook godoc
// @Summary Handle Xendit webhook
// @Description Handle payment webhook from Xendit
// @Tags webhooks
// @Accept json
// @Produce json
// @Param payload body models.XenditWebhookPayload true "Xendit webhook payload"
// @Success 200 {object} map[string]string "message: webhook received"
// @Failure 400 {object} map[string]string "error: message"
// @Failure 500 {object} map[string]string "error: message"
// @Router /api/v1/webhooks/xendit [post]
func (h *WebhookHandler) XenditWebhook(c *gin.Context) {
	var payload models.XenditWebhookPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming external_id is the reservation ID
	reservationID, err := strconv.ParseUint(payload.ExternalID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid external_id"})
		return
	}

	// Update reservation status based on payment status
	err = h.reservationService.UpdatePaymentStatus(uint(reservationID), payload.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "webhook received"})
}
