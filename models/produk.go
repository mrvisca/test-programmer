package models

import "github.com/jinzhu/gorm"

type Produk struct {
	gorm.Model
	KategoriId uint     `gorm:"default:0"`
	StatusId   uint     `gorm:"default:0"`
	Kategori   Kategori `gorm:"foreignKey:KategoriId"`
	Setatus    Status   `gorm:"foreignKey:StatusId"`
	Nama       string   `gorm:"type:varchar(200)"`
	Harga      int64    `gorm:"default:0"`
}

type ResProduk struct {
	ID           uint   `json:"id"`
	KategoriId   uint   `json:"kategori_id"`
	KategoriNama string `json:"kategori_nama"`
	StatusId     uint   `json:"status_id"`
	StatusNama   string `json:"status_nama"`
	Nama         string `json:"nama"`
	Harga        int64  `json:"harga"`
}
