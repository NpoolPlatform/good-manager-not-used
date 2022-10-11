package comment

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/comment"

	"github.com/google/uuid"
)

func validate(info *npool.CommentReq) error { //nolint
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

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	if info.OrderID == nil {
		logger.Sugar().Errorw("validate", "OrderID", info.OrderID)
		return status.Error(codes.InvalidArgument, "OrderID is empty")
	}

	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		logger.Sugar().Errorw("validate", "OrderID", info.GetOrderID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("OrderID is invalid: %v", err))
	}

	if info.GetContent() == "" {
		logger.Sugar().Errorw("validate", "Content", info.GetContent())
		return status.Error(codes.InvalidArgument, "Content is empty")
	}

	if info.ReplyToID == nil {
		logger.Sugar().Errorw("validate", "ReplyToID", info.ReplyToID)
		return status.Error(codes.InvalidArgument, "ReplyToID is empty")
	}

	if _, err := uuid.Parse(info.GetReplyToID()); err != nil {
		logger.Sugar().Errorw("validate", "ReplyToID", info.GetReplyToID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("ReplyToID is invalid: %v", err))
	}

	return nil
}
