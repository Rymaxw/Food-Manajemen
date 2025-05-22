package operasi

import (
    "manajemen-stok-makanan/models"
    "fmt"
)

// HapusBahanMakanan menghapus bahan makanan berdasarkan ID
func HapusBahanMakanan(ID int) {
    // Validasi ID
    if ID < 1 || ID > models.JumlahData {
        fmt.Println("ID tidak valid.")
        return
    }


    indeksArray := ID - 1

    for i := indeksArray; i < models.JumlahData-1; i++ {
        models.StokBahanMakanan[i] = models.StokBahanMakanan[i+1]
    }
    models.JumlahData--
    fmt.Println("Bahan makanan berhasil dihapus.")
}
