package services

import (
	"errors"
	"fmt"
	"time"

	"diro-be/internal/models"
	"diro-be/internal/repositories"
)

// ReservationService handles reservation business logic
type ReservationService struct {
	reservationRepo *repositories.ReservationRepository
	paymentService  *PaymentService
}

// NewReservationService creates a new reservation service
func NewReservationService(reservationRepo *repositories.ReservationRepository, xenditUsername, xenditPassword string) *ReservationService {
	return &ReservationService{
		reservationRepo: reservationRepo,
		paymentService:  NewPaymentService(xenditUsername, xenditPassword),
	}
}

// CreateReservation creates a new reservation with payment
func (s *ReservationService) CreateReservation(courtID, timeslotID uint, date time.Time, customer models.XenditCustomer) (*models.Reservation, string, error) {
	// Check if the slot is still available
	available, err := s.reservationRepo.CheckSlotAvailability(courtID, timeslotID, date)
	if err != nil {
		return nil, "", err
	}
	if !available {
		return nil, "", errors.New("slot is already reserved")
	}

	// Create the reservation
	reservation := &models.Reservation{
		CourtID:       courtID,
		TimeslotID:    timeslotID,
		Date:          date,
		Status:        "pending",
		TotalPrice:    50000, // Fixed price for now
		PaymentStatus: "PENDING",
	}

	if err := s.reservationRepo.CreateReservation(reservation); err != nil {
		return nil, "", err
	}

	// Create Xendit invoice
	invoiceResp, err := s.paymentService.CreateInvoice(reservation, customer)
	if err != nil {
		// Invoice creation failed, delete reservation
		s.reservationRepo.DeleteReservation(reservation.ID)
		return nil, "", fmt.Errorf("failed to create invoice: %w", err)
	}

	// Update reservation with payment info
	reservation.PaymentID = invoiceResp.ID
	reservation.InvoiceURL = invoiceResp.InvoiceURL
	reservation.PaymentStatus = invoiceResp.Status
	if err := s.reservationRepo.UpdateReservation(reservation); err != nil {
		return nil, "", err
	}

	// Load relations
	reservation, err = s.reservationRepo.GetReservationByID(reservation.ID)
	if err != nil {
		return nil, "", err
	}

	return reservation, invoiceResp.InvoiceURL, nil
}

// UpdatePaymentStatus updates the payment status of a reservation
func (s *ReservationService) UpdatePaymentStatus(reservationID uint, paymentStatus string) error {
	reservation, err := s.reservationRepo.GetReservationByID(reservationID)
	if err != nil {
		return err
	}

	reservation.PaymentStatus = paymentStatus
	if paymentStatus == "PAID" {
		reservation.Status = "paid"
	}

	return s.reservationRepo.UpdateReservation(reservation)
}

// GetDayAvailability returns availability for a specific day
func (s *ReservationService) GetDayAvailability(date time.Time) (*models.DayAvailability, error) {
	return s.reservationRepo.GetDayAvailability(date)
}
