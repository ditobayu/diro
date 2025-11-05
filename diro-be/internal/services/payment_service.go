package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"diro-be/internal/models"
)

// PaymentService handles payment processing
type PaymentService struct {
	xenditUsername string
	xenditPassword string
	xenditBaseURL  string
}

// NewPaymentService creates a new payment service
func NewPaymentService(username, password string) *PaymentService {
	return &PaymentService{
		xenditUsername: username,
		xenditPassword: password,
		xenditBaseURL:  "https://api.xendit.co",
	}
}

// CreateInvoice creates a payment invoice via Xendit
func (s *PaymentService) CreateInvoice(reservation *models.Reservation, customer models.XenditCustomer) (*models.XenditInvoiceResponse, error) {
	request := models.XenditInvoiceRequest{
		ExternalID:         strconv.Itoa(int(reservation.ID)),
		Amount:             reservation.TotalPrice,
		Description:        fmt.Sprintf("Reservation for %s at %s", reservation.Court.Name, reservation.Date.Format("2006-01-02")),
		InvoiceDuration:    86400, // 24 hours
		Customer:           customer,
		SuccessRedirectURL: "http://localhost:3000/success",
		FailureRedirectURL: "http://localhost:3000/failed",
		Currency:           "IDR",
		Items: []models.XenditInvoiceItem{
			{
				Name:     fmt.Sprintf("Court %s - %s to %s", reservation.Court.Name, reservation.Timeslot.StartTime, reservation.Timeslot.EndTime),
				Quantity: 1,
				Price:    reservation.TotalPrice,
				Category: "Sports",
				URL:      "http://localhost:3000",
			},
		},
		Metadata: map[string]interface{}{
			"reservation_id": reservation.ID,
			"court_id":       reservation.CourtID,
			"date":           reservation.Date.Format("2006-01-02"),
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", s.xenditBaseURL+"/v2/invoices", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	auth := s.xenditUsername + ":" + s.xenditPassword
	// log usernam and password

	fmt.Printf("Username: %s, Password: %s\n", s.xenditUsername, s.xenditPassword)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)
	req.Header.Set("X-API-VERSION", "2020-02-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("xendit API error: %s", string(body))
	}

	var invoiceResp models.XenditInvoiceResponse
	if err := json.Unmarshal(body, &invoiceResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &invoiceResp, nil
}
