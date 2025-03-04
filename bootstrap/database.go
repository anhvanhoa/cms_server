package bootstrap

import (
	"cms-server/pkg/database"
	pkglog "cms-server/pkg/logger"

	"github.com/go-pg/pg/extra/pgdebug/v10"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func NewPostgresDB(env *Env, entities []any, log pkglog.Logger) *pg.DB {
	// Connect to the database
	opt, err := pg.ParseURL(env.URL_DB)
	if err != nil {
		log.Fatal("Error parsing the database URL: " + err.Error())
	}

	db := pg.Connect(opt)
	// Check the connection
	if err := db.Ping(db.Context()); err != nil {
		log.Fatal("Error connecting to the database: " + err.Error())
	}

	if !env.IsProduction() {
		db.AddQueryHook(pgdebug.NewDebugHook())
		db.AddQueryHook(database.NewQueryHook())
	}
	// Create the tables
	for _, entity := range entities {
		err := db.Model(entity).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			log.Fatal("Error creating table: " + err.Error())
		}
	}

	return db
}
