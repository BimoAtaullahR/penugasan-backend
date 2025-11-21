# Penugasan Backend OmahTI 2025

Ini adalah submission untuk penugasan Backend Division OmahTI 2025. Proyek ini adalah RESTful API yang dibangun menggunakan **Golang**, **Gin**, **GORM**, dan **PostgreSQL**.

## 🌐 Live Demo (Deployment)

API ini sudah di-deploy dan dapat diakses secara publik menggunakan **Hugging Face Spaces (Docker)**.

* **Base URL:** `https://bimoar-penugasan-backend.hf.space`

## 📚 Cara Testing (Postman)

File koleksi Postman (`Penugasan Backend OmahTI.postman_collection.json`) telah disertakan dalam repository ini.

**Langkah Testing:**
1.  Download file `.json` postman dari repo ini.
2.  Import ke Postman.
3.  Gunakan URL Deployment di atas sebagai environment, atau ganti `{{url}}` di Postman dengan link Hugging Face tersebut.

### Daftar Endpoint Tersedia

| Method | Endpoint | Auth | Deskripsi |
| :--- | :--- | :--- | :--- |
| `POST` | `/register` | Public | Mendaftarkan user baru (Email & Password). |
| `POST` | `/login` | Public | Login & mendapatkan **JWT Token**. |
| `GET` | `/api/me` | **Bearer Token** | Endpoint rahasia. Menampilkan profil user dari token. |

## 🛠️ Tech Stack
* **Language:** Go (Golang) 1.25.4
* **Framework:** Gin Gonic
* **ORM:** GORM
* **Database:** PostgreSQL (Supabase via Connection Pooler)
* **Deployment:** Hugging Face Spaces (Docker Container)

## 📝 Catatan Tambahan
Fitur autentikasi menggunakan JWT (JSON Web Token) dengan algoritma HS256. Password di-hash menggunakan bcrypt sebelum disimpan ke database.