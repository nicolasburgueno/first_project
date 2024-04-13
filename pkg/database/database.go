package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/nicolasburgueno/first_project/internal/config"
)

func GetDB() *sqlx.DB {
	db, err := sqlx.Connect(config.DBDriver, config.DBDSN)
	if err != nil {
		log.Panic("[event:db_connection_error][]", err)
		panic(err)
	}

	log.Println("db connection established")
	return db
}
