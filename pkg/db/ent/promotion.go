// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/promotion"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Promotion is the model entity for the Promotion schema.
type Promotion struct {
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
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt uint32 `json:"start_at,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt uint32 `json:"end_at,omitempty"`
	// Price holds the value of the "price" field.
	Price decimal.Decimal `json:"price,omitempty"`
	// Posters holds the value of the "posters" field.
	Posters []string `json:"posters,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Promotion) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case promotion.FieldPosters:
			values[i] = new([]byte)
		case promotion.FieldPrice:
			values[i] = new(decimal.Decimal)
		case promotion.FieldCreatedAt, promotion.FieldUpdatedAt, promotion.FieldDeletedAt, promotion.FieldStartAt, promotion.FieldEndAt:
			values[i] = new(sql.NullInt64)
		case promotion.FieldMessage:
			values[i] = new(sql.NullString)
		case promotion.FieldID, promotion.FieldAppID, promotion.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Promotion", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Promotion fields.
func (pr *Promotion) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case promotion.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case promotion.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = uint32(value.Int64)
			}
		case promotion.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = uint32(value.Int64)
			}
		case promotion.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = uint32(value.Int64)
			}
		case promotion.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				pr.AppID = *value
			}
		case promotion.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				pr.GoodID = *value
			}
		case promotion.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				pr.Message = value.String
			}
		case promotion.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				pr.StartAt = uint32(value.Int64)
			}
		case promotion.FieldEndAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				pr.EndAt = uint32(value.Int64)
			}
		case promotion.FieldPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value != nil {
				pr.Price = *value
			}
		case promotion.FieldPosters:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field posters", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Posters); err != nil {
					return fmt.Errorf("unmarshal field posters: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Promotion.
// Note that you need to call Promotion.Unwrap() before calling this method if this Promotion
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Promotion) Update() *PromotionUpdateOne {
	return (&PromotionClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Promotion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Promotion) Unwrap() *Promotion {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Promotion is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Promotion) String() string {
	var builder strings.Builder
	builder.WriteString("Promotion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.AppID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.GoodID))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(pr.Message)
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.StartAt))
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.EndAt))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("posters=")
	builder.WriteString(fmt.Sprintf("%v", pr.Posters))
	builder.WriteByte(')')
	return builder.String()
}

// Promotions is a parsable slice of Promotion.
type Promotions []*Promotion

func (pr Promotions) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
