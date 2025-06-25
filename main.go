package main

import (
	"fmt"
	"linklist/model"
	"linklist/node"
	"net/http"
	"strconv"
	"text/template"
)

func seeder() {
	jab1 := node.Jabatan{
		IdJabatan:   1,
		NamaJabatan: "Manager",
		Gaji:        10000000,
	}

	jab2 := node.Jabatan{
		IdJabatan:   2,
		NamaJabatan: "Supervisor",
		Gaji:        8000000,
	}

	jab3 := node.Jabatan{
		IdJabatan:   3,
		NamaJabatan: "Staff",
		Gaji:        6000000,
	}

	jab4 := node.Jabatan{
		IdJabatan:   4,
		NamaJabatan: "Intern",
		Gaji:        3000000,
	}

	model.CreateJabatan(jab1)
	model.CreateJabatan(jab2)
	model.CreateJabatan(jab3)
	model.CreateJabatan(jab4)

	pegawai1 := node.Pegawai{
		ID:     1,
		Nama:   "kurniawan",
		NoTelp: "031",
		Alamat: node.Address{
			Kota:  "surabaya",
			Nomer: 10,
			Jalan: "sidosermo",
		},
		Jabatan: 1,
		Email:   "kurniawan@gmail.com",
	}
	model.CreatePegawai(pegawai1)

	pegawai2 := node.Pegawai{
		ID:     2,
		Nama:   "aan",
		NoTelp: "031",
		Alamat: node.Address{
			Kota:  "surabaya",
			Nomer: 10,
			Jalan: "sidoresmo",
		},
		Jabatan: 1,
		Email:   "aan@gmail.com",
	}
	model.CreatePegawai(pegawai2)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	dataPegawai := model.ReadPegawai()
	fmt.Println(dataPegawai)
	tmpl.Execute(w, dataPegawai)
}

func InsertForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func InsertProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		id, _ := strconv.Atoi(r.FormValue("id"))
		pegawai := node.Pegawai{
			ID:     id,
			Nama:   r.FormValue("nama"),
			NoTelp: r.FormValue("notelp"),
			Email:  r.FormValue("email"),
			Alamat: node.Address{
				Kota:  r.FormValue("kota"),
				Jalan: r.FormValue("jalan"),
			},
		}

		// Parsing angka
		fmt.Sscanf(r.FormValue("nomer"), "%d", &pegawai.Alamat.Nomer)
		fmt.Sscanf(r.FormValue("jabatan"), "%d", &pegawai.Jabatan)

		model.CreatePegawai(pegawai)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditForm(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	pegawai, ok := model.GetPegawaiById(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	tmpl, _ := template.ParseFiles("templates/edit.html")
	tmpl.Execute(w, pegawai)
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()

	// parsing numerik
	id, _ := strconv.Atoi(r.FormValue("id"))
	nomer, _ := strconv.Atoi(r.FormValue("nomer"))
	jabatan, _ := strconv.Atoi(r.FormValue("jabatan"))

	updated := node.Pegawai{
		ID:      id,
		Nama:    r.FormValue("nama"),
		NoTelp:  r.FormValue("notelp"),
		Email:   r.FormValue("email"),
		Jabatan: jabatan,
		Alamat: node.Address{
			Kota:  r.FormValue("kota"),
			Jalan: r.FormValue("jalan"),
			Nomer: nomer,
		},
	}

	model.UpdatePegawai(updated, id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func DeletePegawai(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	model.DeletePegawai(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func main() {
	seeder()
	http.HandleFunc("/", Index)
	http.HandleFunc("/tambah", InsertForm)
	http.HandleFunc("/insert", InsertProcess)
	http.HandleFunc("/edit", EditForm)
	http.HandleFunc("/update", UpdateProcess)
	http.HandleFunc("/delete", DeletePegawai)

	http.ListenAndServe(":8080", nil)
}
