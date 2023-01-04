package db

import (
	"context"

	"entgo.io/ent/dialect"

	"notify-api/utils/config"
	"notify-api/utils/ds"

	"notify-api/ent/generate"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Client *generate.Client

var RWLock ds.Mutex

func Init() {
	var err error

	var db string
	switch config.Config.Database.Type {
	case config.Sqlite:
		db = dialect.SQLite
	case config.Mysql:
		db = dialect.MySQL
	case config.Pgsql:
		db = dialect.Postgres
	default:
		panic("Unsupported database type")
	}

	Client, err = generate.Open(db, config.Config.Database.DSN)
	if err != nil {
		zap.S().Fatalf("Failed to connect database: %+v", err)
	}

	if err := Client.Schema.Create(context.Background()); err != nil {
		zap.S().Fatalf("Failed to create schema resources: %+v", err)
	}
}