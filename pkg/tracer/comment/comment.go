package comment

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
)

func trace(span trace1.Span, in *npool.AppGoodReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.Bool(fmt.Sprintf("Online.%v", index), in.GetOnline()),
		attribute.Bool(fmt.Sprintf("Visible.%v", index), in.GetVisible()),
		attribute.String(fmt.Sprintf("GoodName.%v", index), in.GetGoodName()),
		attribute.String(fmt.Sprintf("Price.%v", index), in.GetPrice()),
		attribute.Int(fmt.Sprintf("DisplayIndex.%v", index), int(in.GetDisplayIndex())),
		attribute.Int(fmt.Sprintf("PurchaseLimit.%v", index), int(in.GetPurchaseLimit())),
		attribute.Int(fmt.Sprintf("CommissionPercent.%v", index), int(in.GetCommissionPercent())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppGoodReq) trace1.Span {
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

func TraceMany(span trace1.Span, infos []*npool.AppGoodReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
