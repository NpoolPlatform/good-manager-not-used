package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// AppDefaultGood holds the schema definition for the AppDefaultGood entity.
type AppDefaultGood struct {
	ent.Schema
}

func (AppDefaultGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (AppDefaultGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			UUID("coin_type_id", uuid.UUID{}),
	}
}

// Edges of the AppDefaultGood.
func (AppDefaultGood) Edges() []ent.Edge {
	return nil
}
