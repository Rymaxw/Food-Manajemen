package operasi

import (
    "manajemen-stok-makanan/models"
    "fmt"
    "time"
)

// UbahBahanMakanan mengubah data bahan makanan berdasarkan indeks
func UbahBahanMakanan(indeks int, nama string, jumlahStok int, tanggalKedaluwarsa string) {
    if indeks < 0 || indeks >= models.JumlahData {
        fmt.Println("Indeks tidak valid.")
        return
    }

    tgl, err := time.Parse("02-01-2006", tanggalKedaluwarsa)
    if err != nil {
        fmt.Println("Format tanggal salah. Gunakan format DD-MM-YYYY.")
        return
    }

    models.StokBahanMakanan[indeks] = models.BahanMakanan{
        Nama:           nama,
        JumlahStok:     jumlahStok,
        TanggalKedaluwarsa: tgl,
    }
    fmt.Println("Data bahan makanan berhasil diperbarui.")
}