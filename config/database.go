package database

import(
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
	"os"
)

var DB *gorm.DB


func ConnectDatabase() {
	godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Database Url is not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil{
		log.Fatal("Failed to connect to database")
	}

	DB = db
	log.Println("Database connection successful!")
}