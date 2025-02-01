package controllers

import (
	"test-programmer/helpers"
	"test-programmer/models"
	"test-programmer/settings"

	"github.com/gin-gonic/gin"
)

func FillResKategori(cat models.Kategori) models.ResKategori {
	return models.ResKategori{
		ID:   cat.ID,
		Nama: cat.Nama,
	}
}

func FillResStatus(sta models.Status) models.ResStatus {
	return models.ResStatus{
		ID:   sta.ID,
		Nama: sta.Nama,
	}
}

func DataKategorySupport(c *gin.Context) {
	// Variabel model Kategori
	modCat := []models.Kategori{}

	// Panggil data kategori
	settings.DB.Find(&modCat)

	// Iterasi data kategori
	resCat := []models.ResKategori{}
	for _, modCat := range modCat {
		resCat = append(resCat, FillResKategori(modCat))
	}

	// Tampilkan data response
	helpers.DataResponse(c, resCat)
}

func DataStatusSupport(c *gin.Context) {
	// Variabel model Status
	modSta := []models.Status{}

	// Panggil data status
	settings.DB.Find(&modSta)

	// Iterasi data status
	resSta := []models.ResStatus{}
	for _, modSta := range modSta {
		resSta = append(resSta, FillResStatus(modSta))
	}

	// Tampilkan data response
	helpers.DataResponse(c, resSta)
}
