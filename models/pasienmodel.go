package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/herulobarto/go-crud-mysql/config"
	"github.com/herulobarto/go-crud-mysql/entities"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel) FindAll() ([]entities.Pasien, error) {

	rows, err := p.conn.Query("SELECT * FROM pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPasien []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(&pasien.Id,
			&pasien.NamaLengkap,
			&pasien.NIK,
			&pasien.JenisKelamin,
			&pasien.TempatLahir,
			&pasien.TanggalLahir,
			&pasien.Alamat,
			&pasien.NoHp)

		if pasien.JenisKelamin == "1" {
			pasien.JenisKelamin = "Laki-laki"
		} else {
			pasien.JenisKelamin = "Perempuan"
		}

		// 2006-01-02 == yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", pasien.TanggalLahir)
		// 02-01-2006 == dd-mm-yyyy
		pasien.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataPasien = append(dataPasien, pasien)
	}

	return dataPasien, nil

}

func (p *PasienModel) Create(pasien entities.Pasien) bool {
	result, err := p.conn.Exec("INSERT INTO pasien (nama_lengkap, NIK, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) VALUES (?, ?, ?, ?, ?, ?, ?)",
		pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
