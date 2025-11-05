# DIRO Badminton Reservation Backend

A simple badminton court reservation system built with Go, GORM, and Gin.

## Features

- **Date Selection**: Get available dates for reservation
- **Timeslot Selection**: Get available timeslots for a selected date
- **Court Selection**: Get available courts for a selected date and timeslot
- **Reservation Creation**: Create reservations with payment processing
- **Payment Integration**: Mock payment gateway integration (bonus feature)

## Tech Stack

- **Go 1.24**: Programming language
- **Gin**: Web framework
- **GORM**: ORM for database operations
- **MySQL**: Database
- **golang-migrate**: Database migrations

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   ├── database/               # Database connection and migrations
│   ├── handlers/               # HTTP request handlers
│   ├── models/                 # Database models
│   └── services/               # Business logic services
├── migrations/                 # Database migration files
├── config/                     # Configuration files
├── .env.example                # Environment variables template
└── README.md                   # This file
```

## Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd diro-be
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Set up MySQL database**
   - Create a MySQL database
   - Update `.env` with your database credentials

5. **Run migrations**
   ```bash
   go run main.go
   ```
   The application will automatically run migrations on startup.

6. **Start the server**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /health` - Check server health

### Reservations
- `GET /api/reservations/dates` - Get available dates
- `GET /api/reservations/timeslots?date=2023-12-01` - Get available timeslots for a date
- `GET /api/reservations/courts?date=2023-12-01&timeslot_id=1` - Get available courts for date and timeslot
- `POST /api/reservations` - Create a new reservation
- `PUT /api/reservations/:id/confirm` - Confirm a reservation
- `PUT /api/reservations/:id/cancel` - Cancel a reservation

### Users
- `GET /api/users/:id/reservations` - Get user reservations

## Request/Response Examples

### Create Reservation
```bash
POST /api/reservations
Content-Type: application/json

{
  "user_id": 1,
  "court_id": 1,
  "timeslot_id": 1,
  "date": "2023-12-01"
}
```

### Response
```json
{
  "reservation": {
    "id": 1,
    "user_id": 1,
    "court_id": 1,
    "timeslot_id": 1,
    "date": "2023-12-01T00:00:00Z",
    "status": "confirmed",
    "total_price": 50000,
    "created_at": "2023-11-01T10:00:00Z",
    "updated_at": "2023-11-01T10:00:00Z",
    "user": {...},
    "court": {...},
    "timeslot": {...}
  }
}
```

## Database Schema

- **users**: User information
- **courts**: Badminton courts
- **timeslots**: Available time slots
- **reservations**: Court reservations

## Payment Integration

The system includes a mock payment service that simulates payment processing. In a production environment, integrate with a real payment gateway like Midtrans or Stripe.

## Development

- Run tests: `go test ./...`
- Format code: `go fmt ./...`
- Lint: Install golangci-lint and run `golangci-lint run`

## License

This project is licensed under the MIT License.