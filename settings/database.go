package settings

import (
	"test-programmer/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:@(localhost)/progtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Koneksi ke database gagal!")
	}

	DB.AutoMigrate(&models.Kategori{})
	DB.AutoMigrate(&models.Status{})
	DB.AutoMigrate(&models.Produk{}).AddForeignKey("kategori_id", "kategoris(id)", "NO ACTION", "NO ACTION").AddForeignKey("status_id", "statuses(id)", "NO ACTION", "NO ACTION")
}
