## Cashier App with GORM

### Description

Pengembangan dari **Cashier App** sebelumnya yaitu, melakukan perubahan metode penyimpanan data menggunakan database **PostgreSQL**. Tentunya akan berinteraksi dengan database menggunakan **GORM**.

Terdapat aplikasi yang bisa dijalankan di project ini dengan perintah `go run main.go` yang akan menjalankan aplikasi web di port `8080`. Disini sudah disediakan API dan endpoint yang bisa dibuka di browser untuk UI dari aplikasi ini.

### Penting untuk mengubah koneksi database lokal menjadi milik anda :

```go
dbCredentials = Credential{
    Host:         "localhost",
    Username:     "postgres", // <- ubah ini
    Password:     "postgres", // <- ubah ini
    DatabaseName: "database", // <- ubah ini
    Port:         5432,
}
```