// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/deviceinfo"
	"github.com/google/uuid"
)

// DeviceInfo is the model entity for the DeviceInfo schema.
type DeviceInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Manufacturer holds the value of the "manufacturer" field.
	Manufacturer string `json:"manufacturer,omitempty"`
	// PowerComsuption holds the value of the "power_comsuption" field.
	PowerComsuption uint32 `json:"power_comsuption,omitempty"`
	// ShipmentAt holds the value of the "shipment_at" field.
	ShipmentAt uint32 `json:"shipment_at,omitempty"`
	// Posters holds the value of the "posters" field.
	Posters []string `json:"posters,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeviceInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deviceinfo.FieldPosters:
			values[i] = new([]byte)
		case deviceinfo.FieldCreatedAt, deviceinfo.FieldUpdatedAt, deviceinfo.FieldDeletedAt, deviceinfo.FieldPowerComsuption, deviceinfo.FieldShipmentAt:
			values[i] = new(sql.NullInt64)
		case deviceinfo.FieldType, deviceinfo.FieldManufacturer:
			values[i] = new(sql.NullString)
		case deviceinfo.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeviceInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeviceInfo fields.
func (di *DeviceInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deviceinfo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				di.ID = *value
			}
		case deviceinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				di.CreatedAt = uint32(value.Int64)
			}
		case deviceinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				di.UpdatedAt = uint32(value.Int64)
			}
		case deviceinfo.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				di.DeletedAt = uint32(value.Int64)
			}
		case deviceinfo.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				di.Type = value.String
			}
		case deviceinfo.FieldManufacturer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer", values[i])
			} else if value.Valid {
				di.Manufacturer = value.String
			}
		case deviceinfo.FieldPowerComsuption:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field power_comsuption", values[i])
			} else if value.Valid {
				di.PowerComsuption = uint32(value.Int64)
			}
		case deviceinfo.FieldShipmentAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shipment_at", values[i])
			} else if value.Valid {
				di.ShipmentAt = uint32(value.Int64)
			}
		case deviceinfo.FieldPosters:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field posters", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &di.Posters); err != nil {
					return fmt.Errorf("unmarshal field posters: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DeviceInfo.
// Note that you need to call DeviceInfo.Unwrap() before calling this method if this DeviceInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (di *DeviceInfo) Update() *DeviceInfoUpdateOne {
	return (&DeviceInfoClient{config: di.config}).UpdateOne(di)
}

// Unwrap unwraps the DeviceInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (di *DeviceInfo) Unwrap() *DeviceInfo {
	_tx, ok := di.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeviceInfo is not a transactional entity")
	}
	di.config.driver = _tx.drv
	return di
}

// String implements the fmt.Stringer.
func (di *DeviceInfo) String() string {
	var builder strings.Builder
	builder.WriteString("DeviceInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", di.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", di.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", di.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", di.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(di.Type)
	builder.WriteString(", ")
	builder.WriteString("manufacturer=")
	builder.WriteString(di.Manufacturer)
	builder.WriteString(", ")
	builder.WriteString("power_comsuption=")
	builder.WriteString(fmt.Sprintf("%v", di.PowerComsuption))
	builder.WriteString(", ")
	builder.WriteString("shipment_at=")
	builder.WriteString(fmt.Sprintf("%v", di.ShipmentAt))
	builder.WriteString(", ")
	builder.WriteString("posters=")
	builder.WriteString(fmt.Sprintf("%v", di.Posters))
	builder.WriteByte(')')
	return builder.String()
}

// DeviceInfos is a parsable slice of DeviceInfo.
type DeviceInfos []*DeviceInfo

func (di DeviceInfos) config(cfg config) {
	for _i := range di {
		di[_i].config = cfg
	}
}
