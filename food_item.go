package main

import (
	"fmt"
	"time"
)

// BahanMakanan merepresentasikan satu item bahan makanan
type BahanMakanan struct {
	Nama          string
	Kategori      string
	Jumlah        int
	Kedaluwarsa   time.Time
	TelahDipakai  bool
}

// BuatBahanMakanan membuat data bahan makanan baru
func BuatBahanMakanan(nama, kategori string, jumlah int, tanggalKedaluwarsa string) (BahanMakanan, error) {
	tanggal, err := time.Parse("2006-01-02", tanggalKedaluwarsa)
	if err != nil {
		return BahanMakanan{}, err
	}
	return BahanMakanan{
		Nama:         nama,
		Kategori:     kategori,
		Jumlah:       jumlah,
		Kedaluwarsa:  tanggal,
		TelahDipakai: false,
	}, nil
}

// Tampilkan menampilkan informasi bahan makanan
func (b BahanMakanan) Tampilkan() {
	status := "Tersedia"
	if b.TelahDipakai {
		status = "Telah Digunakan"
	}
	fmt.Printf("Nama: %s | Kategori: %s | Jumlah: %d | Kedaluwarsa: %s | Status: %s\n",
		b.Nama, b.Kategori, b.Jumlah, b.Kedaluwarsa.Format("DD-YY-MM"), status)
}

// HampirKedaluwarsa mengecek apakah bahan makanan akan kedaluwarsa dalam 3 hari
func (b BahanMakanan) HampirKedaluwarsa() bool {
	return time.Until(b.Kedaluwarsa).Hours() <= 72
}

// TandaiSebagaiDipakai menandai bahan makanan sebagai sudah digunakan
func (b *BahanMakanan) TandaiSebagaiDipakai() {
	b.TelahDipakai = true
}

// Fungsi utama untuk pengujian
func main() {
	bahan, err := BuatBahanMakanan("Gula", "Bahan Makanan", 3, "2027-03-10")
	if err != nil {
		fmt.Println("Gagal membuat bahan makanan:", err)
		return
	}

	bahan.Tampilkan()

	if bahan.HampirKedaluwarsa() {
		fmt.Println("PERINGATAN: Bahan ini hampir kedaluwarsa!")
	}

	bahan.TandaiSebagaiDipakai()
	fmt.Println("\nSetelah digunakan:")
	bahan.Tampilkan()
}
