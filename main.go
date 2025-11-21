package main

import (
	// "time"

	// "github.com/gin-gonic/gin"
	// // "github.com/jackc/pgx/v5"
	// "net/http"
	// "github.com/BimoAtaullahR/penugasan-backend/config"
	// "github.com/BimoAtaullahR/penugasan-backend/models"
	// "github.com/BimoAtaullahR/penugasan-backend/controllers"
	"github.com/BimoAtaullahR/penugasan-backend/router"
) 


//fungsi untuk ketika menjalankannya di local
func main(){
	r := router.SetupRouter()
	r.Run(":7860")
	
}
