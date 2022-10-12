package appgood

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	"github.com/google/uuid"
)

func validate(info *npool.AppGoodReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.GoodID == nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID)
		return status.Error(codes.InvalidArgument, "GoodID is empty")
	}

	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GetGoodID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if info.GetGoodName() == "" {
		logger.Sugar().Errorw("validate", "GoodName", info.GetGoodName())
		return status.Error(codes.InvalidArgument, "GoodName is empty")
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

	return nil
}

func duplicate(infos []*npool.AppGoodReq) error {
	keys := map[string]struct{}{}
	apps := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v:%v", info.AppID, info.GoodID, info.GoodName)
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AppID:UserID:CoinTypeID")
		}

		keys[key] = struct{}{}
		apps[info.GetAppID()] = struct{}{}
	}

	if len(apps) > 1 {
		return status.Error(codes.InvalidArgument, "Infos has different AppID")
	}

	return nil
}
