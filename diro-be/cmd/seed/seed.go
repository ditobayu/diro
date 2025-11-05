// Package main provides database seeding utilities
// This command-line tool populates the database with initial data
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"diro-be/internal/config"
	"diro-be/internal/models"
)

// main is the entry point for the seeding command
func main() {
	var action string
	flag.StringVar(&action, "action", "", "Seeding action: seed (populate data), clear (remove all data)")
	flag.Parse()

	if action == "" {
		fmt.Println("Usage:")
		fmt.Println("  go run cmd/seed/seed.go -action=seed   # Populate database with initial data")
		fmt.Println("  go run cmd/seed/seed.go -action=clear  # Clear all seeded data")
		os.Exit(1)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	dsn := cfg.GetDBDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database successfully")

	switch action {
	case "seed":
		if err := seedDatabase(db); err != nil {
			log.Fatal("Failed to seed database:", err)
		}
		fmt.Println("Database seeded successfully")

	case "clear":
		if err := clearDatabase(db); err != nil {
			log.Fatal("Failed to clear database:", err)
		}
		fmt.Println("Database cleared successfully")

	default:
		log.Fatal("Invalid action. Use: seed or clear")
	}
}

// seedDatabase populates the database with initial data
func seedDatabase(db *gorm.DB) error {
	// Seed courts
	courts := []models.Court{
		{
			Name:        "Lapangan A",
			Description: "Lapangan badminton utama dengan pencahayaan LED",
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Lapangan B",
			Description: "Lapangan badminton dengan lantai sintetis",
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Lapangan C",
			Description: "Lapangan badminton indoor dengan AC",
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Lapangan D",
			Description: "Lapangan badminton outdoor",
			IsActive:    false, // Inactive for testing
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := db.Create(&courts).Error; err != nil {
		return fmt.Errorf("failed to seed courts: %w", err)
	}
	fmt.Println("Seeded 4 courts")

	// Seed timeslots
	timeslots := []models.Timeslot{
		{
			StartTime: "08:00",
			EndTime:   "09:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "09:00",
			EndTime:   "10:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "10:00",
			EndTime:   "11:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "11:00",
			EndTime:   "12:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "13:00",
			EndTime:   "14:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "14:00",
			EndTime:   "15:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "15:00",
			EndTime:   "16:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "16:00",
			EndTime:   "17:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "18:00",
			EndTime:   "19:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "19:00",
			EndTime:   "20:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "20:00",
			EndTime:   "21:00",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			StartTime: "21:00",
			EndTime:   "22:00",
			IsActive:  false, // Inactive for testing
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(&timeslots).Error; err != nil {
		return fmt.Errorf("failed to seed timeslots: %w", err)
	}
	fmt.Println("Seeded 12 timeslots")

	// Get the first court and timeslot IDs to use for reservations
	var firstCourt models.Court
	var firstTimeslot models.Timeslot
	if err := db.First(&firstCourt).Error; err != nil {
		return fmt.Errorf("failed to get first court: %w", err)
	}
	if err := db.First(&firstTimeslot).Error; err != nil {
		return fmt.Errorf("failed to get first timeslot: %w", err)
	}

	// Seed reservations (using actual IDs from seeded data)
	tomorrow := time.Now().AddDate(0, 0, 1)
	reservations := []models.Reservation{
		{
			CourtID:    firstCourt.ID,    // First court
			TimeslotID: firstTimeslot.ID, // First timeslot
			Date:       tomorrow,
			Status:     "confirmed",
			TotalPrice: 50000,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			CourtID:    firstCourt.ID,
			TimeslotID: firstTimeslot.ID + 1, // Second timeslot
			Date:       tomorrow,
			Status:     "pending",
			TotalPrice: 50000,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			CourtID:    firstCourt.ID + 1,    // Second court
			TimeslotID: firstTimeslot.ID + 2, // Third timeslot
			Date:       tomorrow,
			Status:     "confirmed",
			TotalPrice: 50000,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			CourtID:    firstCourt.ID + 2,           // Third court
			TimeslotID: firstTimeslot.ID + 4,        // Fifth timeslot
			Date:       time.Now().AddDate(0, 0, 2), // Day after tomorrow
			Status:     "cancelled",
			TotalPrice: 50000,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			CourtID:    firstCourt.ID + 1,           // Second court
			TimeslotID: firstTimeslot.ID + 6,        // Seventh timeslot
			Date:       time.Now().AddDate(0, 0, 3), // 3 days from now
			Status:     "confirmed",
			TotalPrice: 50000,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	if err := db.Create(&reservations).Error; err != nil {
		return fmt.Errorf("failed to seed reservations: %w", err)
	}
	fmt.Println("Seeded 5 reservations")

	return nil
}

// clearDatabase removes all seeded data
func clearDatabase(db *gorm.DB) error {
	// Clear in reverse order due to foreign key constraints
	if err := db.Exec("DELETE FROM reservations").Error; err != nil {
		return fmt.Errorf("failed to clear reservations: %w", err)
	}
	fmt.Println("Cleared reservations")

	if err := db.Exec("DELETE FROM timeslots").Error; err != nil {
		return fmt.Errorf("failed to clear timeslots: %w", err)
	}
	fmt.Println("Cleared timeslots")

	if err := db.Exec("DELETE FROM courts").Error; err != nil {
		return fmt.Errorf("failed to clear courts: %w", err)
	}
	fmt.Println("Cleared courts")

	return nil
}
