package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var instance *sqlx.DB
var err error

func Connect() *sqlx.DB {
	if instance == nil {
		path := os.Getenv("POUPANIQUEL_DB_PATH")
		if path == "" {
			path = "./poupaniquel.db"
		}
		instance, err = sqlx.Connect("sqlite3", path)
	}
	if err != nil {
		log.Fatalln("Error connecting to the database", err)
	}
	instance.SetMaxIdleConns(3)
	return instance
}

func Create() *sqlx.DB {
	transactions := `create table if not exists transactions (
		id integer primary key autoincrement,
		createdAt date not null default CURRENT_DATE,
		type text not null default "expense",
		description text not null,
		amount numeric not null default 0,
		tags text not null default "",
		parentId integer not null default 0
	)`
	tables := []string{transactions}
	
	db := Connect()
	for _, createTable := range tables {
		db.MustExec(createTable)
	}
	return db;
}

func BulkInsert(db *sqlx.DB, data []string) {
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec(item)
	}
	tx.Commit()
}