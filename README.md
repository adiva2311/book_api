
# 📚 Book API

Book API adalah layanan RESTful yang memungkinkan pengguna untuk mengelola data buku, termasuk operasi CRUD (Create, Read, Update, Delete).

---

## ✨ Fitur

- ✅ **Menambahkan Buku**  
  Tambahkan buku baru dengan informasi seperti judul, penulis, dan tahun terbit.
- 📖 **Melihat Daftar Buku**  
  Dapatkan daftar semua buku yang tersedia dalam database.
- ✏️ **Memperbarui Informasi Buku**  
  Perbarui detail buku yang sudah ada.
- 🗑️ **Menghapus Buku**  
  Hapus buku dari database.

---

## 🗂️ Struktur Proyek

```
book_api/
├── app/
│   ├── controller/      # Handler untuk setiap endpoint API
│   ├── helper/          # Fungsi-fungsi utilitas
│   ├── model/           # Definisi struktur data
│   ├── repository/      # Lapisan akses data
│   └── services/        # Logika bisnis aplikasi
├── main.go              # Entry point aplikasi
└── go.mod               # File konfigurasi Go module
```

---

## 🚀 Instalasi

### 1. Kloning Repositori

```bash
git clone https://github.com/adiva2311/book_api.git
```

### 2. Masuk ke Direktori Proyek

```bash
cd book_api
```

### 3. Instalasi Dependensi

Pastikan Anda memiliki [Go](https://golang.org/doc/install) terinstal di sistem Anda.

```bash
go mod tidy
```

---

## 🧪 Penggunaan

### 1. Menjalankan Aplikasi

```bash
go run main.go
```

### 2. Endpoint API

| Method | Endpoint        | Deskripsi                              |
|--------|------------------|----------------------------------------|
| POST   | `/books`         | Tambahkan buku baru                    |
| GET    | `/books`         | Dapatkan daftar semua buku            |
| GET    | `/books/{id}`    | Dapatkan detail buku berdasarkan ID   |
| PUT    | `/books/{id}`    | Perbarui informasi buku berdasarkan ID|
| DELETE | `/books/{id}`    | Hapus buku berdasarkan ID             |

---

## 🤝 Kontribusi

Kontribusi sangat terbuka! Silakan fork proyek ini, buat cabang baru, dan ajukan pull request untuk perubahan yang Anda buat.

---

## 📄 Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
