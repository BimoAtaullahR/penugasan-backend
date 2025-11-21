package router

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

//setuprouter mengatur semua koneksi dan rute, dan mengembalikan *gin.engine
func SetupRouter() *gin.Engine{
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{})
	r := gin.Default()

	//public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	//protected routes
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

	return r
}


