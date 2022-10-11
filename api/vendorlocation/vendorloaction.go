//nolint:nolintlint,dupl
package vendorlocation

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/vendorlocation"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/vendorlocation"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/vendorlocation"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/vendorlocation"

	"github.com/google/uuid"
)

func (s *Server) CreateVendorLocation(ctx context.Context, in *npool.CreateVendorLocationRequest) (*npool.CreateVendorLocationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateVendorLocation")
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
		return &npool.CreateVendorLocationResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create VendorLocation: %v", err.Error())
		return &npool.CreateVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateVendorLocationResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateVendorLocations(ctx context.Context, in *npool.CreateVendorLocationsRequest) (*npool.CreateVendorLocationsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateVendorLocations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateVendorLocationsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create VendorLocations: %v", err)
		return &npool.CreateVendorLocationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateVendorLocationsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateVendorLocation(ctx context.Context, in *npool.UpdateVendorLocationRequest) (*npool.UpdateVendorLocationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateVendorLocations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create VendorLocations: %v", err)
		return &npool.UpdateVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateVendorLocationResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetVendorLocation(ctx context.Context, in *npool.GetVendorLocationRequest) (*npool.GetVendorLocationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetVendorLocation")
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
		return &npool.GetVendorLocationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get VendorLocation: %v", err)
		return &npool.GetVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetVendorLocationResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetVendorLocationOnly(ctx context.Context, in *npool.GetVendorLocationOnlyRequest) (*npool.GetVendorLocationOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetVendorLocationOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get VendorLocations: %v", err)
		return &npool.GetVendorLocationOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetVendorLocationOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetVendorLocations(ctx context.Context, in *npool.GetVendorLocationsRequest) (*npool.GetVendorLocationsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetVendorLocations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get VendorLocations: %v", err)
		return &npool.GetVendorLocationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetVendorLocationsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistVendorLocation(ctx context.Context, in *npool.ExistVendorLocationRequest) (*npool.ExistVendorLocationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistVendorLocation")
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
		return &npool.ExistVendorLocationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check VendorLocation: %v", err)
		return &npool.ExistVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistVendorLocationResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistVendorLocationConds(ctx context.Context,
	in *npool.ExistVendorLocationCondsRequest) (*npool.ExistVendorLocationCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistVendorLocationConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check VendorLocation: %v", err)
		return &npool.ExistVendorLocationCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistVendorLocationCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountVendorLocations(ctx context.Context, in *npool.CountVendorLocationsRequest) (*npool.CountVendorLocationsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountVendorLocations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count VendorLocations: %v", err)
		return &npool.CountVendorLocationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountVendorLocationsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteVendorLocation(ctx context.Context, in *npool.DeleteVendorLocationRequest) (*npool.DeleteVendorLocationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateVendorLocations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "VendorLocation", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteVendorLocationResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create VendorLocations: %v", err)
		return &npool.DeleteVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteVendorLocationResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
