package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, erro := sql.Open("mysql", "root:rootroot@/?parseTime=true")

	db.Exec("create database if not exists dev")
	db.Exec("use dev")
	db.Exec(`create table if not exists user(
		id integer auto_increment,
		name VARCHAR(60) NOT NULL,
		email VARCHAR(60) NOT NULL UNIQUE,
		password VARCHAR(60) NOT NULL,
		created TIMESTAMP DEFAULT CURRENT_TIMESTAMP() ,
		PRIMARY KEY(id)
	)`)

	if erro != nil {
		panic(erro)
	}
	if erro := db.Ping(); erro != nil {
		panic(erro)
	}
	return db, nil
}
