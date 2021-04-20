package test

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wenerme/ent-demo/ent"
	"github.com/wenerme/ent-demo/ent/enttest"
	"log"
	"testing"
)

func TestClear(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
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
