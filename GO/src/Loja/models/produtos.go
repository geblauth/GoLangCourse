package models

import "loja/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	insereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	insereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}
func BuscaTodosProdutos() []Produto {

	db := db.ConectaBanco()

	selectDeTodosProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos

}
func DeletaProduto(id string) {
	db := db.ConectaBanco()
	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
	defer db.Close()

}

func EditaProduto(id string) Produto {
	db := db.ConectaBanco()
	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoAtualizar.Id = id
		produtoAtualizar.Nome = nome
		produtoAtualizar.Descricao = descricao
		produtoAtualizar.Preco = preco
		produtoAtualizar.Quantidade = quantidade

	}
	defer db.Close()

	return produtoAtualizar

}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
