# DIRO - Badminton Court Reservation System

Sistem reservasi lapangan badminton yang terdiri dari backend API (Go) dan frontend web (Next.js) dengan integrasi pembayaran.

## 🚀 Fitur Utama

- **Pemilihan Tanggal**: Lihat tanggal yang tersedia untuk reservasi
- **Pemilihan Waktu**: Lihat slot waktu yang tersedia berdasarkan tanggal
- **Pemilihan Lapangan**: Lihat lapangan yang tersedia berdasarkan tanggal dan waktu
- **Pembuatan Reservasi**: Buat reservasi dengan integrasi pembayaran
- **Integrasi Pembayaran**: Mock payment gateway (Xendit)
- **UI Responsif**: Antarmuka web yang responsif dan modern

## 🛠️ Tech Stack

### Backend
- **Go 1.24**: Bahasa pemrograman
- **Gin**: Web framework
- **GORM**: ORM untuk operasi database
- **MySQL**: Database
- **golang-migrate**: Database migrations
- **Swagger**: API documentation

### Frontend
- **Next.js 16**: React framework
- **React 19**: Library UI
- **TypeScript**: Type-safe JavaScript
- **Tailwind CSS v4**: Utility-first CSS framework
- **Radix UI**: Component library

### Infrastructure
- **Docker & Docker Compose**: Containerization
- **MySQL 8.0**: Database server

## 📁 Struktur Project

```
diro/
├── diro-be/                    # Backend (Go)
│   ├── cmd/
│   │   ├── migrate/           # Database migration tool
│   │   └── seed/              # Database seeding tool
│   ├── internal/
│   │   ├── config/            # Configuration management
│   │   ├── database/          # Database connection
│   │   ├── handlers/          # HTTP request handlers
│   │   ├── models/            # Database models
│   │   ├── repositories/      # Data access layer
│   │   ├── routes/            # Route definitions
│   │   └── services/          # Business logic
│   ├── migrations/            # SQL migration files
│   ├── docs/                  # Swagger documentation
│   ├── .env.example           # Environment variables template
│   ├── Dockerfile             # Docker build file
│   ├── go.mod                 # Go modules
│   └── main.go                # Application entry point
├── diro-fe/                    # Frontend (Next.js)
│   ├── app/
│   │   ├── components/        # Reusable React components
│   │   ├── hooks/             # Custom React hooks
│   │   └── (pages)/           # Next.js pages
│   ├── lib/                   # Utility functions
│   ├── public/                # Static assets
│   ├── types/                 # TypeScript type definitions
│   ├── Dockerfile             # Docker build file
│   ├── package.json           # Node.js dependencies
│   └── tailwind.config.js     # Tailwind CSS config
├── docker-compose.yml         # Docker Compose configuration
├── .env                       # Environment variables
└── README.md                  # This file
```

## � Quick Start

```bash
# 1. Clone dan setup
git clone <repository-url>
cd diro
cp .env.example .env
# Edit .env dengan kredensial Anda

# 2. Jalankan aplikasi
docker-compose up --build

# 3. Setup database (jalankan di terminal baru)
docker-compose exec backend go run cmd/migrate/migrate.go -action=up
docker-compose exec backend go run cmd/seed/seed.go -action=seed

# 4. Akses aplikasi
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
# API Docs: http://localhost:8080/swagger/index.html
```

## �📋 Prerequisites

Sebelum menjalankan project ini, pastikan Anda memiliki:

- **Docker & Docker Compose** (versi terbaru)
- **Git** (untuk cloning repository)
- **Terminal/Command Line** yang mendukung shell commands

## 🚀 Cara Menjalankan

### Menggunakan Docker Compose

1. **Clone repository**
   ```bash
   git clone <repository-url>
   cd diro
   ```

2. **Setup environment variables**
   ```bash
   cp .env.example .env
   # Edit .env file dengan konfigurasi Anda
   ```

3. **Build dan jalankan containers**
   ```bash
   docker-compose up --build
   ```

4. **Akses aplikasi**
   - **Frontend**: http://localhost:3000
   - **Backend API**: http://localhost:8080
   - **API Documentation**: http://localhost:8080/swagger/index.html

## ⚙️ Konfigurasi Environment Variables

### File .env (Root Directory)
```env
# Database Configuration
MYSQL_ROOT_PASSWORD=your_secure_mysql_root_password
MYSQL_DATABASE=diro_db
MYSQL_USER=diro_user
MYSQL_PASSWORD=your_secure_mysql_password
MYSQL_PORT=3306

# Backend Configuration
BACKEND_PORT=8080
DB_HOST=mysql
DB_PORT=3306
DB_USER=diro_user
DB_PASSWORD=your_secure_mysql_password
DB_NAME=diro_db

# Frontend Configuration
FRONTEND_PORT=3000

# Payment Configuration (Xendit)
XENDIT_USERNAME=your_xendit_api_key
XENDIT_PASSWORD=your_xendit_secret_key
```

### Setup Xendit Payment Gateway

**1. Daftar Akun Xendit:**
- Kunjungi [https://dashboard.xendit.co](https://dashboard.xendit.co)
- Pilih "Sign Up" dan buat akun business/personal
- Verifikasi email dan nomor telepon

**2. Dapatkan API Keys:**
- Login ke Xendit Dashboard
- Pergi ke **Settings** → **API Keys**
- Klik **"Generate API Key"**
- Pilih environment: **Development** (untuk testing)
- Copy **API Key** dan **Secret Key**

**3. Konfigurasi Environment:**
```env
XENDIT_USERNAME=xnd_development_xxxxxxxxxxxxxxxxxxxxxxxxx
XENDIT_PASSWORD=xnd_secret_xxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

**4. Testing Payment:**
- Gunakan test card dari Xendit documentation
- Semua transaksi di development environment **GRATIS**
- Untuk production, ganti dengan production API keys

**📝 Catatan:** Jangan commit API keys ke repository. Selalu gunakan environment variables.

### Setup Database & Data Awal

Setelah containers berjalan dan environment sudah dikonfigurasi:

1. **Jalankan migrations** (membuat tabel database):
   ```bash
   docker-compose exec backend go run cmd/migrate/migrate.go -action=up
   ```

2. **Seed database** (isi data awal untuk testing):
   ```bash
   docker-compose exec backend go run cmd/seed/seed.go -action=seed
   ```

**Data yang akan dibuat:**
- ✅ 4 lapangan badminton (courts)
- ✅ 12 slot waktu (timeslots) 
- ✅ 5 sample reservasi dengan berbagai status

### File .env (diro-be Directory)
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password_here
DB_NAME=diro_db

# Payment Configuration
XENDIT_USERNAME=your_xendit_username
XENDIT_PASSWORD=your_xendit_password
```

## 📚 API Documentation

### Health Check
- `GET /health` - Cek kesehatan server

### Reservasi
- `GET /api/reservations/dates` - Dapatkan tanggal yang tersedia
- `GET /api/reservations/timeslots?date=2023-12-01` - Dapatkan slot waktu untuk tanggal tertentu
- `GET /api/reservations/courts?date=2023-12-01&timeslot_id=1` - Dapatkan lapangan yang tersedia
- `POST /api/reservations` - Buat reservasi baru
- `PUT /api/reservations/:id/confirm` - Konfirmasi reservasi
- `PUT /api/reservations/:id/cancel` - Batalkan reservasi

### User
- `GET /api/users/:id/reservations` - Dapatkan reservasi user

### Contoh Request/Response

#### Buat Reservasi
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

#### Response
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
    "updated_at": "2023-11-01T10:00:00Z"
  }
}
```

## 🗄️ Database Operations

### Migrations (Membuat/Mengubah Struktur Database)

Migrations digunakan untuk membuat dan mengubah struktur database secara terkontrol.

```bash
# Jalankan semua migration (membuat tabel database)
docker-compose exec backend go run cmd/migrate/migrate.go -action=up

# Rollback 1 migration terakhir
docker-compose exec backend go run cmd/migrate/migrate.go -action=down -steps=1

# Cek status migration saat ini
docker-compose exec backend go run cmd/migrate/migrate.go -action=version

# Buat migration baru
docker-compose exec backend go run cmd/migrate/migrate.go -action=create -name=nama_migration_baru
```

### Seeding (Mengisi Data Awal)

Seeding mengisi database dengan data awal untuk testing dan development.

```bash
# Isi database dengan data awal
docker-compose exec backend go run cmd/seed/seed.go -action=seed

# Hapus semua data yang di-seed
docker-compose exec backend go run cmd/seed/seed.go -action=clear
```

#### Data yang Akan Dibuat Saat Seeding:

**🏓 Courts (Lapangan):**
- Lapangan A: Lapangan badminton utama dengan pencahayaan LED
- Lapangan B: Lapangan badminton dengan lantai sintetis  
- Lapangan C: Lapangan badminton indoor dengan AC
- Lapangan D: Lapangan badminton outdoor (tidak aktif)

**⏰ Timeslots (Slot Waktu):**
- 08:00 - 22:00 (12 slot waktu)
- Slot 21:00-22:00 tidak aktif untuk testing

**📅 Sample Reservations:**
- 5 reservasi dengan berbagai status (confirmed, pending, cancelled)
- Menggunakan tanggal besok dan lusa untuk testing

### Mengakses Database

```bash
# Masuk ke MySQL container
docker-compose exec mysql mysql -u diro_user -p diro_db

# Password: sesuai yang di-set di .env (MYSQL_PASSWORD)
```

## 🧪 Testing

### Backend Testing
```bash
# Jalankan unit tests backend
docker-compose exec backend go test ./...
```

### Frontend Testing
```bash
# Jalankan linter frontend
docker-compose exec frontend npm run lint
```

## 🚀 Deployment

### Production dengan Docker

1. **Build images untuk production**
   ```bash
   docker-compose build
   ```

2. **Jalankan dalam mode detached**
   ```bash
   docker-compose up -d
   ```

3. **Cek logs**
   ```bash
   docker-compose logs -f
   ```

### Environment Variables untuk Production

Pastikan untuk mengatur environment variables yang sesuai untuk production:
- Database credentials yang aman
- Payment gateway credentials (Xendit production keys)
- Domain dan port yang sesuai

## 🐛 Troubleshooting

### Common Issues

1. **Port already in use**
   - Ubah port di `.env` file
   - Atau stop service yang menggunakan port tersebut

2. **Database connection failed**
   - Pastikan MySQL container sedang berjalan
   - Cek kredensial database di `.env`

3. **Migration failed**
   - Cek status migration: `docker-compose exec backend go run cmd/migrate/migrate.go -action=version`
   - Rollback jika diperlukan: `docker-compose exec backend go run cmd/migrate/migrate.go -action=down -steps=1`

4. **Frontend build failed**
   - Pastikan Node.js version >= 20.9.0
   - Clear cache: `docker-compose exec frontend rm -rf node_modules .next && npm install`

5. **Container fails to start**
   - Cek logs: `docker-compose logs <service_name>`
   - Pastikan semua environment variables sudah di-set dengan benar

### Logs & Monitoring

```bash
# Lihat logs semua container
docker-compose logs

# Lihat logs spesifik container
docker-compose logs backend
docker-compose logs frontend
docker-compose logs mysql

# Follow logs real-time
docker-compose logs -f

# Cek status container
docker-compose ps

# Monitor resource usage
docker stats
```

### Container Management

```bash
# Stop containers
docker-compose down

# Stop containers dan hapus volumes (data akan hilang!)
docker-compose down -v

# Rebuild spesifik service
docker-compose up --build backend

# Masuk ke container untuk debugging
docker-compose exec backend sh
docker-compose exec mysql mysql -u diro_user -p diro_db
```

## 🤝 Contributing

1. Fork repository
2. Buat feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📝 License

Project ini menggunakan lisensi MIT. Lihat file `LICENSE` untuk detail lebih lanjut.

## 📞 Support

Jika ada pertanyaan atau masalah, silakan buat issue di repository ini atau hubungi tim development.

---

**Happy coding! 🎾**