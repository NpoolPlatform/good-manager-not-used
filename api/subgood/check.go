package subgood

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

	"github.com/google/uuid"
)

func validate(info *npool.SubGoodReq) error { //nolint
	if info.MainGoodID == nil {
		logger.Sugar().Errorw("validate", "MainGoodID", info.MainGoodID)
		return status.Error(codes.InvalidArgument, "MainGoodID is empty")
	}

	if _, err := uuid.Parse(info.GetMainGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "MainGoodID", info.GetMainGoodID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("MainGoodID is invalid: %v", err))
	}
	if info.SubGoodID == nil {
		logger.Sugar().Errorw("validate", "SubGoodID", info.SubGoodID)
		return status.Error(codes.InvalidArgument, "SubGoodID is empty")
	}

	if _, err := uuid.Parse(info.GetSubGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "SubGoodID", info.GetSubGoodID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("SubGoodID is invalid: %v", err))
	}

	return nil
}
