package test

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
	"github.com/wenerme/ent-demo/ent"
	"github.com/wenerme/ent-demo/ent/enttest"
	"github.com/wenerme/ent-demo/ent/migrate"
	"log"
	"testing"
)

func TestPGMir(t *testing.T) {
	db, err := sql.Open("postgres", "user=test password=test dbname=test host=127.0.0.1 port=15432 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	var driver *entsql.Driver
	driver = entsql.OpenDB("postgres", db)

	mir, err := schema.NewMigrate(driver)
	if err != nil {
		panic(err)
	}
	err = mir.Create(context.Background(),
		migrate.UsersTable,
		migrate.PetsTable,
	)
	if err != nil {
		panic(err)
	}
}

func TestPG(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "postgres", "user=test password=test dbname=test host=127.0.0.1 port=15432 sslmode=disable TimeZone=Asia/Shanghai", opts...)
	client = client.Debug()
	defer client.Close()
	ctx := context.Background()

	userA := client.User.Create().SetName("name").SaveX(ctx)
	spew.Dump(*userA)
	petA := client.Pet.Create().SetName("dog").SetOwningUserID(userA.ID).SaveX(ctx)
	log.Println("clear user id")
	spew.Dump(*petA)
	petA.Update().ClearOwningUserID().SaveX(ctx)
	log.Println("clear user")
	petA.Update().ClearOwningUser().SaveX(ctx)
}
