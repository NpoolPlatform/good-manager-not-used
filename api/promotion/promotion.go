//nolint:nolintlint,dupl
package promotion

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/promotion"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/promotion"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/promotion"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/promotion"

	"github.com/google/uuid"
)

func (s *Server) CreatePromotion(ctx context.Context, in *npool.CreatePromotionRequest) (*npool.CreatePromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePromotion")
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
		return &npool.CreatePromotionResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "Promotion", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Promotion: %v", err.Error())
		return &npool.CreatePromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePromotionResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreatePromotions(ctx context.Context, in *npool.CreatePromotionsRequest) (*npool.CreatePromotionsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePromotions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreatePromotionsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "Promotion", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create Promotions: %v", err)
		return &npool.CreatePromotionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePromotionsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdatePromotion(ctx context.Context, in *npool.UpdatePromotionRequest) (*npool.UpdatePromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePromotions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Promotion", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdatePromotionResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Promotions: %v", err)
		return &npool.UpdatePromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdatePromotionResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetPromotion(ctx context.Context, in *npool.GetPromotionRequest) (*npool.GetPromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetPromotion")
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
		return &npool.GetPromotionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Promotion", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get Promotion: %v", err)
		return &npool.GetPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPromotionResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetPromotionOnly(ctx context.Context, in *npool.GetPromotionOnlyRequest) (*npool.GetPromotionOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetPromotionOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Promotion", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get Promotions: %v", err)
		return &npool.GetPromotionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPromotionOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetPromotions(ctx context.Context, in *npool.GetPromotionsRequest) (*npool.GetPromotionsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetPromotions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "Promotion", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Promotions: %v", err)
		return &npool.GetPromotionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPromotionsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistPromotion(ctx context.Context, in *npool.ExistPromotionRequest) (*npool.ExistPromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistPromotion")
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
		return &npool.ExistPromotionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Promotion", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check Promotion: %v", err)
		return &npool.ExistPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistPromotionResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistPromotionConds(ctx context.Context,
	in *npool.ExistPromotionCondsRequest) (*npool.ExistPromotionCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistPromotionConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Promotion", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check Promotion: %v", err)
		return &npool.ExistPromotionCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistPromotionCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountPromotions(ctx context.Context, in *npool.CountPromotionsRequest) (*npool.CountPromotionsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountPromotions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Promotion", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count Promotions: %v", err)
		return &npool.CountPromotionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountPromotionsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeletePromotion(ctx context.Context, in *npool.DeletePromotionRequest) (*npool.DeletePromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePromotions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Promotion", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeletePromotionResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create Promotions: %v", err)
		return &npool.DeletePromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeletePromotionResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
