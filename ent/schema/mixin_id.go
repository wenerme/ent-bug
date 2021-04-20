package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/wenerme/ent-demo/models"
	"time"
)

var idType = new(models.ID)

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(idType).Unique().Immutable().Annotations(),
	}
}
func (IDMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		HookID(),
	}
}

func HookID() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			if mutation.Op().Is(ent.OpCreate) {
				s := mutation.(interface {
					SetID(v models.ID)
				})
				id, err := NextID(ctx, mutation)
				if err != nil {
					return nil, err
				}
				s.SetID(id)
			}
			return next.Mutate(ctx, mutation)
		})
	}
}

var NextID = func(ctx context.Context, mutation ent.Mutation) (models.ID, error) {
	// test only
	return models.ID(time.Now().String()), nil
}
