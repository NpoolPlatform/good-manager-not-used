// Code generated by ent, DO NOT EDIT.

package good

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the good type in the database.
	Label = "good"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeviceInfoID holds the string denoting the device_info_id field in the database.
	FieldDeviceInfoID = "device_info_id"
	// FieldDurationDays holds the string denoting the duration_days field in the database.
	FieldDurationDays = "duration_days"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldInheritFromGoodID holds the string denoting the inherit_from_good_id field in the database.
	FieldInheritFromGoodID = "inherit_from_good_id"
	// FieldVendorLocationID holds the string denoting the vendor_location_id field in the database.
	FieldVendorLocationID = "vendor_location_id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldBenefitType holds the string denoting the benefit_type field in the database.
	FieldBenefitType = "benefit_type"
	// FieldGoodType holds the string denoting the good_type field in the database.
	FieldGoodType = "good_type"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// FieldUnitAmount holds the string denoting the unit_amount field in the database.
	FieldUnitAmount = "unit_amount"
	// FieldSupportCoinTypeIds holds the string denoting the support_coin_type_ids field in the database.
	FieldSupportCoinTypeIds = "support_coin_type_ids"
	// FieldDeliveryAt holds the string denoting the delivery_at field in the database.
	FieldDeliveryAt = "delivery_at"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldTestOnly holds the string denoting the test_only field in the database.
	FieldTestOnly = "test_only"
	// Table holds the table name of the good in the database.
	Table = "goods"
)

// Columns holds all SQL columns for good fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldDeviceInfoID,
	FieldDurationDays,
	FieldCoinTypeID,
	FieldInheritFromGoodID,
	FieldVendorLocationID,
	FieldPrice,
	FieldBenefitType,
	FieldGoodType,
	FieldTitle,
	FieldUnit,
	FieldUnitAmount,
	FieldSupportCoinTypeIds,
	FieldDeliveryAt,
	FieldStartAt,
	FieldTestOnly,
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
	// DefaultDurationDays holds the default value on creation for the "duration_days" field.
	DefaultDurationDays int32
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice decimal.Decimal
	// DefaultBenefitType holds the default value on creation for the "benefit_type" field.
	DefaultBenefitType string
	// DefaultGoodType holds the default value on creation for the "good_type" field.
	DefaultGoodType string
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultUnit holds the default value on creation for the "unit" field.
	DefaultUnit string
	// DefaultUnitAmount holds the default value on creation for the "unit_amount" field.
	DefaultUnitAmount int32
	// DefaultSupportCoinTypeIds holds the default value on creation for the "support_coin_type_ids" field.
	DefaultSupportCoinTypeIds []uuid.UUID
	// DefaultDeliveryAt holds the default value on creation for the "delivery_at" field.
	DefaultDeliveryAt uint32
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt uint32
	// DefaultTestOnly holds the default value on creation for the "test_only" field.
	DefaultTestOnly bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
