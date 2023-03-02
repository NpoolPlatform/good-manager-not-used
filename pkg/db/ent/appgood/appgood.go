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
	// FieldSaleStartAt holds the string denoting the sale_start_at field in the database.
	FieldSaleStartAt = "sale_start_at"
	// FieldSaleEndAt holds the string denoting the sale_end_at field in the database.
	FieldSaleEndAt = "sale_end_at"
	// FieldServiceStartAt holds the string denoting the service_start_at field in the database.
	FieldServiceStartAt = "service_start_at"
	// FieldTechnicalFeeRatio holds the string denoting the technical_fee_ratio field in the database.
	FieldTechnicalFeeRatio = "technical_fee_ratio"
	// FieldElectricityFeeRatio holds the string denoting the electricity_fee_ratio field in the database.
	FieldElectricityFeeRatio = "electricity_fee_ratio"
	// FieldDailyRewardAmount holds the string denoting the daily_reward_amount field in the database.
	FieldDailyRewardAmount = "daily_reward_amount"
	// FieldCommissionSettleType holds the string denoting the commission_settle_type field in the database.
	FieldCommissionSettleType = "commission_settle_type"
	// FieldDescriptions holds the string denoting the descriptions field in the database.
	FieldDescriptions = "descriptions"
	// FieldGoodBanner holds the string denoting the good_banner field in the database.
	FieldGoodBanner = "good_banner"
	// FieldDisplayNames holds the string denoting the display_names field in the database.
	FieldDisplayNames = "display_names"
	// FieldEnablePurchase holds the string denoting the enable_purchase field in the database.
	FieldEnablePurchase = "enable_purchase"
	// FieldEnableProductPage holds the string denoting the enable_product_page field in the database.
	FieldEnableProductPage = "enable_product_page"
	// FieldCancelMode holds the string denoting the cancel_mode field in the database.
	FieldCancelMode = "cancel_mode"
	// FieldUserPurchaseLimit holds the string denoting the user_purchase_limit field in the database.
	FieldUserPurchaseLimit = "user_purchase_limit"
	// FieldDisplayColors holds the string denoting the display_colors field in the database.
	FieldDisplayColors = "display_colors"
	// FieldCancellableBeforeStart holds the string denoting the cancellable_before_start field in the database.
	FieldCancellableBeforeStart = "cancellable_before_start"
	// FieldProductPage holds the string denoting the product_page field in the database.
	FieldProductPage = "product_page"
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
	FieldSaleStartAt,
	FieldSaleEndAt,
	FieldServiceStartAt,
	FieldTechnicalFeeRatio,
	FieldElectricityFeeRatio,
	FieldDailyRewardAmount,
	FieldCommissionSettleType,
	FieldDescriptions,
	FieldGoodBanner,
	FieldDisplayNames,
	FieldEnablePurchase,
	FieldEnableProductPage,
	FieldCancelMode,
	FieldUserPurchaseLimit,
	FieldDisplayColors,
	FieldCancellableBeforeStart,
	FieldProductPage,
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
	// DefaultSaleStartAt holds the default value on creation for the "sale_start_at" field.
	DefaultSaleStartAt uint32
	// DefaultSaleEndAt holds the default value on creation for the "sale_end_at" field.
	DefaultSaleEndAt uint32
	// DefaultServiceStartAt holds the default value on creation for the "service_start_at" field.
	DefaultServiceStartAt uint32
	// DefaultTechnicalFeeRatio holds the default value on creation for the "technical_fee_ratio" field.
	DefaultTechnicalFeeRatio uint32
	// DefaultElectricityFeeRatio holds the default value on creation for the "electricity_fee_ratio" field.
	DefaultElectricityFeeRatio uint32
	// DefaultDailyRewardAmount holds the default value on creation for the "daily_reward_amount" field.
	DefaultDailyRewardAmount decimal.Decimal
	// DefaultCommissionSettleType holds the default value on creation for the "commission_settle_type" field.
	DefaultCommissionSettleType string
	// DefaultDescriptions holds the default value on creation for the "descriptions" field.
	DefaultDescriptions []string
	// DefaultGoodBanner holds the default value on creation for the "good_banner" field.
	DefaultGoodBanner string
	// DefaultDisplayNames holds the default value on creation for the "display_names" field.
	DefaultDisplayNames []string
	// DefaultEnablePurchase holds the default value on creation for the "enable_purchase" field.
	DefaultEnablePurchase bool
	// DefaultEnableProductPage holds the default value on creation for the "enable_product_page" field.
	DefaultEnableProductPage bool
	// DefaultCancelMode holds the default value on creation for the "cancel_mode" field.
	DefaultCancelMode string
	// DefaultDisplayColors holds the default value on creation for the "display_colors" field.
	DefaultDisplayColors []string
	// DefaultCancellableBeforeStart holds the default value on creation for the "cancellable_before_start" field.
	DefaultCancellableBeforeStart uint32
	// DefaultProductPage holds the default value on creation for the "product_page" field.
	DefaultProductPage string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
