package pasiencontroller

import (
	"log"
	"net/http"
	"text/template"

	"github.com/raihanfaiq72/go-crud/entities"
	"github.com/raihanfaiq72/go-crud/models"
)

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

// tambah data
func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			// panic(err)
			log.Println("Error:", err)
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
		pasien.NoHP = request.Form.Get("no_hp")

		pasienModel.Create(pasien)
		data := map[string]interface{}{
			"pesan": "data pasien berhasil ditambahkan",
		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}
}

// edit data
func Edit(response http.ResponseWriter, request *http.Request) {

}

// delete
func Delete(response http.ResponseWriter, request *http.Request) {

}
