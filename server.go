package main

import (
	"test-programmer/routes"
	"test-programmer/seeders"
	"test-programmer/settings"
)

func main() {
	// Panggil fungsi koneksi ke database
	settings.InitDB()
	defer settings.DB.Close()

	// Seed data otomatis
	seeders.StatusSeed(nil)
	seeders.ProdukSeed(nil)

	// Panggil fungsi route webapp
	routes.WebAppRoute()
}
