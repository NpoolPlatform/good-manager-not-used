// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/appgood"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppGood is the model entity for the AppGood schema.
type AppGood struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// Online holds the value of the "online" field.
	Online bool `json:"online,omitempty"`
	// Visible holds the value of the "visible" field.
	Visible bool `json:"visible,omitempty"`
	// GoodName holds the value of the "good_name" field.
	GoodName string `json:"good_name,omitempty"`
	// Price holds the value of the "price" field.
	Price decimal.Decimal `json:"price,omitempty"`
	// DisplayIndex holds the value of the "display_index" field.
	DisplayIndex int32 `json:"display_index,omitempty"`
	// PurchaseLimit holds the value of the "purchase_limit" field.
	PurchaseLimit int32 `json:"purchase_limit,omitempty"`
	// CommissionPercent holds the value of the "commission_percent" field.
	CommissionPercent int32 `json:"commission_percent,omitempty"`
	// SaleStartAt holds the value of the "sale_start_at" field.
	SaleStartAt uint32 `json:"sale_start_at,omitempty"`
	// SaleEndAt holds the value of the "sale_end_at" field.
	SaleEndAt uint32 `json:"sale_end_at,omitempty"`
	// ServiceStartAt holds the value of the "service_start_at" field.
	ServiceStartAt uint32 `json:"service_start_at,omitempty"`
	// TechnicalFeeRatio holds the value of the "technical_fee_ratio" field.
	TechnicalFeeRatio uint32 `json:"technical_fee_ratio,omitempty"`
	// ElectricityFeeRatio holds the value of the "electricity_fee_ratio" field.
	ElectricityFeeRatio uint32 `json:"electricity_fee_ratio,omitempty"`
	// DailyRewardAmount holds the value of the "daily_reward_amount" field.
	DailyRewardAmount decimal.Decimal `json:"daily_reward_amount,omitempty"`
	// CommissionSettleType holds the value of the "commission_settle_type" field.
	CommissionSettleType string `json:"commission_settle_type,omitempty"`
	// Descriptions holds the value of the "descriptions" field.
	Descriptions []string `json:"descriptions,omitempty"`
	// GoodBanner holds the value of the "good_banner" field.
	GoodBanner string `json:"good_banner,omitempty"`
	// DisplayNames holds the value of the "display_names" field.
	DisplayNames []string `json:"display_names,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppGood) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appgood.FieldDescriptions, appgood.FieldDisplayNames:
			values[i] = new([]byte)
		case appgood.FieldPrice, appgood.FieldDailyRewardAmount:
			values[i] = new(decimal.Decimal)
		case appgood.FieldOnline, appgood.FieldVisible:
			values[i] = new(sql.NullBool)
		case appgood.FieldCreatedAt, appgood.FieldUpdatedAt, appgood.FieldDeletedAt, appgood.FieldDisplayIndex, appgood.FieldPurchaseLimit, appgood.FieldCommissionPercent, appgood.FieldSaleStartAt, appgood.FieldSaleEndAt, appgood.FieldServiceStartAt, appgood.FieldTechnicalFeeRatio, appgood.FieldElectricityFeeRatio:
			values[i] = new(sql.NullInt64)
		case appgood.FieldGoodName, appgood.FieldCommissionSettleType, appgood.FieldGoodBanner:
			values[i] = new(sql.NullString)
		case appgood.FieldID, appgood.FieldAppID, appgood.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppGood", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppGood fields.
func (ag *AppGood) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appgood.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ag.ID = *value
			}
		case appgood.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ag.CreatedAt = uint32(value.Int64)
			}
		case appgood.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ag.UpdatedAt = uint32(value.Int64)
			}
		case appgood.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ag.DeletedAt = uint32(value.Int64)
			}
		case appgood.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ag.AppID = *value
			}
		case appgood.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				ag.GoodID = *value
			}
		case appgood.FieldOnline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field online", values[i])
			} else if value.Valid {
				ag.Online = value.Bool
			}
		case appgood.FieldVisible:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field visible", values[i])
			} else if value.Valid {
				ag.Visible = value.Bool
			}
		case appgood.FieldGoodName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field good_name", values[i])
			} else if value.Valid {
				ag.GoodName = value.String
			}
		case appgood.FieldPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value != nil {
				ag.Price = *value
			}
		case appgood.FieldDisplayIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field display_index", values[i])
			} else if value.Valid {
				ag.DisplayIndex = int32(value.Int64)
			}
		case appgood.FieldPurchaseLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field purchase_limit", values[i])
			} else if value.Valid {
				ag.PurchaseLimit = int32(value.Int64)
			}
		case appgood.FieldCommissionPercent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field commission_percent", values[i])
			} else if value.Valid {
				ag.CommissionPercent = int32(value.Int64)
			}
		case appgood.FieldSaleStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sale_start_at", values[i])
			} else if value.Valid {
				ag.SaleStartAt = uint32(value.Int64)
			}
		case appgood.FieldSaleEndAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sale_end_at", values[i])
			} else if value.Valid {
				ag.SaleEndAt = uint32(value.Int64)
			}
		case appgood.FieldServiceStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field service_start_at", values[i])
			} else if value.Valid {
				ag.ServiceStartAt = uint32(value.Int64)
			}
		case appgood.FieldTechnicalFeeRatio:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field technical_fee_ratio", values[i])
			} else if value.Valid {
				ag.TechnicalFeeRatio = uint32(value.Int64)
			}
		case appgood.FieldElectricityFeeRatio:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field electricity_fee_ratio", values[i])
			} else if value.Valid {
				ag.ElectricityFeeRatio = uint32(value.Int64)
			}
		case appgood.FieldDailyRewardAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field daily_reward_amount", values[i])
			} else if value != nil {
				ag.DailyRewardAmount = *value
			}
		case appgood.FieldCommissionSettleType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commission_settle_type", values[i])
			} else if value.Valid {
				ag.CommissionSettleType = value.String
			}
		case appgood.FieldDescriptions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field descriptions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ag.Descriptions); err != nil {
					return fmt.Errorf("unmarshal field descriptions: %w", err)
				}
			}
		case appgood.FieldGoodBanner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field good_banner", values[i])
			} else if value.Valid {
				ag.GoodBanner = value.String
			}
		case appgood.FieldDisplayNames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field display_names", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ag.DisplayNames); err != nil {
					return fmt.Errorf("unmarshal field display_names: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppGood.
// Note that you need to call AppGood.Unwrap() before calling this method if this AppGood
// was returned from a transaction, and the transaction was committed or rolled back.
func (ag *AppGood) Update() *AppGoodUpdateOne {
	return (&AppGoodClient{config: ag.config}).UpdateOne(ag)
}

// Unwrap unwraps the AppGood entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ag *AppGood) Unwrap() *AppGood {
	_tx, ok := ag.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppGood is not a transactional entity")
	}
	ag.config.driver = _tx.drv
	return ag
}

// String implements the fmt.Stringer.
func (ag *AppGood) String() string {
	var builder strings.Builder
	builder.WriteString("AppGood(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ag.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.AppID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.GoodID))
	builder.WriteString(", ")
	builder.WriteString("online=")
	builder.WriteString(fmt.Sprintf("%v", ag.Online))
	builder.WriteString(", ")
	builder.WriteString("visible=")
	builder.WriteString(fmt.Sprintf("%v", ag.Visible))
	builder.WriteString(", ")
	builder.WriteString("good_name=")
	builder.WriteString(ag.GoodName)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", ag.Price))
	builder.WriteString(", ")
	builder.WriteString("display_index=")
	builder.WriteString(fmt.Sprintf("%v", ag.DisplayIndex))
	builder.WriteString(", ")
	builder.WriteString("purchase_limit=")
	builder.WriteString(fmt.Sprintf("%v", ag.PurchaseLimit))
	builder.WriteString(", ")
	builder.WriteString("commission_percent=")
	builder.WriteString(fmt.Sprintf("%v", ag.CommissionPercent))
	builder.WriteString(", ")
	builder.WriteString("sale_start_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.SaleStartAt))
	builder.WriteString(", ")
	builder.WriteString("sale_end_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.SaleEndAt))
	builder.WriteString(", ")
	builder.WriteString("service_start_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.ServiceStartAt))
	builder.WriteString(", ")
	builder.WriteString("technical_fee_ratio=")
	builder.WriteString(fmt.Sprintf("%v", ag.TechnicalFeeRatio))
	builder.WriteString(", ")
	builder.WriteString("electricity_fee_ratio=")
	builder.WriteString(fmt.Sprintf("%v", ag.ElectricityFeeRatio))
	builder.WriteString(", ")
	builder.WriteString("daily_reward_amount=")
	builder.WriteString(fmt.Sprintf("%v", ag.DailyRewardAmount))
	builder.WriteString(", ")
	builder.WriteString("commission_settle_type=")
	builder.WriteString(ag.CommissionSettleType)
	builder.WriteString(", ")
	builder.WriteString("descriptions=")
	builder.WriteString(fmt.Sprintf("%v", ag.Descriptions))
	builder.WriteString(", ")
	builder.WriteString("good_banner=")
	builder.WriteString(ag.GoodBanner)
	builder.WriteString(", ")
	builder.WriteString("display_names=")
	builder.WriteString(fmt.Sprintf("%v", ag.DisplayNames))
	builder.WriteByte(')')
	return builder.String()
}

// AppGoods is a parsable slice of AppGood.
type AppGoods []*AppGood

func (ag AppGoods) config(cfg config) {
	for _i := range ag {
		ag[_i].config = cfg
	}
}
