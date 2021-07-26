package routes

import (
	"loja/controllers"
	"net/http"
)


func CarregaRotas()  {
	http.HandleFunc("/", controllers.Index)
}
