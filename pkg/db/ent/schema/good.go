package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-manager/pkg/db/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"

	timeconst "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
)

// Good holds the schema definition for the Good entity.
type Good struct {
	ent.Schema
}

func (Good) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Good.
func (Good) Fields() []ent.Field {
	const benefitHours = 24

	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("device_info_id", uuid.UUID{}),
		field.
			Int32("duration_days").
			Optional().
			Default(timeconst.DaysPerYear),
		field.
			UUID("coin_type_id", uuid.UUID{}),
		field.
			UUID("inherit_from_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("vendor_location_id", uuid.UUID{}),
		field.
			Other("price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("benefit_type").
			Optional().
			Default(good.BenefitType_DefaultBenefitType.String()),
		field.
			String("good_type").
			Optional().
			Default(good.GoodType_DefaultGoodType.String()),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			String("unit").
			Optional().
			Default(""),
		field.
			Int32("unit_amount").
			Optional().
			Default(0),
		field.
			JSON("support_coin_type_ids", []uuid.UUID{}).
			Optional().
			Default([]uuid.UUID{}),
		field.
			Uint32("delivery_at").
			Optional().
			Default(0),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Bool("test_only").
			Optional().
			Default(false),
		field.
			Uint32("benefit_interval_hours").
			Optional().
			Default(benefitHours),
	}
}

// Edges of the Good.
func (Good) Edges() []ent.Edge {
	return nil
}

func (Good) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("device_info_id", "vendor_location_id", "benefit_type", "good_type"),
	}
}
