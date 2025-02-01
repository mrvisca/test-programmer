package controllers

import (
	"test-programmer/helpers"
	"test-programmer/models"
	"test-programmer/settings"

	"github.com/gin-gonic/gin"
)

func FillResProduct(pro models.Produk) models.ResProduk {
	return models.ResProduk{
		ID:           pro.ID,
		KategoriId:   pro.KategoriId,
		KategoriNama: pro.Kategori.Nama,
		StatusId:     pro.StatusId,
		StatusNama:   pro.Setatus.Nama,
		Nama:         pro.Nama,
		Harga:        pro.Harga,
	}
}

func ListProduk(c *gin.Context) {
	// Model Produk
	promod := []models.Produk{}

	settings.DB.Where("status_id = ?", 1).Preload("Kategori").Preload("Setatus").Find(&promod)

	// Iterasi data
	itData := []models.ResProduk{}
	for _, promod := range promod {
		itData = append(itData, FillResProduct(promod))
	}

	// Tampilkan data dengan helper response
	helpers.DataResponse(c, itData)
}
