# Quiz API

Quiz API adalah RESTful API berbasis Golang + Gin untuk mengelola sistem kuis dengan kategori, pertanyaan, hasil kuis, dan laporan.  
Menggunakan PostgreSQL sebagai database, dengan dukungan JWT Authentication dan Role-based Access (User & Admin).  
API ini sudah live di Railway.

Base URL Production:  
https://quiz-api-v1-production.up.railway.app/api/v1

### Daftar Endpoint
## Users & Authentication  
POST   `/users/register` → Registrasi user baru  
POST   `/users/login` → Login & mendapatkan JWT token  
GET    `/users/profile` → Ambil profil user yang sedang login  

### Categories  
GET    `/categories` → Ambil semua kategori (User & Admin)  
POST   `/categories` → Tambah kategori baru (Admin)  
PUT    `/categories/:id` → Update kategori berdasarkan ID (Admin)  
DELETE `/categories/:id` → Hapus kategori berdasarkan ID (Admin)  

### Questions  
GET    `/questions` → Ambil semua pertanyaan (Admin)  
GET    `/questions/:id` → Ambil pertanyaan berdasarkan ID (Admin)  
POST   `/questions` → Tambah pertanyaan baru (Admin)  
PUT    `/questions/:id` → Update pertanyaan berdasarkan ID (Admin)  
DELETE `/questions/:id` → Hapus pertanyaan berdasarkan ID (Admin)  

### Quiz  
GET    `/quiz/start?category={id}` → Mulai kuis berdasarkan kategori  
POST   `/quiz/submit` → Submit jawaban kuis  
GET    `/quiz/results` → Ambil hasil kuis user yang sedang login  

### Reports  
GET    `/reports/top-scores` → Ambil laporan skor tertinggi (Admin)  
GET    `/reports/users` → Ambil laporan aktivitas user (Admin)  

Authentication  
Gunakan JWT Token untuk mengakses endpoint yang memerlukan autentikasi.  
Tambahkan header:  
Authorization: Bearer <your_token_here>
