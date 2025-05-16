package operasi

import (
    "manajemen-stok-makanan/models"
    "fmt"
)

// HapusBahanMakanan menghapus bahan makanan berdasarkan indeks
func HapusBahanMakanan(indeks int) {
    if indeks < 0 || indeks >= models.JumlahData {
        fmt.Println("Indeks tidak valid.")
        return
    }

    // Geser elemen-elemen setelah indeks ke kiri
    for i := indeks; i < models.JumlahData-1; i++ {
        models.StokBahanMakanan[i] = models.StokBahanMakanan[i+1]
    }
    models.JumlahData--
    fmt.Println("Bahan makanan berhasil dihapus.")
}