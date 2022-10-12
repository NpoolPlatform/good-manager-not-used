package recommend

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/recommend"

	"github.com/google/uuid"
)

func validate(info *npool.RecommendReq) error {
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

	if info.RecommenderID == nil {
		logger.Sugar().Errorw("validate", "RecommenderID", info.RecommenderID)
		return status.Error(codes.InvalidArgument, "RecommenderID is empty")
	}

	if _, err := uuid.Parse(info.GetRecommenderID()); err != nil {
		logger.Sugar().Errorw("validate", "RecommenderID", info.GetRecommenderID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("RecommenderID is invalid: %v", err))
	}

	if info.Message == nil {
		logger.Sugar().Errorw("validate", "Message", info.Message)
		return status.Error(codes.InvalidArgument, "Message is empty")
	}
	return nil
}
