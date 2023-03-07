//nolint:nolintlint,dupl
package appdefaultgood

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/appdefaultgood"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/appdefaultgood"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/appdefaultgood"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appdefaultgood"

	"github.com/google/uuid"
)

func (s *Server) CreateAppDefaultGood(
	ctx context.Context,
	in *npool.CreateAppDefaultGoodRequest,
) (
	*npool.CreateAppDefaultGoodResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppDefaultGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		return &npool.CreateAppDefaultGoodResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create appdefaultgood: %v", err.Error())
		return &npool.CreateAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppDefaultGoodResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppDefaultGoods(
	ctx context.Context,
	in *npool.CreateAppDefaultGoodsRequest,
) (
	*npool.CreateAppDefaultGoodsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppDefaultGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppDefaultGoodsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateAppDefaultGoodsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create appdefaultgoods: %v", err)
		return &npool.CreateAppDefaultGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppDefaultGoodsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppDefaultGood(
	ctx context.Context,
	in *npool.UpdateAppDefaultGoodRequest,
) (
	*npool.UpdateAppDefaultGoodResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppDefaultGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create appdefaultgoods: %v", err)
		return &npool.UpdateAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppDefaultGoodResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetAppDefaultGood(
	ctx context.Context,
	in *npool.GetAppDefaultGoodRequest,
) (
	*npool.GetAppDefaultGoodResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppDefaultGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get appdefaultgood: %v", err)
		return &npool.GetAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppDefaultGoodResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppDefaultGoodOnly(
	ctx context.Context,
	in *npool.GetAppDefaultGoodOnlyRequest,
) (
	*npool.GetAppDefaultGoodOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppDefaultGoodOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get appdefaultgoods: %v", err)
		return &npool.GetAppDefaultGoodOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppDefaultGoodOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppDefaultGoods(
	ctx context.Context,
	in *npool.GetAppDefaultGoodsRequest,
) (
	*npool.GetAppDefaultGoodsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppDefaultGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get appdefaultgoods: %v", err)
		return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppDefaultGoodsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppDefaultGood(
	ctx context.Context,
	in *npool.ExistAppDefaultGoodRequest,
) (
	*npool.ExistAppDefaultGoodResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppDefaultGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check appdefaultgood: %v", err)
		return &npool.ExistAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppDefaultGoodResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppDefaultGoodConds(
	ctx context.Context,
	in *npool.ExistAppDefaultGoodCondsRequest,
) (
	*npool.ExistAppDefaultGoodCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppDefaultGoodConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check appdefaultgood: %v", err)
		return &npool.ExistAppDefaultGoodCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppDefaultGoodCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppDefaultGoods(
	ctx context.Context,
	in *npool.CountAppDefaultGoodsRequest,
) (
	*npool.CountAppDefaultGoodsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppDefaultGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count appdefaultgoods: %v", err)
		return &npool.CountAppDefaultGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppDefaultGoodsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppDefaultGood(
	ctx context.Context,
	in *npool.DeleteAppDefaultGoodRequest,
) (
	*npool.DeleteAppDefaultGoodResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppDefaultGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "appdefaultgood", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create appdefaultgoods: %v", err)
		return &npool.DeleteAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppDefaultGoodResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
