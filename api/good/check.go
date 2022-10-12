package good

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	"github.com/google/uuid"
)

//nolint
func validate(info *npool.GoodReq) error {
	if info.DeviceInfoID == nil {
		logger.Sugar().Errorw("validate", "DeviceInfoID", info.DeviceInfoID)
		return status.Error(codes.InvalidArgument, "DeviceInfoID is empty")
	}

	if _, err := uuid.Parse(info.GetDeviceInfoID()); err != nil {
		logger.Sugar().Errorw("validate", "DeviceInfoID", info.GetDeviceInfoID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("DeviceInfoID is invalid: %v", err))
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.VendorLocationID == nil {
		logger.Sugar().Errorw("validate", "VendorLocationID", info.VendorLocationID)
		return status.Error(codes.InvalidArgument, "VendorLocationID is empty")
	}

	if _, err := uuid.Parse(info.GetVendorLocationID()); err != nil {
		logger.Sugar().Errorw("validate", "VendorLocationID", info.GetVendorLocationID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("VendorLocationID is invalid: %v", err))
	}

	if info.Price == nil {
		logger.Sugar().Errorw("validate", "Price", info.Price)
		return status.Error(codes.InvalidArgument, "Price is empty")
	}

	price, err := decimal.NewFromString(info.GetPrice())
	if err != nil {
		logger.Sugar().Errorw("validate", "Price", info.GetPrice(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Price is invalid: %v", err))
	}

	if price.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "Price", info.GetPrice(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "GetPrice is Less than or equal to 0")
	}

	switch info.GetBenefitType() {
	case npool.BenefitType_BenefitTypePlatform:
	case npool.BenefitType_BenefitTypePool:
	default:
		logger.Sugar().Errorw("validate", "BenefitType", info.GetBenefitType(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("BenefitType is invalid: %v", err))
	}

	switch info.GetGoodType() {
	case npool.GoodType_GoodTypeClassicMining:
	case npool.GoodType_GoodTypeUnionMining:
	case npool.GoodType_GoodTypeTechniqueFee:
	case npool.GoodType_GoodTypeElectricityFee:
	default:
		logger.Sugar().Errorw("validate", "GoodType", info.GetGoodType(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodType is invalid: %v", err))
	}

	if info.GetTitle() == "" {
		logger.Sugar().Errorw("validate", "Title", info.GetTitle())
		return status.Error(codes.InvalidArgument, "Title is empty")
	}

	if info.GetUnit() == "" {
		logger.Sugar().Errorw("validate", "Unit", info.GetUnit())
		return status.Error(codes.InvalidArgument, "Unit is empty")
	}

	if info.GetUnitAmount() <= 0 {
		logger.Sugar().Errorw("validate", "UnitAmount", info.GetUnitAmount())
		return status.Error(codes.InvalidArgument, "GetPrice is Less than or equal to 0")
	}

	if len(info.GetSupportCoinTypeIDs()) == 0 {
		logger.Sugar().Errorw("validate", "SupportCoinTypeIDs", info.GetSupportCoinTypeIDs())
		return status.Error(codes.InvalidArgument, fmt.Sprintf("SupportCoinTypeIDs is invalid: %v", err))
	}

	if info.GetDeliveryAt() <= 0 {
		logger.Sugar().Errorw("validate", "DeliveryAt", info.GetDeliveryAt())
		return status.Error(codes.InvalidArgument, "GetPrice is Less than or equal to 0")
	}

	if info.GetStartAt() <= 0 {
		logger.Sugar().Errorw("validate", "StartAt", info.GetStartAt())
		return status.Error(codes.InvalidArgument, "GetPrice is Less than or equal to 0")
	}

	return nil
}

func duplicate(infos []*npool.GoodReq) error {
	keys := map[string]struct{}{}
	s := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v", info.Title)
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate Title")
		}

		keys[key] = struct{}{}
		s[info.GetID()] = struct{}{}
	}

	if len(s) > 1 {
		return status.Error(codes.InvalidArgument, "Infos has different ID")
	}

	return nil
}
