package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// SubGood holds the schema definition for the SubGood entity.
type SubGood struct {
	ent.Schema
}

func (SubGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the SubGood.
func (SubGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("main_good_id", uuid.UUID{}),
		field.
			UUID("sub_good_id", uuid.UUID{}),
		field.
			Bool("must").
			Optional().
			Default(false),
		field.
			Bool("commission").
			Optional().
			Default(false),
	}
}

// Edges of the SubGood.
func (SubGood) Edges() []ent.Edge {
	return nil
}

func (SubGood) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "main_good_id", "sub_good_id"),
	}
}
