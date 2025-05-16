package utils

import (
    "encoding/json"
    "os"
    "manajemen-stok-makanan/models"
)

// SimpanKeJSON menyimpan data ke file JSON
func SimpanKeJSON() error {
    file, err := os.Create("stok_bahan_makanan.json")
    if err != nil {
        return err
    }
    defer file.Close()

    data := make([]models.BahanMakanan, models.JumlahData)
    copy(data, models.StokBahanMakanan[:models.JumlahData])

    encoder := json.NewEncoder(file)
    return encoder.Encode(data)
}

// MuatDariJSON memuat data dari file JSON
func MuatDariJSON() error {
    file, err := os.Open("stok_bahan_makanan.json")
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    var data []models.BahanMakanan
    err = decoder.Decode(&data)
    if err != nil {
        return err
    }

    for i := range data {
        models.StokBahanMakanan[i] = data[i]
    }
    models.JumlahData = len(data)

    return nil
}