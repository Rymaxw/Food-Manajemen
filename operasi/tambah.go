package operasi

import (
    "manajemen-stok-makanan/models"
    "fmt"
    "time"
)

// TambahBahanMakanan menambahkan bahan makanan baru ke array
func TambahBahanMakanan(nama string, jumlahStok int, tanggalKedaluwarsa string) {
    if models.JumlahData >= 100 {
        fmt.Println("Maaf, kapasitas maksimum sudah tercapai.")
        return
    }

    tgl, err := time.Parse("02-01-2006", tanggalKedaluwarsa)
    if err != nil {
        fmt.Println("Format tanggal salah. Gunakan format DD-MM-YYYY.")
        return
    }

    models.StokBahanMakanan[models.JumlahData] = models.BahanMakanan{
        Nama:           nama,
        JumlahStok:     jumlahStok,
        TanggalKedaluwarsa: tgl,
    }
    models.JumlahData++
    fmt.Println("Bahan makanan berhasil ditambahkan.")
}