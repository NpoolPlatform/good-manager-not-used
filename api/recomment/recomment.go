//nolint:nolintlint,dupl
package recommend

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/recommend"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/recommend"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/recommend"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/recommend"

	"github.com/google/uuid"
)

func (s *Server) CreateRecommend(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.CreateRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateRecommend")
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
		return &npool.CreateRecommendResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "Recommend", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Recommend: %v", err.Error())
		return &npool.CreateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRecommendResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateRecommends(ctx context.Context, in *npool.CreateRecommendsRequest) (*npool.CreateRecommendsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateRecommends")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateRecommendsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "Recommend", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create Recommends: %v", err)
		return &npool.CreateRecommendsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRecommendsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateRecommend(ctx context.Context, in *npool.UpdateRecommendRequest) (*npool.UpdateRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateRecommends")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Recommend", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateRecommendResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Recommends: %v", err)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateRecommendResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetRecommend(ctx context.Context, in *npool.GetRecommendRequest) (*npool.GetRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetRecommend")
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
		return &npool.GetRecommendResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Recommend", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get Recommend: %v", err)
		return &npool.GetRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRecommendResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetRecommendOnly(ctx context.Context, in *npool.GetRecommendOnlyRequest) (*npool.GetRecommendOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetRecommendOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Recommend", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get Recommends: %v", err)
		return &npool.GetRecommendOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRecommendOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetRecommends(ctx context.Context, in *npool.GetRecommendsRequest) (*npool.GetRecommendsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetRecommends")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "Recommend", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Recommends: %v", err)
		return &npool.GetRecommendsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRecommendsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistRecommend(ctx context.Context, in *npool.ExistRecommendRequest) (*npool.ExistRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistRecommend")
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
		return &npool.ExistRecommendResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Recommend", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check Recommend: %v", err)
		return &npool.ExistRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRecommendResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistRecommendConds(ctx context.Context,
	in *npool.ExistRecommendCondsRequest) (*npool.ExistRecommendCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistRecommendConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Recommend", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check Recommend: %v", err)
		return &npool.ExistRecommendCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRecommendCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountRecommends(ctx context.Context, in *npool.CountRecommendsRequest) (*npool.CountRecommendsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountRecommends")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Recommend", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count Recommends: %v", err)
		return &npool.CountRecommendsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountRecommendsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteRecommend(ctx context.Context, in *npool.DeleteRecommendRequest) (*npool.DeleteRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateRecommends")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Recommend", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteRecommendResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create Recommends: %v", err)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteRecommendResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
