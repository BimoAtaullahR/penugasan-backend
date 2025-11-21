package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/BimoAtaullahR/penugasan-backend/config"
	"github.com/BimoAtaullahR/penugasan-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context){
	//mendapatkan token dari Header HTTP
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 || tokenString[0] != "Bearer"{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
		return 
	}

	//parse dan validasi token
	token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error){
		//validasi algoritma enkripsi
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	//cek claims dan expiration
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			return
		}
		//cari user di DB
		userID := claims["sub"]
		var user models.User
		database.DB.First(&user, userID)
		if userID==""{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user invalid"})
			return
		}
		
		//attach ke request dan next
		c.Set("user", user)
		c.Next()
	}else{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
	}
	
}