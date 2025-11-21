package controllers

import (
	"net/http"

	"github.com/BimoAtaullahR/penugasan-backend/config"
	"github.com/BimoAtaullahR/penugasan-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
)

type RegisterInput struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"` //validasi minimal 8 karakter
}

type LoginInput struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context){
	var input RegisterInput

	//validasi input
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	//hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses password"})
		return 
	}

	user := models.User{
		Email: input.Email,
		Password: string(hashedPassword),
	}	

	if err := database.DB.Create(&user).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Gagal mendaftar, email mungkin sudah terdaftar sebelumnya"})
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Registrasi berhasil!"})
}

func Login(c *gin.Context){
	var input LoginInput

	//validasi input JSON
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	var user models.User

	//cari user di database berdasarkan email
	if err := database.DB.Where("email: ?", input.Email).First(&user).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email atau password salah"})
		return
	}

	//verifikasi password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email atau password salah"})
		return
	}

	//generate JWT Token
	//membuat payload atau isi data dalam token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub" : user.ID,								//subject: pemilik dari token ini (id user)
		"exp" : time.Now().Add(time.Hour*24).Unix(),	//token kadaluarsa dalam 24 jam
	})

	//menandatangani token dengan secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	//mengirim token ke user
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}