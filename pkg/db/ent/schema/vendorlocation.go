package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// VendorLocation holds the schema definition for the VendorLocation entity.
type VendorLocation struct {
	ent.Schema
}

func (VendorLocation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the VendorLocation.
func (VendorLocation) Fields() []ent.Field {
	maxLen := 128
	addressMaxLen := 256
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("country").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.String("province").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.String("city").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.String("address").
			Optional().
			Default("").
			MaxLen(addressMaxLen),
	}
}

// Edges of the VendorLocation.
func (VendorLocation) Edges() []ent.Edge {
	return nil
}
