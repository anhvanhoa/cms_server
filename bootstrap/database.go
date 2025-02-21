package bootstrap

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func NewPostgresDB(env *Env, entities []interface{}) *pg.DB {

	// Connect to the database
	opt, err := pg.ParseURL(env.URL_DB)
	if err != nil {
		log.Fatal("Error parsing the database URL:", err)
	}

	db := pg.Connect(opt)
	// Check the connection
	if err := db.Ping(db.Context()); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Create the tables
	for _, entity := range entities {
		err := db.Model(entity).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			log.Fatal("Error creating table:", err)
		}
	}

	return db
}

func ClosePostgresDB(db *pg.DB) {
	db.Close()
}
