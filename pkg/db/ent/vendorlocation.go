// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/vendorlocation"
	"github.com/google/uuid"
)

// VendorLocation is the model entity for the VendorLocation schema.
type VendorLocation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VendorLocation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case vendorlocation.FieldCreatedAt, vendorlocation.FieldUpdatedAt, vendorlocation.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case vendorlocation.FieldCountry, vendorlocation.FieldProvince, vendorlocation.FieldCity, vendorlocation.FieldAddress:
			values[i] = new(sql.NullString)
		case vendorlocation.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type VendorLocation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VendorLocation fields.
func (vl *VendorLocation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vendorlocation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				vl.ID = *value
			}
		case vendorlocation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vl.CreatedAt = uint32(value.Int64)
			}
		case vendorlocation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vl.UpdatedAt = uint32(value.Int64)
			}
		case vendorlocation.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				vl.DeletedAt = uint32(value.Int64)
			}
		case vendorlocation.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				vl.Country = value.String
			}
		case vendorlocation.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				vl.Province = value.String
			}
		case vendorlocation.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				vl.City = value.String
			}
		case vendorlocation.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				vl.Address = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this VendorLocation.
// Note that you need to call VendorLocation.Unwrap() before calling this method if this VendorLocation
// was returned from a transaction, and the transaction was committed or rolled back.
func (vl *VendorLocation) Update() *VendorLocationUpdateOne {
	return (&VendorLocationClient{config: vl.config}).UpdateOne(vl)
}

// Unwrap unwraps the VendorLocation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vl *VendorLocation) Unwrap() *VendorLocation {
	_tx, ok := vl.config.driver.(*txDriver)
	if !ok {
		panic("ent: VendorLocation is not a transactional entity")
	}
	vl.config.driver = _tx.drv
	return vl
}

// String implements the fmt.Stringer.
func (vl *VendorLocation) String() string {
	var builder strings.Builder
	builder.WriteString("VendorLocation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vl.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", vl.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", vl.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", vl.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(vl.Country)
	builder.WriteString(", ")
	builder.WriteString("province=")
	builder.WriteString(vl.Province)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(vl.City)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(vl.Address)
	builder.WriteByte(')')
	return builder.String()
}

// VendorLocations is a parsable slice of VendorLocation.
type VendorLocations []*VendorLocation

func (vl VendorLocations) config(cfg config) {
	for _i := range vl {
		vl[_i].config = cfg
	}
}
