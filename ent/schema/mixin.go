package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type AuditorMixin struct {
	mixin.Schema
}

func (AuditorMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("createdByID").Optional().Nillable(),
	}
}
func (AuditorMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("createdByID"),
	}
}
func (AuditorMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("createdBy", User.Type).From("creator").Field("createdByID").Unique(),
	}
}
