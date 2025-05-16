package main

import (
    "fmt"
    "manajemen-stok-makanan/operasi"
    "manajemen-stok-makanan/utils"
)

func main() {
    // Muat data dari file JSON saat aplikasi dimulai
    utils.MuatDariJSON()

    for {
        fmt.Println("\nMenu Aplikasi Manajemen Stok Bahan Makanan:")
        fmt.Println("1. Tambah Bahan Makanan")
        fmt.Println("2. Ubah Bahan Makanan")
        fmt.Println("3. Hapus Bahan Makanan")
        fmt.Println("4. Cari Bahan Makanan")
        fmt.Println("5. Urutkan Bahan Makanan")
        fmt.Println("6. Tampilkan Laporan")
        fmt.Println("7. Periksa Kedaluwarsa")
        fmt.Println("8. Keluar")

        var pilihan int
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            var nama string
            var jumlahStok int
            var tanggalKedaluwarsa string
            fmt.Print("Nama bahan makanan: ")
            fmt.Scan(&nama)
            fmt.Print("Jumlah stok: ")
            fmt.Scan(&jumlahStok)
            fmt.Print("Tanggal kedaluwarsa (DD-MM-YYYY): ")
            fmt.Scan(&tanggalKedaluwarsa)
            operasi.TambahBahanMakanan(nama, jumlahStok, tanggalKedaluwarsa)
        case 2:
            var indeks int
            var nama string
            var jumlahStok int
            var tanggalKedaluwarsa string
            fmt.Print("Indeks bahan makanan: ")
            fmt.Scan(&indeks)
            fmt.Print("Nama baru: ")
            fmt.Scan(&nama)
            fmt.Print("Jumlah stok baru: ")
            fmt.Scan(&jumlahStok)
            fmt.Print("Tanggal kedaluwarsa baru (DD-MM-YYYY): ")
            fmt.Scan(&tanggalKedaluwarsa)
            operasi.UbahBahanMakanan(indeks, nama, jumlahStok, tanggalKedaluwarsa)
        case 3:
            var indeks int
            fmt.Print("Indeks bahan makanan yang akan dihapus: ")
            fmt.Scan(&indeks)
            operasi.HapusBahanMakanan(indeks)
        case 4:
            var nama string
            fmt.Print("Cari bahan makanan berdasarkan nama: ")
            fmt.Scan(&nama)
            indeks := utils.SequentialSearch(nama)
            if indeks != -1 {
                fmt.Printf("Bahan makanan ditemukan di indeks %d.\n", indeks)
            } else {
                fmt.Println("Bahan makanan tidak ditemukan.")
            }
        case 5:
            fmt.Println("1. Urutkan berdasarkan tanggal kedaluwarsa")
            fmt.Println("2. Urutkan berdasarkan jumlah stok")
            var pilihanUrut int
            fmt.Print("Pilih metode pengurutan: ")
            fmt.Scan(&pilihanUrut)
            switch pilihanUrut {
            case 1:
                utils.SelectionSort()
            case 2:
                utils.InsertionSort()
            default:
                fmt.Println("Pilihan tidak valid.")
            }
        case 6:
            operasi.TampilkanLaporan()
        case 7:
            operasi.PeriksaKedaluwarsa()
        case 8:
            fmt.Println("Simpan data sebelum keluar...")
            utils.SimpanKeJSON()
            fmt.Println("Terima kasih telah menggunakan aplikasi!")
            return
        default:
            fmt.Println("Pilihan tidak valid.")
        }
    }
}