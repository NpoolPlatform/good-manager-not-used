//nolint:nolintlint,dupl
package stock

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/stock"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/stock"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/stock"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"

	"github.com/google/uuid"
)

func (s *Server) CreateStock(ctx context.Context, in *npool.CreateStockRequest) (*npool.CreateStockResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateStock")
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
		return &npool.CreateStockResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "Stock", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Stock: %v", err.Error())
		return &npool.CreateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStockResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateStocks(ctx context.Context, in *npool.CreateStocksRequest) (*npool.CreateStocksResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateStocks")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateStocksResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "Stock", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create Stocks: %v", err)
		return &npool.CreateStocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStocksResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateStock(ctx context.Context, in *npool.UpdateStockRequest) (*npool.UpdateStockResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateStocks")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Stock", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateStockResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Stocks: %v", err)
		return &npool.UpdateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateStockResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetStock(ctx context.Context, in *npool.GetStockRequest) (*npool.GetStockResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetStock")
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
		return &npool.GetStockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Stock", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get Stock: %v", err)
		return &npool.GetStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStockResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetStockOnly(ctx context.Context, in *npool.GetStockOnlyRequest) (*npool.GetStockOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetStockOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Stock", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get Stocks: %v", err)
		return &npool.GetStockOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStockOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetStocks(ctx context.Context, in *npool.GetStocksRequest) (*npool.GetStocksResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetStocks")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "Stock", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Stocks: %v", err)
		return &npool.GetStocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStocksResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistStock(ctx context.Context, in *npool.ExistStockRequest) (*npool.ExistStockResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistStock")
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
		return &npool.ExistStockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Stock", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check Stock: %v", err)
		return &npool.ExistStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistStockResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistStockConds(ctx context.Context,
	in *npool.ExistStockCondsRequest) (*npool.ExistStockCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistStockConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Stock", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check Stock: %v", err)
		return &npool.ExistStockCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistStockCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountStocks(ctx context.Context, in *npool.CountStocksRequest) (*npool.CountStocksResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountStocks")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Stock", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count Stocks: %v", err)
		return &npool.CountStocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountStocksResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteStock(ctx context.Context, in *npool.DeleteStockRequest) (*npool.DeleteStockResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateStocks")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Stock", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteStockResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create Stocks: %v", err)
		return &npool.DeleteStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteStockResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
