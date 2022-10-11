//nolint:nolintlint,dupl
package subgood

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/subgood"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/subgood"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/subgood"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

	"github.com/google/uuid"
)

func (s *Server) CreateSubGood(ctx context.Context, in *npool.CreateSubGoodRequest) (*npool.CreateSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSubGood")
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
		return &npool.CreateSubGoodResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "SubGood", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create SubGood: %v", err.Error())
		return &npool.CreateSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSubGoodResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateSubGoods(ctx context.Context, in *npool.CreateSubGoodsRequest) (*npool.CreateSubGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSubGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateSubGoodsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "SubGood", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create SubGoods: %v", err)
		return &npool.CreateSubGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSubGoodsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateSubGood(ctx context.Context, in *npool.UpdateSubGoodRequest) (*npool.UpdateSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSubGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "SubGood", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateSubGoodResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create SubGoods: %v", err)
		return &npool.UpdateSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSubGoodResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetSubGood(ctx context.Context, in *npool.GetSubGoodRequest) (*npool.GetSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSubGood")
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
		return &npool.GetSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "SubGood", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get SubGood: %v", err)
		return &npool.GetSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubGoodResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSubGoodOnly(ctx context.Context, in *npool.GetSubGoodOnlyRequest) (*npool.GetSubGoodOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSubGoodOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "SubGood", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get SubGoods: %v", err)
		return &npool.GetSubGoodOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubGoodOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSubGoods(ctx context.Context, in *npool.GetSubGoodsRequest) (*npool.GetSubGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSubGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "SubGood", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get SubGoods: %v", err)
		return &npool.GetSubGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubGoodsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistSubGood(ctx context.Context, in *npool.ExistSubGoodRequest) (*npool.ExistSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSubGood")
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
		return &npool.ExistSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "SubGood", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check SubGood: %v", err)
		return &npool.ExistSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSubGoodResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSubGoodConds(ctx context.Context,
	in *npool.ExistSubGoodCondsRequest) (*npool.ExistSubGoodCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSubGoodConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "SubGood", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check SubGood: %v", err)
		return &npool.ExistSubGoodCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSubGoodCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountSubGoods(ctx context.Context, in *npool.CountSubGoodsRequest) (*npool.CountSubGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountSubGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "SubGood", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count SubGoods: %v", err)
		return &npool.CountSubGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountSubGoodsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteSubGood(ctx context.Context, in *npool.DeleteSubGoodRequest) (*npool.DeleteSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSubGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "SubGood", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteSubGoodResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create SubGoods: %v", err)
		return &npool.DeleteSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSubGoodResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
