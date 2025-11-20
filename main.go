package main

import (
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/jackc/pgx/v5"
	// "net/http"
	"github.com/BimoAtaullahR/penugasan-backend/config"
	"github.com/BimoAtaullahR/penugasan-backend/models"
	"github.com/BimoAtaullahR/penugasan-backend/controllers"
) 

func main(){
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{})
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.Run()
	// router.Run()
}
