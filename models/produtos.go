package models

import "loja/config"

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindAll() []Produto {
	db := config.ConectaDB()

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
		defer db.Close()

		return produtos
}
