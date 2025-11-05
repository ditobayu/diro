package repositories

import (
	"time"

	"gorm.io/gorm"

	"diro-be/internal/models"
)

// ReservationRepository handles database operations for reservations
type ReservationRepository struct {
	db *gorm.DB
}

// NewReservationRepository creates a new reservation repository
func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

// CreateReservation creates a new reservation in the database
func (r *ReservationRepository) CreateReservation(reservation *models.Reservation) error {
	return r.db.Create(reservation).Error
}

// GetReservationByID gets a reservation by ID with relations
func (r *ReservationRepository) GetReservationByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation
	err := r.db.Preload("Court").Preload("Timeslot").First(&reservation, id).Error
	return &reservation, err
}

// UpdateReservation updates a reservation
func (r *ReservationRepository) UpdateReservation(reservation *models.Reservation) error {
	return r.db.Save(reservation).Error
}

// DeleteReservation deletes a reservation
func (r *ReservationRepository) DeleteReservation(id uint) error {
	return r.db.Delete(&models.Reservation{}, id).Error
}

// CheckSlotAvailability checks if a slot is available
func (r *ReservationRepository) CheckSlotAvailability(courtID, timeslotID uint, date time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&models.Reservation{}).
		Where("court_id = ? AND timeslot_id = ? AND date = ? AND status = ?",
			courtID, timeslotID, date.Format("2006-01-02"), "paid").
		Count(&count).Error
	return count == 0, err
}

// GetDayAvailability returns availability for a specific day
func (r *ReservationRepository) GetDayAvailability(date time.Time) (*models.DayAvailability, error) {
	var courts []models.Court
	if err := r.db.Where("is_active = ?", true).Find(&courts).Error; err != nil {
		return nil, err
	}

	var courtAvailabilities []models.CourtAvailability

	for _, court := range courts {
		// Get all active timeslots
		var allTimeslots []models.Timeslot
		if err := r.db.Where("is_active = ?", true).Find(&allTimeslots).Error; err != nil {
			return nil, err
		}

		// Get reserved timeslot IDs for this court and date
		var reservedTimeslotIDs []uint
		r.db.Model(&models.Reservation{}).
			Where("court_id = ? AND date = ? AND status = ?",
				court.ID, date.Format("2006-01-02"), "paid").
			Pluck("timeslot_id", &reservedTimeslotIDs)

		// Create timeslots with status
		var timeslotsWithStatus []models.TimeslotWithStatus
		for _, ts := range allTimeslots {
			isBooked := false
			for _, reservedID := range reservedTimeslotIDs {
				if ts.ID == reservedID {
					isBooked = true
					break
				}
			}
			timeslotsWithStatus = append(timeslotsWithStatus, models.TimeslotWithStatus{
				Timeslot: ts,
				IsBooked: isBooked,
			})
		}

		courtAvailability := models.CourtAvailability{
			Court:     court,
			Timeslots: timeslotsWithStatus,
		}
		courtAvailabilities = append(courtAvailabilities, courtAvailability)
	}

	dayAvailability := &models.DayAvailability{
		Date:   date.Format("2006-01-02"),
		Courts: courtAvailabilities,
	}

	return dayAvailability, nil
}
