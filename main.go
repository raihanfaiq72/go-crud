package main

import (
	"net/http"

	"github.com/raihanfaiq72/go-crud/controllers/pasiencontroller"
)

func main() {
	http.HandleFunc("/", pasiencontroller.Index)
	http.HandleFunc("/tambah-pasien", pasiencontroller.Add)
	http.HandleFunc("/edit-pasien", pasiencontroller.Edit)
	http.HandleFunc("/hapus-pasien", pasiencontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
