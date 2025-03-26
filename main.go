package main

import (
	"bufio"
	"fmt"
	"linklist/view"
	"os"
	"strconv"
)

func menuUtama() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Data Pegawai")
		fmt.Println("2. Tampilkan Data Pegawai")
		fmt.Println("3. Update Data Pegawai")
		fmt.Println("4. Hapus Data Pegawai")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		// Membaca input dari user
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Menghapus karakter newline (`\n`)

		// Konversi input menjadi integer
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, silakan masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Anda memilih: Tambah Data Pegawai")
			view.Insert()
		case 2:
			fmt.Println("Anda memilih: Tampilkan Data Pegawai")
			view.Views()
		case 3:
			fmt.Println("Anda memilih: Update Data Pegawai")
			view.Update()
		case 4:
			fmt.Println("Anda memilih: Hapus Data Pegawai")
			view.Delete()
		case 5:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func main() {
	menuUtama()
}
