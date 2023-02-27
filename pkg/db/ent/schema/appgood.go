package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-manager/pkg/db/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
)

// AppGood holds the schema definition for the AppGood entity.
type AppGood struct {
	ent.Schema
}

func (AppGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

//nolint:funlen
func (AppGood) Fields() []ent.Field {
	lDef := 3000
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
			Bool("online").
			Optional().
			Default(false),
		field.
			Bool("visible").
			Optional().
			Default(true),
		field.
			String("good_name").
			Optional().
			Default(""),
		field.
			Other("price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Int32("display_index").
			Optional().
			Default(0),
		field.
			Int32("purchase_limit").
			Optional().
			Default(int32(lDef)),
		field.
			Int32("commission_percent").
			Optional().
			Default(0),
		field.
			Uint32("sale_start_at").
			Optional().
			Default(0),
		field.
			Uint32("sale_end_at").
			Optional().
			Default(0),
		field.
			Uint32("service_start_at").
			Optional().
			Default(0),
		field.
			Uint32("technical_fee_ratio").
			Optional().
			Default(0),
		field.
			Uint32("electricity_fee_ratio").
			Optional().
			Default(0),
		field.
			Other("daily_reward_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("commission_settle_type").
			Optional().
			Default(commmgrpb.SettleType_NoCommission.String()),
		field.
			JSON("descriptions", []string{}).
			Optional().
			Default([]string{}),
		field.
			String("good_banner").
			Optional().
			Default(""),
		field.
			JSON("display_names", []string{}).
			Optional().
			Default([]string{}),
		field.
			Bool("open_purchase").
			Optional().
			Default(true),
		field.
			Bool("into_product_page").
			Optional().
			Default(true),
		field.
			Bool("can_cancel").
			Optional().
			Default(false),
		field.
			Other("user_purchase_limit", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional(),
	}
}

// Edges of the AppGood.
func (AppGood) Edges() []ent.Edge {
	return nil
}

func (AppGood) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "app_id", "online"),
	}
}
