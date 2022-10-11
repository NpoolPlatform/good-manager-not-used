//nolint:nolintlint,dupl
package appgood

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/appgood"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/appgood"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/appgood"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	"github.com/google/uuid"
)

func (s *Server) CreateAppGood(ctx context.Context, in *npool.CreateAppGoodRequest) (*npool.CreateAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppGood")
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
		return &npool.CreateAppGoodResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "appgood", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create appgood: %v", err.Error())
		return &npool.CreateAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppGoodResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppGoods(ctx context.Context, in *npool.CreateAppGoodsRequest) (*npool.CreateAppGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppGoodsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateAppGoodsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appgood", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create appgoods: %v", err)
		return &npool.CreateAppGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppGoodsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppGood(ctx context.Context, in *npool.UpdateAppGoodRequest) (*npool.UpdateAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "appgood", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create appgoods: %v", err)
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppGoodResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetAppGood(ctx context.Context, in *npool.GetAppGoodRequest) (*npool.GetAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppGood")
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
		return &npool.GetAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appgood", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get appgood: %v", err)
		return &npool.GetAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppGoodOnly(ctx context.Context, in *npool.GetAppGoodOnlyRequest) (*npool.GetAppGoodOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppGoodOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appgood", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get appgoods: %v", err)
		return &npool.GetAppGoodOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppGoods(ctx context.Context, in *npool.GetAppGoodsRequest) (*npool.GetAppGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appgood", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get appgoods: %v", err)
		return &npool.GetAppGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppGood(ctx context.Context, in *npool.ExistAppGoodRequest) (*npool.ExistAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppGood")
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
		return &npool.ExistAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appgood", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check appgood: %v", err)
		return &npool.ExistAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppGoodResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppGoodConds(ctx context.Context,
	in *npool.ExistAppGoodCondsRequest) (*npool.ExistAppGoodCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppGoodConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appgood", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check appgood: %v", err)
		return &npool.ExistAppGoodCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppGoodCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppGoods(ctx context.Context, in *npool.CountAppGoodsRequest) (*npool.CountAppGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appgood", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count appgoods: %v", err)
		return &npool.CountAppGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppGoodsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppGood(ctx context.Context, in *npool.DeleteAppGoodRequest) (*npool.DeleteAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "appgood", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteAppGoodResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create appgoods: %v", err)
		return &npool.DeleteAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppGoodResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
