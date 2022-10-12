package deviceinfo

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"
)

func validate(info *npool.DeviceInfoReq) error {
	if info.GetType() == "" {
		logger.Sugar().Errorw("validate", "Type", info.GetType())
		return status.Error(codes.InvalidArgument, "Type is empty")
	}

	if info.GetManufacturer() == "" {
		logger.Sugar().Errorw("validate", "Manufacturer", info.GetManufacturer())
		return status.Error(codes.InvalidArgument, "Manufacturer is empty")
	}

	if info.Manufacturer == nil {
		logger.Sugar().Errorw("validate", "Manufacturer", info.Manufacturer)
		return status.Error(codes.InvalidArgument, "Manufacturer is empty")
	}

	if info.ShipmentAt == nil {
		logger.Sugar().Errorw("validate", "ShipmentAt", info.ShipmentAt)
		return status.Error(codes.InvalidArgument, "ShipmentAt is empty")
	}

	return nil
}
