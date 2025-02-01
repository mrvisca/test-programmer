package models

import "github.com/jinzhu/gorm"

type Status struct {
	gorm.Model
	Nama string `gorm:"type:varchar(200)"`
}

type ResStatus struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
}
