package routes

import (
	"log"
	"net/http"
	"test-programmer/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebAppRoute() {
	// Memanggil fungsi route dari framework gin golang
	router := gin.Default()

	// Menambahkan cors pada settingan route gin golang
	router.Use(cors.Default())

	// Menggunakan middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://example.com"}, // Ganti dengan origin yang diizinkan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Memuat template dengan ekstensi .tmpl dari direktori view
	router.LoadHTMLGlob("views/*.tmpl")

	// Menyajikan file statis dari direktori assets
	router.Static("/assets", "./assets")

	// Route Website
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Index.tmpl", nil) // Render template Login.tmpl
	})

	// Route API
	v1 := router.Group("api/v1/")
	{
		proMenu := v1.Group("/produk/")
		{
			proMenu.GET("/list", controllers.ListProduk)
		}

		supportMenu := v1.Group("/support-data/")
		{
			supportMenu.GET("/kategori", controllers.DataKategorySupport)
			supportMenu.GET("/status", controllers.DataStatusSupport)
		}
	}

	// Menampilkan log server berjalan dengan port 8080
	log.Println("Server started on: http://127.0.0.1:8080")

	// Menjalankan server ke port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
