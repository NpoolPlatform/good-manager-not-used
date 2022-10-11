//nolint:nolintlint,dupl
package comment

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/comment"
	crud "github.com/NpoolPlatform/good-manager/pkg/crud/comment"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/comment"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/comment"

	"github.com/google/uuid"
)

func (s *Server) CreateComment(ctx context.Context, in *npool.CreateCommentRequest) (*npool.CreateCommentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateComment")
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
		return &npool.CreateCommentResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "comment", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create comment: %v", err.Error())
		return &npool.CreateCommentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCommentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateComments(ctx context.Context, in *npool.CreateCommentsRequest) (*npool.CreateCommentsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateComments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCommentsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "comment", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create comments: %v", err)
		return &npool.CreateCommentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCommentsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateComment(ctx context.Context, in *npool.UpdateCommentRequest) (*npool.UpdateCommentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateComments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "comment", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateCommentResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create comments: %v", err)
		return &npool.UpdateCommentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCommentResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetComment(ctx context.Context, in *npool.GetCommentRequest) (*npool.GetCommentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetComment")
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
		return &npool.GetCommentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "comment", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get comment: %v", err)
		return &npool.GetCommentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCommentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCommentOnly(ctx context.Context, in *npool.GetCommentOnlyRequest) (*npool.GetCommentOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCommentOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "comment", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get comments: %v", err)
		return &npool.GetCommentOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCommentOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetComments(ctx context.Context, in *npool.GetCommentsRequest) (*npool.GetCommentsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetComments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "comment", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get comments: %v", err)
		return &npool.GetCommentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCommentsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistComment(ctx context.Context, in *npool.ExistCommentRequest) (*npool.ExistCommentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistComment")
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
		return &npool.ExistCommentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "comment", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check comment: %v", err)
		return &npool.ExistCommentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCommentResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCommentConds(ctx context.Context,
	in *npool.ExistCommentCondsRequest) (*npool.ExistCommentCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCommentConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "comment", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check comment: %v", err)
		return &npool.ExistCommentCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCommentCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountComments(ctx context.Context, in *npool.CountCommentsRequest) (*npool.CountCommentsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountComments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "comment", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count comments: %v", err)
		return &npool.CountCommentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCommentsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteComment(ctx context.Context, in *npool.DeleteCommentRequest) (*npool.DeleteCommentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateComments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "comment", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteCommentResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create comments: %v", err)
		return &npool.DeleteCommentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCommentResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
