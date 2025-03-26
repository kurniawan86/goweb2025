package view

import (
	"bufio"
	"fmt"
	"linklist/model"
	"linklist/node"
	"os"
)

func Insert() {
	var kota, nama, notelp, email string
	var id, nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pegawai: ")
	fmt.Scan(&id)

	fmt.Print("Masukkan Nama Pegawai: ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukkan Jalan: ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukkan Nomer rumah: ")
	fmt.Scan(&nomer)

	fmt.Print("Masukkan Kota: ")
	fmt.Scan(&kota)

	fmt.Print("Masukkan Nomor Telepon: ")
	fmt.Scan(&notelp)

	fmt.Print("Masukkan Email: ")
	fmt.Scan(&email)

	// create new pegawai
	pegawai := node.Pegawai{
		ID:     id,
		Nama:   nama,
		Alamat: node.Address{jalan, kota, nomer},
		NoTelp: notelp,
		Email:  email,
	}

	// insert to DaftarPegawai
	cek := model.CreatePegawai(pegawai)
	if cek {
		fmt.Println("== Pegawai berhasil ditambahkan ==")
	} else {
		fmt.Println("Pegawai gagal ditambahkan")
	}
	fmt.Println()
}

func Views() {
	fmt.Println("Daftar Pegawai")
	for i, emp := range model.ReadPegawai() {
		fmt.Println("--- Pegawai ke -", i+1, " ---")
		fmt.Println("ID Pegawai\t: ", emp.ID)
		fmt.Println("Nama Pegawai\t: ", emp.Nama)
		fmt.Println("Alamat\t\t: ", emp.Alamat.Jalan, emp.Alamat.Nomer, emp.Alamat.Kota)
		fmt.Println("No Telepon\t: ", emp.NoTelp)
		fmt.Println("Email\t\t: ", emp.Email)
		fmt.Println()
	}
}

func Update() {
	var id, nomer int
	var nama, kota, notelp, email string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pegawai yang akan diupdate: ")
	fmt.Scan(&id)

	fmt.Print("Masukkan Nama Pegawai: ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukkan Kota: ")
	fmt.Scan(&kota)

	fmt.Print("Masukkan Jalan: ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukkan nomer rumah: ")
	fmt.Scan(&nomer)

	fmt.Print("Masukkan Nomor Telepon: ")
	fmt.Scan(&notelp)

	fmt.Print("Masukkan Email: ")
	fmt.Scan(&email)

	// create new pegawai
	pegawai := node.Pegawai{
		ID:     id,
		Nama:   nama,
		Alamat: node.Address{jalan, kota, nomer},
		NoTelp: notelp,
		Email:  email,
	}

	// update pegawai
	cek := model.UpdatePegawai(pegawai, id)
	if cek {
		fmt.Println("== Pegawai berhasil diupdate ==")
	} else {
		fmt.Println("Pegawai gagal diupdate")
	}
	fmt.Println()
}

func Delete() {
	var id int
	fmt.Print("Masukkan ID Pegawai yang akan dihapus: ")
	fmt.Scan(&id)

	cek := model.DeletePegawai(id)
	if cek {
		fmt.Println("== Pegawai berhasil dihapus ==")
	} else {
		fmt.Println("Pegawai gagal dihapus")
	}
	fmt.Println()
}
