package servidor

import (
	"crud-simples/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID   uint32 `json:"id"`
	Nome string `json:"nome"`
}

// CriarUsuario insere um novo usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}

	fmt.Println(usuario)

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	stratement, erro := db.Prepare("insert into usuarios (nome) values (?)")
	if erro != nil {
		fmt.Println(erro)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer stratement.Close()

	insercao, erro := stratement.Exec(usuario.Nome)
	if erro != nil {
		w.Write([]byte("Erro ao inserir o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	// status code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

}
