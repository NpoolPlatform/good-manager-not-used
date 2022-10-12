package vendorlocation

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/vendorlocation"
)

func validate(info *npool.VendorLocationReq) error {
	if info.GetCountry() == "" {
		logger.Sugar().Errorw("validate", "Country", info.GetCountry())
		return status.Error(codes.InvalidArgument, "Country is empty")
	}

	if info.GetProvince() == "" {
		logger.Sugar().Errorw("validate", "Province", info.GetProvince())
		return status.Error(codes.InvalidArgument, "Province is empty")
	}

	if info.GetCity() == "" {
		logger.Sugar().Errorw("validate", "City", info.GetCity())
		return status.Error(codes.InvalidArgument, "City is empty")
	}

	if info.GetAddress() == "" {
		logger.Sugar().Errorw("validate", "Address", info.GetAddress())
		return status.Error(codes.InvalidArgument, "Address is empty")
	}

	return nil
}
