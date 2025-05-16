package operasi

import (
    "manajemen-stok-makanan/models"
    "fmt"
    "time"
)

// TampilkanLaporan menampilkan laporan total bahan makanan
func TampilkanLaporan() {
    fmt.Println("\nLaporan Stok Bahan Makanan:")
    fmt.Printf("%-5s %-20s %-10s %-15s %-10s\n", "ID", "Nama", "Jumlah", "Kedaluwarsa", "Status")
    fmt.Println("-------------------------------------------------------------")

    var totalStok int = 0
    for i := 0; i < models.JumlahData; i++ {
        tanggalFormatted := models.StokBahanMakanan[i].TanggalKedaluwarsa.Format("02-01-2006")
        status := getStatusKedaluwarsa(models.StokBahanMakanan[i].TanggalKedaluwarsa)
        fmt.Printf("%-5d %-20s %-10d %-15s %-10s\n", i+1, models.StokBahanMakanan[i].Nama, models.StokBahanMakanan[i].JumlahStok, tanggalFormatted, status)
        totalStok += models.StokBahanMakanan[i].JumlahStok
    }

    fmt.Println("-------------------------------------------------------------")
    fmt.Printf("Total bahan makanan yang tersedia: %d\n", models.JumlahData)
    fmt.Printf("Total jumlah stok: %d\n", totalStok)
}

// getStatusKedaluwarsa menentukan status kedaluwarsa bahan makanan
func getStatusKedaluwarsa(tanggalKedaluwarsa time.Time) string {
    saatIni := time.Now()
    hariKedaluwarsa := tanggalKedaluwarsa.Sub(saatIni).Hours() / 24

    if hariKedaluwarsa <= 0 {
        return "Kedaluwarsa"
    } else if hariKedaluwarsa <= 7 {
        return "Segera"
    } else {
        return "Aman"
    }
}