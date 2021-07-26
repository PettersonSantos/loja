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

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.Delete(idProduto)
	http.Redirect(w,r,"/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.FindById(id)
	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if(r.Method == "POST") {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("erro na conversao do Id:", err)
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("erro na conversao do preço:", err)
		}
		quantConvert, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("erro na conversao do quantidade:", err)
		}
		models.Update(idConvertido, nome, descricao, precoConvertido, quantConvert)

		http.Redirect(w,r, "/", 301)
	}
}

