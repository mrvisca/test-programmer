package seeders

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"test-programmer/models"
	"test-programmer/settings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
)

func StatusSeed(c *gin.Context) {
	// Variabel Hitung data
	var hitungStatus int64

	// Variabel Model Status
	modStatus := []models.Status{}

	// Hitung Record
	settings.DB.Model(&modStatus).Count(&hitungStatus)

	// Kondisi Bila Belum Ada Data / Sudah Ada Data
	if hitungStatus == 0 {
		// Struct Simpan Data
		simStatus1 := models.Status{
			Nama: "bisa dijual",
		}

		simStatus2 := models.Status{
			Nama: "tidak bisa dijual",
		}

		// Simpan Struct Data
		settings.DB.Create(&simStatus1)
		settings.DB.Create(&simStatus2)

		// Tampilkan Printline Seed Sukses
		log.Println("Seeder Status Produk Sukses Dibuat!")
	} else {
		// Tampilkan Printline Seed Sukses
		log.Println("Seeder Status Produk Siap Digunakan!")
	}
}

func ProdukSeed(c *gin.Context) {
	// Variabel Hitung data
	var hitungProduk int64

	// Variabel Model Status
	modPro := []models.Produk{}

	// Hitung Record
	settings.DB.Model(&modPro).Count(&hitungProduk)

	// Kondisi Bila Belum Ada Data / Sudah Ada Data
	if hitungProduk == 0 {

		client := resty.New()
		url := "https://recruitment.fastprint.co.id/tes/api_tes_programmer"

		// Kirim POST request dengan Form Data
		response, err := client.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(map[string]string{
				"username": "tesprogrammer010225C21",
				"password": "439e9ae1df04d71e83defc0478ff65bc",
			}).
			Post(url)

		if err != nil {
			log.Println("Error saat request API:", err)
			return
		}

		// Parsing response body ke dalam map
		var data map[string]interface{}
		err = json.Unmarshal(response.Body(), &data)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			return
		}

		// Ambil value dari key "data", yang berisi array produk
		productsData, ok := data["data"].([]interface{})
		if !ok {
			log.Println("Data tidak ditemukan atau format tidak sesuai.")
			return
		}

		// Loop melalui array produk dan ambil id_produk, kategori, dan nama_produk
		for _, product := range productsData {
			productMap, ok := product.(map[string]interface{})
			if !ok {
				log.Println("Gagal parse produk.")
				continue
			}

			// Ambil nilai key yang diperlukan
			namaProduk := fmt.Sprintf("%v", productMap["nama_produk"])
			kategori := fmt.Sprintf("%v", productMap["kategori"])
			harga := fmt.Sprintf("%v", productMap["harga"])
			status := fmt.Sprintf("%v", productMap["status"])

			// Model Kategori
			var cat models.Kategori

			// Konversi tipe data id dan harga
			konv, _ := strconv.Atoi(harga)

			// Kondisi bila data kategori tidak ada
			if errors.Is(settings.DB.Where("nama LIKE ?", "%"+kategori+"%").First(&cat).Error, gorm.ErrRecordNotFound) {
				// Struct data simpan kategori baru
				newCat := models.Kategori{
					Nama: kategori,
				}
				settings.DB.Create(&newCat)

				if status == "bisa dijual" {
					// Struct data simpan produk baru
					newPro := models.Produk{
						KategoriId: newCat.ID,
						StatusId:   1,
						Nama:       namaProduk,
						Harga:      int64(konv),
					}

					// Simpan Data
					settings.DB.Create(&newPro)
				} else {
					// Struct data simpan produk baru
					newPro := models.Produk{
						KategoriId: newCat.ID,
						StatusId:   2,
						Nama:       namaProduk,
						Harga:      int64(konv),
					}

					// Simpan Data
					settings.DB.Create(&newPro)
				}
			}

			// Kondisi bila "bisa dijual" / "tidak bisa dijual"
			if status == "bisa dijual" {
				// Styruct data simpan produk baru
				proNew1 := models.Produk{
					KategoriId: cat.ID,
					StatusId:   1,
					Nama:       namaProduk,
					Harga:      int64(konv),
				}

				// Simpan data
				settings.DB.Create(&proNew1)
			} else {
				// Styruct data simpan produk baru
				proNew2 := models.Produk{
					KategoriId: cat.ID,
					StatusId:   2,
					Nama:       namaProduk,
					Harga:      int64(konv),
				}

				// Simpan data
				settings.DB.Create(&proNew2)
			}
		}

		// Tampilkan Printline Seed Sukses
		log.Println("Seeder Fetch Produk Sukses Dibuat!")
	} else {
		// Tampilkan Printline Seed Sukses
		log.Println("Seeder Fetch Produk Siap Digunakan!")
	}
}
