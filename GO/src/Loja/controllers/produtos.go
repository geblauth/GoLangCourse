package controllers

import (
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todoOsProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todoOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Error", err)
		}

		quantidadeCovertido, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Error", err)
		}
		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeCovertido)

	}

	http.Redirect(w, r, "/", 301)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvert, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}

		precoConvert, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println(err)
		}

		quantidadeConvert, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("a", err)
		}

		models.AtualizaProduto(idConvert, nome, descricao, precoConvert, quantidadeConvert)

	}
	http.Redirect(w, r, "/", 301)

}
