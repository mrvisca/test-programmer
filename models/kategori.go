package models

import "github.com/jinzhu/gorm"

type Kategori struct {
	gorm.Model
	Nama string `gorm:"type:varchar(200)"`
}

type ResKategori struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
}
