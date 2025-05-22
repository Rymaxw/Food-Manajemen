package operasi

import (
    "manajemen-stok-makanan/models"
    "fmt"
    "time"
)


func UbahBahanMakanan(ID int, nama string, jumlahStok int, tanggalKedaluwarsa string) {

    if ID < 1 || ID > models.JumlahData {
        fmt.Println("ID tidak valid.")
        return
    }

   
    indeksArray := ID - 1

  
    tgl, err := time.Parse("02-01-2006", tanggalKedaluwarsa)
    if err != nil {
        fmt.Println("Format tanggal salah. Gunakan format DD-MM-YYYY.")
        return
    }


    models.StokBahanMakanan[indeksArray] = models.BahanMakanan{
        Nama:           nama,
        JumlahStok:     jumlahStok,
        TanggalKedaluwarsa: tgl,
    }
    fmt.Println("Data bahan makanan berhasil diperbarui.")
}
