package controllers

import (
	"log"
	"loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	findAll := models.FindAll()
	templates.ExecuteTemplate(w, "Index", findAll)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("erro na conversao do preço:", err)
		}
		quantConvert, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("erro na conversao do quantidade:", err)
		}
		models.Save(nome, descricao, precoConvertido, quantConvert)

		http.Redirect(w,r, "/", 301)
	}
}

