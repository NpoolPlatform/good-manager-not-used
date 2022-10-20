package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// Recommend holds the schema definition for the Recommend entity.
type Recommend struct {
	ent.Schema
}

func (Recommend) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Recommend.
func (Recommend) Fields() []ent.Field {
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
			UUID("recommender_id", uuid.UUID{}).
			Optional().
			Default(
				func() uuid.UUID {
					return uuid.UUID{}
				}),
		field.
			String("message").
			Optional().
			Default(""),
		field.
			Float("recommend_index").
			Optional().
			Default(0),
	}
}

// Indexes of the Recommend.
func (Recommend) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "good_id"),
	}
}
