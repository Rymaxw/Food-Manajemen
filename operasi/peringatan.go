package operasi

import (
    "manajemen-stok-makanan/models"
    "fmt"
    "time"
)

// PeriksaKedaluwarsa memeriksa bahan makanan yang mendekati tanggal kedaluwarsa
func PeriksaKedaluwarsa() {
    saatIni := time.Now()
    for i := 0; i < models.JumlahData; i++ {
        hariKedaluwarsa := models.StokBahanMakanan[i].TanggalKedaluwarsa.Sub(saatIni).Hours() / 24
        if hariKedaluwarsa < 7 { // Jika kurang dari 7 hari
            fmt.Printf("PERINGATAN: %s akan kedaluwarsa dalam %.0f hari.\n", models.StokBahanMakanan[i].Nama, hariKedaluwarsa)
        }
    }
}