package utils

import (
    "manajemen-stok-makanan/models"
    "fmt"
)

// SelectionSort mengurutkan bahan makanan berdasarkan tanggal kedaluwarsa
func SelectionSort() {
    for i := 0; i < models.JumlahData-1; i++ {
        minIndex := i
        for j := i + 1; j < models.JumlahData; j++ {
            if models.StokBahanMakanan[j].TanggalKedaluwarsa.Before(models.StokBahanMakanan[minIndex].TanggalKedaluwarsa) {
                minIndex = j
            }
        }
        // Tukar elemen
        models.StokBahanMakanan[i], models.StokBahanMakanan[minIndex] = models.StokBahanMakanan[minIndex], models.StokBahanMakanan[i]
    }
    fmt.Println("Daftar bahan makanan telah diurutkan berdasarkan tanggal kedaluwarsa.")
}

// InsertionSort mengurutkan bahan makanan berdasarkan jumlah stok
func InsertionSort() {
    for i := 1; i < models.JumlahData; i++ {
        current := models.StokBahanMakanan[i]
        j := i - 1
        for j >= 0 && models.StokBahanMakanan[j].JumlahStok > current.JumlahStok {
            models.StokBahanMakanan[j+1] = models.StokBahanMakanan[j]
            j--
        }
        models.StokBahanMakanan[j+1] = current
    }
    fmt.Println("Daftar bahan makanan telah diurutkan berdasarkan jumlah stok.")
}