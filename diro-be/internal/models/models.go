package models

import (
	"time"
)

// Court represents a badminton court
type Court struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Timeslot represents available time slots
type Timeslot struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StartTime string    `json:"start_time" gorm:"not null"` // Format: "HH:MM"
	EndTime   string    `json:"end_time" gorm:"not null"`   // Format: "HH:MM"
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Reservation represents a booking reservation
type Reservation struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	CourtID       uint      `json:"court_id" gorm:"not null"`
	TimeslotID    uint      `json:"timeslot_id" gorm:"not null"`
	Date          time.Time `json:"date" gorm:"type:date;not null"`
	Status        string    `json:"status" gorm:"default:'pending'"` // pending, confirmed, cancelled, paid
	TotalPrice    float64   `json:"total_price" gorm:"default:0"`
	PaymentID     string    `json:"payment_id" gorm:"default:''"`     // Xendit invoice ID
	InvoiceURL    string    `json:"invoice_url" gorm:"default:''"`    // Xendit invoice URL
	PaymentStatus string    `json:"payment_status" gorm:"default:''"` // PENDING, PAID, FAILED, EXPIRED
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// Relations
	Court    Court    `json:"court" gorm:"foreignKey:CourtID"`
	Timeslot Timeslot `json:"timeslot" gorm:"foreignKey:TimeslotID"`
}

// DayAvailability represents availability for a specific day
type DayAvailability struct {
	Date   string              `json:"date"`
	Courts []CourtAvailability `json:"courts"`
}

// CourtAvailability represents availability for a specific court on a day
type CourtAvailability struct {
	Court     Court                `json:"court"`
	Timeslots []TimeslotWithStatus `json:"timeslots"`
}

// TimeslotWithStatus represents a timeslot with its booking status
type TimeslotWithStatus struct {
	Timeslot Timeslot `json:"timeslot"`
	IsBooked bool     `json:"is_booked"`
}

// XenditWebhookPayload represents the payload from Xendit webhook
type XenditWebhookPayload struct {
	ID                 string              `json:"id"`
	Items              []XenditWebhookItem `json:"items"`
	Amount             int                 `json:"amount"`
	Status             string              `json:"status"`
	Created            string              `json:"created"`
	IsHigh             bool                `json:"is_high"`
	PaidAt             *string             `json:"paid_at,omitempty"`
	Updated            string              `json:"updated"`
	UserID             string              `json:"user_id"`
	Currency           string              `json:"currency"`
	PaymentID          string              `json:"payment_id"`
	Description        string              `json:"description"`
	ExternalID         string              `json:"external_id"`
	PaidAmount         int                 `json:"paid_amount"`
	EwalletType        string              `json:"ewallet_type"`
	MerchantName       string              `json:"merchant_name"`
	PaymentMethod      string              `json:"payment_method"`
	PaymentChannel     string              `json:"payment_channel"`
	PaymentMethodID    string              `json:"payment_method_id"`
	FailureRedirectURL string              `json:"failure_redirect_url"`
	SuccessRedirectURL string              `json:"success_redirect_url"`
}

// XenditWebhookItem represents an item in the webhook payload
type XenditWebhookItem struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}

// XenditInvoiceRequest represents the request payload to create Xendit invoice
type XenditInvoiceRequest struct {
	ExternalID         string                 `json:"external_id"`
	Amount             float64                `json:"amount"`
	Description        string                 `json:"description"`
	InvoiceDuration    int                    `json:"invoice_duration"`
	Customer           XenditCustomer         `json:"customer"`
	SuccessRedirectURL string                 `json:"success_redirect_url"`
	FailureRedirectURL string                 `json:"failure_redirect_url"`
	Currency           string                 `json:"currency"`
	Items              []XenditInvoiceItem    `json:"items"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}

// XenditCustomer represents customer info for Xendit invoice
type XenditCustomer struct {
	GivenNames   string `json:"given_names"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
}

// XenditInvoiceItem represents an item in the invoice
type XenditInvoiceItem struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
	URL      string  `json:"url"`
}

// XenditInvoiceResponse represents the response from Xendit invoice creation
type XenditInvoiceResponse struct {
	ID                        string                 `json:"id"`
	ExternalID                string                 `json:"external_id"`
	UserID                    string                 `json:"user_id"`
	Status                    string                 `json:"status"`
	MerchantName              string                 `json:"merchant_name"`
	MerchantProfilePictureURL string                 `json:"merchant_profile_picture_url"`
	Amount                    float64                `json:"amount"`
	Description               string                 `json:"description"`
	ExpiryDate                string                 `json:"expiry_date"`
	InvoiceURL                string                 `json:"invoice_url"`
	AvailableBanks            []XenditAvailableBank  `json:"available_banks"`
	AvailableRetailOutlets    []XenditRetailOutlet   `json:"available_retail_outlets"`
	AvailableEwallets         []XenditEwallet        `json:"available_ewallets"`
	AvailableQRCodes          []XenditQRCode         `json:"available_qr_codes"`
	AvailableDirectDebits     []XenditDirectDebit    `json:"available_direct_debits"`
	AvailablePaylaters        []interface{}          `json:"available_paylaters"`
	ShouldExcludeCreditCard   bool                   `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                   `json:"should_send_email"`
	SuccessRedirectURL        string                 `json:"success_redirect_url"`
	FailureRedirectURL        string                 `json:"failure_redirect_url"`
	Created                   string                 `json:"created"`
	Updated                   string                 `json:"updated"`
	Currency                  string                 `json:"currency"`
	Items                     []XenditInvoiceItem    `json:"items"`
	Customer                  XenditCustomer         `json:"customer"`
	Metadata                  map[string]interface{} `json:"metadata"`
}

// XenditAvailableBank represents available bank for payment
type XenditAvailableBank struct {
	BankCode          string `json:"bank_code"`
	CollectionType    string `json:"collection_type"`
	TransferAmount    int    `json:"transfer_amount"`
	BankBranch        string `json:"bank_branch"`
	AccountHolderName string `json:"account_holder_name"`
	IdentityAmount    int    `json:"identity_amount"`
}

// XenditRetailOutlet represents available retail outlet
type XenditRetailOutlet struct {
	RetailOutletName string `json:"retail_outlet_name"`
}

// XenditEwallet represents available e-wallet
type XenditEwallet struct {
	EwalletType string `json:"ewallet_type"`
}

// XenditQRCode represents available QR code
type XenditQRCode struct {
	QRCodeType string `json:"qr_code_type"`
}

// XenditDirectDebit represents available direct debit
type XenditDirectDebit struct {
	DirectDebitType string `json:"direct_debit_type"`
}
