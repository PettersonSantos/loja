package controllers

import (
	"loja/models"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	findAll := models.FindAll()
	templates.ExecuteTemplate(w, "Index", findAll)
}

