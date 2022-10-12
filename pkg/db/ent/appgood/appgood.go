// Code generated by ent, DO NOT EDIT.

package appgood

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the appgood type in the database.
	Label = "app_good"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldOnline holds the string denoting the online field in the database.
	FieldOnline = "online"
	// FieldVisible holds the string denoting the visible field in the database.
	FieldVisible = "visible"
	// FieldGoodName holds the string denoting the good_name field in the database.
	FieldGoodName = "good_name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldDisplayIndex holds the string denoting the display_index field in the database.
	FieldDisplayIndex = "display_index"
	// FieldPurchaseLimit holds the string denoting the purchase_limit field in the database.
	FieldPurchaseLimit = "purchase_limit"
	// FieldCommissionPercent holds the string denoting the commission_percent field in the database.
	FieldCommissionPercent = "commission_percent"
	// Table holds the table name of the appgood in the database.
	Table = "app_goods"
)

// Columns holds all SQL columns for appgood fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldGoodID,
	FieldOnline,
	FieldVisible,
	FieldGoodName,
	FieldPrice,
	FieldDisplayIndex,
	FieldPurchaseLimit,
	FieldCommissionPercent,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/good-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultOnline holds the default value on creation for the "online" field.
	DefaultOnline bool
	// DefaultVisible holds the default value on creation for the "visible" field.
	DefaultVisible bool
	// DefaultGoodName holds the default value on creation for the "good_name" field.
	DefaultGoodName string
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice decimal.Decimal
	// DefaultDisplayIndex holds the default value on creation for the "display_index" field.
	DefaultDisplayIndex int32
	// DefaultPurchaseLimit holds the default value on creation for the "purchase_limit" field.
	DefaultPurchaseLimit int32
	// DefaultCommissionPercent holds the default value on creation for the "commission_percent" field.
	DefaultCommissionPercent int32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)