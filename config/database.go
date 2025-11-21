package database

import(
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	// "fmt"
	"log"
	"os"
)

var DB *gorm.DB


func ConnectDatabase() {
	// err := godotenv.Load()
	// if err != nil{
	// 	log.Fatal("Error loading .env file")
	// }

	godotenv.Load()

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