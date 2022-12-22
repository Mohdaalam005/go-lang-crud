package database

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/mohdaalam005/crud-with-postgres/models"
)

func Connect() *pg.DB {
	psqlConn := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}
	var db *pg.DB = pg.Connect(psqlConn)

	if db == nil {
		log.Fatal("failed to connect to database")
	}
	log.Println("successfully connected to database")

	err := createSchema(db)
	CheckNilErr(err)
	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.Book)(nil),
		(*models.Author)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Fatal(err)
		}

	}
	return nil
}

func CheckNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
