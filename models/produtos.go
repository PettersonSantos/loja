package models

import "loja/config"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindAll() []Produto {
	db := config.ConectaDB()

	findAll, err := db.Query("Select * from produtos order by id asc")

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
		p.Id = id
		p.Nome = nome
		p.Preco = preco
		p.Descricao = descricao
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
		defer db.Close()

		return produtos
}

func Save(nome, descricao string, preco float64, quantidade int) {
	db := config.ConectaDB()
	insert, err := db.Prepare("Insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func Delete(id string) {
	db := config.ConectaDB()
	delete, err := db.Prepare("Delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
}

func FindById(id string) Produto {
	db := config.ConectaDB()

	findById, err := db.Query("Select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	for findById.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = findById.Scan(&id, &descricao, &preco, &quantidade, &nome)

		if err != nil {
			panic(err.Error())
		}
		produto.Id = id
		produto.Nome = nome
		produto.Preco = preco
		produto.Descricao = descricao
		produto.Quantidade = quantidade
	}
	defer db.Close()
	return produto
}

func Update(id int, nome, descricao string, preco float64, quantidade int)  {
	db := config.ConectaDB()
	insert, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5  ")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
