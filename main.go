package main

import (
	"github.com/BimoAtaullahR/penugasan-backend/router"
) 

func main(){
	r := router.SetupRouter()
	r.Run(":7860")
	
}
