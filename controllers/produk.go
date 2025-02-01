package controllers

import (
	"errors"
	"strconv"
	"test-programmer/helpers"
	"test-programmer/models"
	"test-programmer/settings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

func CreateProduk(c *gin.Context) {
	// Definisi body request ke variabel
	kategoriid := c.PostForm("kategori_id")
	nama := c.PostForm("nama")
	harga := c.PostForm("harga")
	statusid := c.PostForm("status_id")

	// Konversi tipe data integer
	konv1, _ := strconv.Atoi(kategoriid)
	konv2, _ := strconv.Atoi(harga)
	konv3, _ := strconv.Atoi(statusid)

	// Struct simpan data
	simpan := models.Produk{
		KategoriId: uint(konv1),
		StatusId:   uint(konv3),
		Nama:       nama,
		Harga:      int64(konv2),
	}

	// Simpan struct data ke database
	settings.DB.Create(&simpan)

	// Simpan struct data ke variabel response
	res := FillResProduct(simpan)

	// Tampilkan response data
	helpers.SuksesWithDataResponse(c, "Berhasil membuat produk baru!", res)
}

func UpdateProduk(c *gin.Context) {
	// Definisikan id param dan body request kedalam variabel
	paramid := c.Param("id")
	idkategori := c.PostForm("kategori_id")
	idstatus := c.PostForm("status_id")
	nama := c.PostForm("nama")
	harga := c.PostForm("harga")

	// Variabel model produk
	var pro models.Produk

	// Kondisi bila data tidak ada
	if errors.Is(settings.DB.Where("id = ?", paramid).First(&pro).Error, gorm.ErrRecordNotFound) {
		helpers.ErrorResponse(c, "Data Produk tidak ditemukan!")
		c.Abort()
		return
	}

	// Konversi tipe data integer
	konv1, _ := strconv.Atoi(idkategori)
	konv2, _ := strconv.Atoi(idstatus)
	konv3, _ := strconv.Atoi(harga)

	// Update data dalam database
	settings.DB.Model(&pro).Where("id = ?", paramid).Updates(models.Produk{
		KategoriId: uint(konv1),
		StatusId:   uint(konv2),
		Nama:       nama,
		Harga:      int64(konv3),
	})

	// Simpan hasil update ke variabel response
	res := FillResProduct(pro)

	// Tampilkan response sukses
	helpers.SuksesWithDataResponse(c, "Berhasil melakukan update data produk!", res)
}

func HapusProduk(c *gin.Context) {
	// Definisikan id param kedalam variabel
	paramid := c.Param("id")

	// Variabel model produk
	var pro models.Produk

	// Kondisi bila data tidak ada
	if errors.Is(settings.DB.Where("id = ?", paramid).First(&pro).Error, gorm.ErrRecordNotFound) {
		helpers.ErrorResponse(c, "Data Produk tidak ditemukan!")
		c.Abort()
		return
	}

	// Hapus data produk berdasarkan param id
	settings.DB.Where("id = ?", paramid).Delete(&pro)

	// Helper response sukses
	helpers.SuksesResponse(c, "Berhasil melalukan hapus data produk!")
}
