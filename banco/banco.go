package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func criacaoBanco(db sql.Result) {

}

func Conectar() (*sql.DB, error) {
	stringConexao := "root:123456Asdf@/cursogo?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	//exec(db, "drop table if exists usuarios")
	exec(db, `create table if not exists usuarios (
		id integer auto_increment,
		nome varchar(80),
		email varchar(80),
		senha varchar(32),
		PRIMARY KEY (id)
		) ENGINE=InnoDB;`)

	// testando conexao com banco
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	//fmt.Println("Conexão banco dados está aberta")
	return db, nil
}
