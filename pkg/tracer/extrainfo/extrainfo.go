package extrainfo

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/extrainfo"
)

func trace(span trace1.Span, in *npool.ExtraInfoReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.StringSlice(fmt.Sprintf("Posters.%v", index), in.GetPosters()),
		attribute.StringSlice(fmt.Sprintf("Labels.%v", index), in.GetLabels()),
		attribute.Int(fmt.Sprintf("VoteCount.%v", index), int(in.GetVoteCount())),
		attribute.Float64(fmt.Sprintf("Rating.%v", index), float64(in.GetRating())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.ExtraInfoReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("GoodID.Op", in.GetGoodID().GetOp()),
		attribute.String("GoodID.Value", in.GetGoodID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.ExtraInfoReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
