// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/subgood"
	"github.com/google/uuid"
)

// SubGood is the model entity for the SubGood schema.
type SubGood struct {
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
	// MainGoodID holds the value of the "main_good_id" field.
	MainGoodID uuid.UUID `json:"main_good_id,omitempty"`
	// SubGoodID holds the value of the "sub_good_id" field.
	SubGoodID uuid.UUID `json:"sub_good_id,omitempty"`
	// Must holds the value of the "must" field.
	Must bool `json:"must,omitempty"`
	// Commission holds the value of the "commission" field.
	Commission bool `json:"commission,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubGood) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case subgood.FieldMust, subgood.FieldCommission:
			values[i] = new(sql.NullBool)
		case subgood.FieldCreatedAt, subgood.FieldUpdatedAt, subgood.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case subgood.FieldID, subgood.FieldAppID, subgood.FieldMainGoodID, subgood.FieldSubGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SubGood", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubGood fields.
func (sg *SubGood) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subgood.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sg.ID = *value
			}
		case subgood.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sg.CreatedAt = uint32(value.Int64)
			}
		case subgood.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sg.UpdatedAt = uint32(value.Int64)
			}
		case subgood.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sg.DeletedAt = uint32(value.Int64)
			}
		case subgood.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				sg.AppID = *value
			}
		case subgood.FieldMainGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field main_good_id", values[i])
			} else if value != nil {
				sg.MainGoodID = *value
			}
		case subgood.FieldSubGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field sub_good_id", values[i])
			} else if value != nil {
				sg.SubGoodID = *value
			}
		case subgood.FieldMust:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field must", values[i])
			} else if value.Valid {
				sg.Must = value.Bool
			}
		case subgood.FieldCommission:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field commission", values[i])
			} else if value.Valid {
				sg.Commission = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SubGood.
// Note that you need to call SubGood.Unwrap() before calling this method if this SubGood
// was returned from a transaction, and the transaction was committed or rolled back.
func (sg *SubGood) Update() *SubGoodUpdateOne {
	return (&SubGoodClient{config: sg.config}).UpdateOne(sg)
}

// Unwrap unwraps the SubGood entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sg *SubGood) Unwrap() *SubGood {
	_tx, ok := sg.config.driver.(*txDriver)
	if !ok {
		panic("ent: SubGood is not a transactional entity")
	}
	sg.config.driver = _tx.drv
	return sg
}

// String implements the fmt.Stringer.
func (sg *SubGood) String() string {
	var builder strings.Builder
	builder.WriteString("SubGood(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sg.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", sg.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", sg.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", sg.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", sg.AppID))
	builder.WriteString(", ")
	builder.WriteString("main_good_id=")
	builder.WriteString(fmt.Sprintf("%v", sg.MainGoodID))
	builder.WriteString(", ")
	builder.WriteString("sub_good_id=")
	builder.WriteString(fmt.Sprintf("%v", sg.SubGoodID))
	builder.WriteString(", ")
	builder.WriteString("must=")
	builder.WriteString(fmt.Sprintf("%v", sg.Must))
	builder.WriteString(", ")
	builder.WriteString("commission=")
	builder.WriteString(fmt.Sprintf("%v", sg.Commission))
	builder.WriteByte(')')
	return builder.String()
}

// SubGoods is a parsable slice of SubGood.
type SubGoods []*SubGood

func (sg SubGoods) config(cfg config) {
	for _i := range sg {
		sg[_i].config = cfg
	}
}