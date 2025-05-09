package main

import (
	"fmt"
	"time"
)

// Struct bahan makanan (bisa dipindah ke file lain kalau ingin lebih modular)
type BahanMakanan struct {
	Nama        string
	Jumlah      int
	Kedaluwarsa time.Time
}

// Fungsi untuk mengubah data bahan makanan
func UbahBahanMakanan(data []BahanMakanan, nama string, jumlahBaru int, kedaluwarsaBaru time.Time) ([]BahanMakanan, bool) {
	for i := 0; i < len(data); i++ {
		if data[i].Nama == nama {
			data[i].Jumlah = jumlahBaru
			data[i].Kedaluwarsa = kedaluwarsaBaru
			fmt.Printf("Data bahan makanan '%s' berhasil diubah.\n", nama)
			return data, true
		}
	}
	fmt.Printf("Bahan makanan '%s' tidak ditemukan.\n", nama)
	return data, false
}
