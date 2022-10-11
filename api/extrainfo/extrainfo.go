//nolint:nolintlint,dupl
package extrainfo

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/extrainfo"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/extrainfo"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/extrainfo"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/extrainfo"

	"github.com/google/uuid"
)

func (s *Server) CreateExtraInfo(ctx context.Context, in *npool.CreateExtraInfoRequest) (*npool.CreateExtraInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateExtraInfo")
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
		return &npool.CreateExtraInfoResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create extrainfo: %v", err.Error())
		return &npool.CreateExtraInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateExtraInfoResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateExtraInfos(ctx context.Context, in *npool.CreateExtraInfosRequest) (*npool.CreateExtraInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateExtraInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateExtraInfosResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create extrainfos: %v", err)
		return &npool.CreateExtraInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateExtraInfosResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateExtraInfo(ctx context.Context, in *npool.UpdateExtraInfoRequest) (*npool.UpdateExtraInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateExtraInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateExtraInfoResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create extrainfos: %v", err)
		return &npool.UpdateExtraInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateExtraInfoResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetExtraInfo(ctx context.Context, in *npool.GetExtraInfoRequest) (*npool.GetExtraInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetExtraInfo")
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
		return &npool.GetExtraInfoResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get extrainfo: %v", err)
		return &npool.GetExtraInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetExtraInfoResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetExtraInfoOnly(ctx context.Context, in *npool.GetExtraInfoOnlyRequest) (*npool.GetExtraInfoOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetExtraInfoOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get extrainfos: %v", err)
		return &npool.GetExtraInfoOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetExtraInfoOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetExtraInfos(ctx context.Context, in *npool.GetExtraInfosRequest) (*npool.GetExtraInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetExtraInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get extrainfos: %v", err)
		return &npool.GetExtraInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetExtraInfosResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistExtraInfo(ctx context.Context, in *npool.ExistExtraInfoRequest) (*npool.ExistExtraInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistExtraInfo")
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
		return &npool.ExistExtraInfoResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check extrainfo: %v", err)
		return &npool.ExistExtraInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistExtraInfoResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistExtraInfoConds(ctx context.Context,
	in *npool.ExistExtraInfoCondsRequest) (*npool.ExistExtraInfoCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistExtraInfoConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check extrainfo: %v", err)
		return &npool.ExistExtraInfoCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistExtraInfoCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountExtraInfos(ctx context.Context, in *npool.CountExtraInfosRequest) (*npool.CountExtraInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountExtraInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count extrainfos: %v", err)
		return &npool.CountExtraInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountExtraInfosResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteExtraInfo(ctx context.Context, in *npool.DeleteExtraInfoRequest) (*npool.DeleteExtraInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateExtraInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "extrainfo", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteExtraInfoResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create extrainfos: %v", err)
		return &npool.DeleteExtraInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteExtraInfoResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
