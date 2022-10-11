package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// DeviceInfo holds the schema definition for the DeviceInfo entity.
type DeviceInfo struct {
	ent.Schema
}

func (DeviceInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the DeviceInfo.
func (DeviceInfo) Fields() []ent.Field {
	maxLen := 64
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			String("type").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.
			String("manufacturer").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.
			Uint32("power_comsuption").
			Optional().
			Default(0),
		field.
			Uint32("shipment_at").
			Optional().
			Default(0),
		field.
			JSON("posters", []string{}).
			Optional().
			Default([]string{}),
	}
}

// Edges of the DeviceInfo.
func (DeviceInfo) Edges() []ent.Edge {
	return nil
}
