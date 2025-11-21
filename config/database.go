package database

import(
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
	"strings"
	"log"
	"os"
)

var DB *gorm.DB


func ConnectDatabase() {
	// err := godotenv.Load()
	// if err != nil{
	// 	log.Fatal("Error loading .env file")
	// }

	_ = godotenv.Load()

	// --- DEBUGGING START ---
	// 1. Cek apakah variabel terbaca
	val := os.Getenv("DATABASE_URL")
	fmt.Printf("DEBUG: Panjang karakter DATABASE_URL adalah: %d\n", len(val))
	
	// 2. Cek apakah ada typo atau variabel lain yang masuk
	fmt.Println("DEBUG: Daftar Env Var yang tersedia:")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		// Cetak nama variabelnya saja
		if len(pair) > 0 {
			fmt.Println("- " + pair[0])
		}
	}
	// --- DEBUGGING END ---
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == ""{
		log.Fatal("Database Url is not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil{
		log.Fatal("Failed to connect to database")
	}

	DB = db
	log.Println("Database connection successful!")
}