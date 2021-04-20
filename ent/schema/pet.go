package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Pet struct {
	ent.Schema
}

func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		HasOwnerMixin{},
	}
}
