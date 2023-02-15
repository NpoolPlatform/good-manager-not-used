package stock

import (
	"fmt"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"

	"github.com/google/uuid"
)

func validate(info *npool.StockReq) error {
	if info.GoodID == nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID)
		return status.Error(codes.InvalidArgument, "GoodID is empty")
	}

	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GetGoodID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	total, err := decimal.NewFromString(info.GetTotal())
	if err != nil {
		logger.Sugar().Errorw("validate", "Total", info.GetTotal())
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if total.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "Total", info.GetTotal())
		return status.Error(codes.InvalidArgument, "Total is Less than or equal to 0")
	}
	return nil
}
