package models

import (
    "time"
)

// BahanMakanan merepresentasikan satu item bahan makanan
type BahanMakanan struct {
    Nama           string
    JumlahStok     int
    TanggalKedaluwarsa time.Time
}

// Array statis dengan dimensi 100
var StokBahanMakanan [100]BahanMakanan
var JumlahData int = 0 // Untuk melacak jumlah data yang ada