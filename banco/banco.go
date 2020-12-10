package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	stringConexao := "root:123456Asdf@/cursogo?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	// testando conexao com banco
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	//fmt.Println("Conexão banco dados está aberta")
	return db, nil
}
