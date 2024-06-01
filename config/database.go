package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go-crud"

	// Sesuaikan hostname dan port MySQL sesuai dengan konfigurasi Anda
	dbHost := "localhost"
	dbPort := "3306"

	// Konstruksi DSN dengan format yang benar
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Membuka koneksi database menggunakan DSN yang sudah dibuat
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}

	// Verifikasi bahwa koneksi berhasil sebelum mengembalikan db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
