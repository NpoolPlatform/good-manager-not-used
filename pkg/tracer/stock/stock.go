package stock

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"
)

func trace(span trace1.Span, in *npool.StockReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.String(fmt.Sprintf("Total.%v", index), in.GetTotal()),
		attribute.String(fmt.Sprintf("Locked.%v", index), in.GetLocked()),
		attribute.String(fmt.Sprintf("InService.%v", index), in.GetInService()),
		attribute.String(fmt.Sprintf("Sold.%v", index), in.GetSold()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.StockReq) trace1.Span {
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

func TraceMany(span trace1.Span, infos []*npool.StockReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
