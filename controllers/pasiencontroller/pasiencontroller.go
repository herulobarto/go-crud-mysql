package pasiencontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/herulobarto/go-crud-mysql/entities"
	"github.com/herulobarto/go-crud-mysql/libraries"
	"github.com/herulobarto/go-crud-mysql/models"
)

var validation = libraries.NewValidation()
var pasienModel = models.NewPasienModel()

func Index(response http.ResponseWriter, request *http.Request) {
	pasien, _ := pasienModel.FindAll()

	data := map[string]interface{}{
		"pasien": pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var pasien entities.Pasien
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors
		} else {
			err := pasienModel.Create(pasien)
			if err != nil {
				data["error"] = "Gagal menyimpan data pasien"
			} else {
				data["pesan"] = "Data pasien berhasil disimpan"
			}
		}

		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		id, err := strconv.ParseInt(request.URL.Query().Get("id"), 10, 64)
		if err != nil {
			panic(err)
		}

		pasien, err := pasienModel.FindByID(id)
		if err != nil {
			panic(err)
		}

		data := map[string]interface{}{
			"pasien": pasien,
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var pasien entities.Pasien
		pasien.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors
		} else {
			err := pasienModel.Update(pasien)
			if err != nil {
				data["error"] = "Gagal mengubah data pasien"
			} else {
				data["pesan"] = "Data pasien berhasil diubah"
			}
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	// Implementasi untuk delete pasien
}
