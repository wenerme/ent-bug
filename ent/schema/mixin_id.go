package schema

import (
	"context"

	"github.com/wenerme/ent-demo/utils/ulids"
	"github.com/wenerme/wego/stdx"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/wenerme/ent-demo/models"
)

var idPtrType = models.ID("")

// var idPtrType = (*models.ID)(nil)
var idType = models.ID("")

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(models.ID("")).Unique().Immutable().Annotations(),
		//field.UUID("uid", new(uuid.UUID)).Default(func() *uuid.UUID {
		//	u := uuid.New()
		//	return &u
		//}).Optional().Annotations(entsql.Annotation{}),
		field.UUID("uid", uuid.New()).Default(uuid.New).Unique().Immutable().Annotations(
			entsql.Annotation{
				// 需要针对 UUID 特殊处理 - https://github.com/ent/ent/pull/1462
				// unsupported default type: [16]byte
				// https://github.com/chris-rock/ent/blob/7b4e3e1aa607fa648fc508b58fbe40b0779c369a/dialect/sql/schema/schema.go#L267
				Default: "gen_random_uuid()",
			},
		),
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
	return models.ID(uuid.New().String()), nil
}

type ID2Mixin struct {
	mixin.Schema
}

func (ID2Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			return stdx.Must(ulids.New()).String()
		}).Unique().Immutable().Annotations(
			&entsql.Annotation{
				Default: "public.generate_ulid()",
			},
			// entgql.OrderField("Id"),
			// FieldAnnotation{CanUpdate: TernaryNo, CanCreate: TernaryNo, CanSort: TernaryYes},
		),
	}
}

type TenantSidMixin struct {
	mixin.Schema
}

func (TenantSidMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sid").Immutable().Annotations(
		// entgql.OrderField("Sid"),
		// FieldAnnotation{CanUpdate: TernaryNo, CanCreate: TernaryNo, CanSort: TernaryYes},
		),
		field.Int("tid").Immutable().Annotations(
			// FieldAnnotation{CanUpdate: TernaryNo, CanCreate: TernaryNo, ServerOnly: TernaryYes},
			&entsql.Annotation{Default: "current_setting('tenant.id')::bigint"},
		),
	}
}
