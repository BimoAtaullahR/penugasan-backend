package handler

import(
	"github.com/BimoAtaullahR/penugasan-backend/router"
	"net/http"
)

var r = router.SetupRouter()

func Handler(w http.ResponseWriter, req *http.Request){
	r.ServeHTTP(w, req)
}