//nolint:nolintlint,dupl
package deviceinfo

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/deviceinfo"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/deviceinfo"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/deviceinfo"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"

	"github.com/google/uuid"
)

func (s *Server) CreateDeviceInfo(ctx context.Context, in *npool.CreateDeviceInfoRequest) (*npool.CreateDeviceInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDeviceInfo")
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
		return &npool.CreateDeviceInfoResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create deviceinfo: %v", err.Error())
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDeviceInfoResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateDeviceInfos(ctx context.Context, in *npool.CreateDeviceInfosRequest) (*npool.CreateDeviceInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDeviceInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateDeviceInfosResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create deviceinfos: %v", err)
		return &npool.CreateDeviceInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDeviceInfosResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateDeviceInfo(ctx context.Context, in *npool.UpdateDeviceInfoRequest) (*npool.UpdateDeviceInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDeviceInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create deviceinfos: %v", err)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateDeviceInfoResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetDeviceInfo(ctx context.Context, in *npool.GetDeviceInfoRequest) (*npool.GetDeviceInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDeviceInfo")
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
		return &npool.GetDeviceInfoResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get deviceinfo: %v", err)
		return &npool.GetDeviceInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDeviceInfoResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetDeviceInfoOnly(ctx context.Context, in *npool.GetDeviceInfoOnlyRequest) (*npool.GetDeviceInfoOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDeviceInfoOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get deviceinfos: %v", err)
		return &npool.GetDeviceInfoOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDeviceInfoOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetDeviceInfos(ctx context.Context, in *npool.GetDeviceInfosRequest) (*npool.GetDeviceInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDeviceInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get deviceinfos: %v", err)
		return &npool.GetDeviceInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDeviceInfosResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistDeviceInfo(ctx context.Context, in *npool.ExistDeviceInfoRequest) (*npool.ExistDeviceInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistDeviceInfo")
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
		return &npool.ExistDeviceInfoResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check deviceinfo: %v", err)
		return &npool.ExistDeviceInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDeviceInfoResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistDeviceInfoConds(ctx context.Context,
	in *npool.ExistDeviceInfoCondsRequest) (*npool.ExistDeviceInfoCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistDeviceInfoConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check deviceinfo: %v", err)
		return &npool.ExistDeviceInfoCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDeviceInfoCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountDeviceInfos(ctx context.Context, in *npool.CountDeviceInfosRequest) (*npool.CountDeviceInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountDeviceInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count deviceinfos: %v", err)
		return &npool.CountDeviceInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountDeviceInfosResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteDeviceInfo(ctx context.Context, in *npool.DeleteDeviceInfoRequest) (*npool.DeleteDeviceInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDeviceInfos")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "deviceinfo", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteDeviceInfoResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create deviceinfos: %v", err)
		return &npool.DeleteDeviceInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteDeviceInfoResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
