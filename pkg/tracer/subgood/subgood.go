package subgood

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"
)

func trace(span trace1.Span, in *npool.SubGoodReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("MainGoodID.%v", index), in.GetMainGoodID()),
		attribute.String(fmt.Sprintf("SubGoodID.%v", index), in.GetSubGoodID()),
		attribute.Bool(fmt.Sprintf("Must.%v", index), in.GetMust()),
		attribute.Bool(fmt.Sprintf("Commission.%v", index), in.GetCommission()),
		attribute.Float64(fmt.Sprintf("CreatedAt.%v", index), float64(in.GetCreatedAt())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.SubGoodReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("MainGoodID.Op", in.GetMainGoodID().GetOp()),
		attribute.String("MainGoodID.Value", in.GetMainGoodID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.SubGoodReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
