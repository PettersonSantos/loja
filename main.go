package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"text/template"
)

func conectaDB() *sql.DB {
	conexao := "user=postgres dbname=loja_1 password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaDB()

	findAll, err := db.Query("Select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for findAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = findAll.Scan(&id, &descricao, &preco, &quantidade, &nome)

		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Preco = preco
		p.Descricao = descricao
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
