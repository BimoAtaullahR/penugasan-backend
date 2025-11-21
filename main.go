package main

import (
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/jackc/pgx/v5"
	"net/http"
	"github.com/BimoAtaullahR/penugasan-backend/config"
	"github.com/BimoAtaullahR/penugasan-backend/models"
	"github.com/BimoAtaullahR/penugasan-backend/controllers"
	"github.com/BimoAtaullahR/penugasan-backend/middleware"
) 

var r *gin.Engine

//dijalankan terlebih dahulu secara otomatis baik di vercel ataupun di local
func init(){
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{})
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middleware.RequireAuth)

	api.GET("/me", func(ctx *gin.Context) {
		user, exists := ctx.Get("user")
		if !exists{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user data not found"})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message" : "ini halaman /me",
			"user": user,
		})
	})
}

//untuk vercel
func Handler(w http.ResponseWriter, req *http.Request){
	r.ServeHTTP(w, req)
}

//fungsi untuk ketika menjalankannya di local
func main(){
	r.Run()
	// router.Run()
}
