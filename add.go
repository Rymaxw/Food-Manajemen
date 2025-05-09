package main

import (
	"fmt"
)

// TambahBahanMakanan menambahkan data bahan makanan ke dalam daftar
func TambahBahanMakanan(daftar *[]BahanMakanan, nama, kategori string, jumlah int, tanggalKedaluwarsa string) {
	bahan, err := BuatBahanMakanan(nama, kategori, jumlah, tanggalKedaluwarsa)
	if err != nil {
		fmt.Println("Gagal menambahkan bahan makanan:", err)
		return
	}
	*daftar = append(*daftar, bahan)
	fmt.Println("Bahan makanan berhasil ditambahkan.")
}