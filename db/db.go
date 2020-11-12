package db

import (
	"cattleai/ent"
	"cattleai/ent/migrate"
	"context"
	"fmt"

	"github.com/facebook/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "root"
	dbPswd = "Smile!@#"
	dbAddr = "localhost:3306"
	dbName = "cattlems"
)

var (
	Client *ent.Client
)

func Init() {
	dataSourceName := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8", dbUser, dbPswd, dbAddr, dbName)
	client, err := ent.Open(dialect.MySQL, dataSourceName)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	if err = client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		panic(err)
	}
	Client = client
	dataInit()
}

func Close() {
	if Client != nil {
		Client.Close()
	}
}
