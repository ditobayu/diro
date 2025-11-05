package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"diro-be/internal/models"
	"diro-be/internal/services"
)

// ReservationHandler handles reservation HTTP requests
type ReservationHandler struct {
	reservationService *services.ReservationService
}

// NewReservationHandler creates a new reservation handler
func NewReservationHandler(reservationService *services.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}

// CreateReservation godoc
// @Summary Create a new reservation
// @Description Create a new reservation for a court at specific date and timeslot
// @Tags reservations
// @Accept json
// @Produce json
// @Param reservation body object true "Reservation data"
// @Success 201 {object} map[string]interface{} "reservation: object, invoice_url: string"
// @Failure 400 {object} map[string]string "error: message"
// @Router /api/reservations [post]
func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var req struct {
		CourtID    uint   `json:"court_id" binding:"required"`
		TimeslotID uint   `json:"timeslot_id" binding:"required"`
		Date       string `json:"date" binding:"required"`
		Customer   struct {
			GivenNames   string `json:"given_names" binding:"required"`
			Surname      string `json:"surname"`
			Email        string `json:"email" binding:"required"`
			MobileNumber string `json:"mobile_number" binding:"required"`
		} `json:"customer" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	customer := models.XenditCustomer{
		GivenNames:   req.Customer.GivenNames,
		Surname:      req.Customer.Surname,
		Email:        req.Customer.Email,
		MobileNumber: req.Customer.MobileNumber,
	}

	reservation, invoiceURL, err := h.reservationService.CreateReservation(req.CourtID, req.TimeslotID, date, customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reservation": reservation,
		"invoice_url": invoiceURL,
	})
}

// GetDayAvailability godoc
// @Summary Get day availability
// @Description Get availability for a specific day including courts and available timeslots
// @Tags reservations
// @Accept json
// @Produce json
// @Param date query string true "Date in YYYY-MM-DD format"
// @Success 200 {object} models.DayAvailability
// @Failure 400 {object} map[string]string "error: message"
// @Failure 500 {object} map[string]string "error: message"
// @Router /api/reservations/availability [get]
func (h *ReservationHandler) GetDayAvailability(c *gin.Context) {
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date parameter is required"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	availability, err := h.reservationService.GetDayAvailability(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, availability)
}
