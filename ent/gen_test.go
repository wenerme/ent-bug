package ent

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
	"testing"
)

func TestGen(t *testing.T) {
	err := entc.Generate("./schema", &gen.Config{
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
