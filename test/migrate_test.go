package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	atlasmigrate "ariga.io/atlas/sql/migrate"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/wenerme/ent-demo/ent"
)

func TestMigration(t *testing.T) {
	fmt.Println("Drivers", sql.Drivers())
	db, err := sql.Open("pgx", "user=test password=test dbname=test host=127.0.0.1 port=15432 sslmode=disable TimeZone=Asia/Shanghai search_path=dev")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("create schema if not exists dev")
	if err != nil {
		panic(err)
	}

	var driver *entsql.Driver
	driver = entsql.OpenDB("postgres", db)
	ctx := context.Background()
	client := ent.NewClient(ent.Driver(driver))

	dir, err := atlasmigrate.NewLocalDir("../migrations")
	if err != nil {
		panic(err)
	}

	err = client.Schema.Create(ctx,
		schema.WithAtlas(true),
		schema.WithDir(dir),
		schema.WithSumFile(),
	)
	if err != nil {
		panic(err)
	}
}
