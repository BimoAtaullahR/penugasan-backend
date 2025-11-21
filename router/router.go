package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/BimoAtaullahR/penugasan-backend/config"
	"github.com/BimoAtaullahR/penugasan-backend/models"
	"github.com/BimoAtaullahR/penugasan-backend/controllers"
	"github.com/BimoAtaullahR/penugasan-backend/middleware"
) 

//setuprouter mengatur semua koneksi dan rute, dan mengembalikan *gin.engine
func SetupRouter() *gin.Engine{
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{})
	r := gin.Default()

	//public routes
	// Rute Halaman Utama (Root)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Hai! API untuk penugasan OmahTI Backend sudah online dan siap menerima request.",
			"version": "1.0.0",
		})
	})
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


