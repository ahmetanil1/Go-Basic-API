# Library Management Project


--TR--
Bu proje, temel bir **kütüphane yönetim sistemi** örneğidir.  
Amaç, kitap ekleme, listeleme, güncelleme ve silme işlemlerini MongoDB ile birlikte **HTTP üzerinden manuel olarak** gerçekleştirmektir.

> ⚠️ Not: Hazır framework veya fonksiyonlardan kaçınıp, temel mantığı kendim anlamak ve uygulamak için her şeyi kendi başıma yazdım.

## Özellikler

- Kitap ekleme (POST /books)
- Tüm kitapları listeleme (GET /books)
- ID ile kitap bilgisi alma (GET /books/{id})
- Kitap güncelleme (PUT /books/{id})
- Kitap silme (DELETE /books/{id})
- MongoDB ile veri depolama
- Basit logging middleware

/internal
/db         → MongoDB bağlantısı
/handlers   → HTTP handler’lar
/models     → Book modeli
/repos      → MongoDB repository
/services   → Business logic katmanı
main.go      → Sunucu başlangıcı ve routing

--EN--
This project is a simple **library management system** example.  
The goal is to manually implement book creation, listing, updating, and deletion with MongoDB over **HTTP requests**.

> ⚠️ Note: I avoided using ready-made frameworks or functions; I wrote everything myself to understand the fundamentals clearly.

## Features

- Add a new book (POST /books)
- List all books (GET /books)
- Get book by ID (GET /books/{id})
- Update book (PUT /books/{id})
- Delete book (DELETE /books/{id})
- Store data in MongoDB
- Simple logging middleware

/internal
/db         → MongoDB connection
/handlers   → HTTP handlers
/models     → Book model
/repos      → MongoDB repository
/services   → Business logic layer
main.go      → Server startup and routing
