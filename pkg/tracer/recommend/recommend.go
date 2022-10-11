package recommend

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/recommend"
)

func trace(span trace1.Span, in *npool.RecommendReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.String(fmt.Sprintf("RecommenderID.%v", index), in.GetRecommenderID()),
		attribute.String(fmt.Sprintf("Message.%v", index), in.GetMessage()),
		attribute.Float64(fmt.Sprintf("RecommendIndex.%v", index), float64(in.GetRecommendIndex())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.RecommendReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("GoodID.Op", in.GetGoodID().GetOp()),
		attribute.String("GoodID.Value", in.GetGoodID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.RecommendReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
