package utils

import (
    "manajemen-stok-makanan/models"
)

// SequentialSearch mencari bahan makanan berdasarkan nama menggunakan Sequential Search
func SequentialSearch(nama string) int {
    for i := 0; i < models.JumlahData; i++ {
        if models.StokBahanMakanan[i].Nama == nama {
            return i
        }
    }
    return -1
}

// BinarySearch mencari bahan makanan berdasarkan nama menggunakan Binary Search
func BinarySearch(nama string) int {
    awal := 0
    akhir := models.JumlahData - 1

    for awal <= akhir {
        tengah := (awal + akhir) / 2
        if models.StokBahanMakanan[tengah].Nama == nama {
            return tengah
        } else if models.StokBahanMakanan[tengah].Nama < nama {
            awal = tengah + 1
        } else {
            akhir = tengah - 1
        }
    }
    return -1
}