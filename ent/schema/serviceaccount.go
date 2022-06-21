package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ServiceAccount holds the schema definition for the ServiceAccount entity.
type ServiceAccount struct {
	ent.Schema
}

// Fields of the ServiceAccount.
func (ServiceAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("displayName").NotEmpty(),
		field.Bool("disabled").Default(false),
		field.String("username"),
		field.String("password"),
	}
}

// Edges of the ServiceAccount.
func (ServiceAccount) Edges() []ent.Edge {
	return nil
}

func (ServiceAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ID2Mixin{},
	}
}
