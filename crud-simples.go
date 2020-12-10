package main

import (
	"crud-simples/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/*
		stringConexao := "root:123456Asdf@/cursogo?charset=utf8&parseTime=True&loc=Local"

		db, erro := sql.Open("mysql", stringConexao)
		if erro != nil {
			log.Fatal(erro)
		}
		defer db.Close()

		// testando conexao com banco
		if erro = db.Ping(); erro != nil {
			log.Fatal(erro)
		}

		fmt.Println("Conexão banco dados está aberta")

		linhas, erro := db.Query("select * from usuarios")
		if erro != nil {
			log.Fatal(erro)
		}
		defer linhas.Close()

		//fmt.Println(linhas)
	*/

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
