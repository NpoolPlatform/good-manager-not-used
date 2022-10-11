package promotion

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/promotion"
)

func trace(span trace1.Span, in *npool.PromotionReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.String(fmt.Sprintf("Message.%v", index), in.GetMessage()),
		attribute.Int(fmt.Sprintf("StartAt.%v", index), int(in.GetStartAt())),
		attribute.Int(fmt.Sprintf("EndAt.%v", index), int(in.GetEndAt())),
		attribute.String(fmt.Sprintf("Price.%v", index), in.GetPrice()),
		attribute.StringSlice(fmt.Sprintf("Posters.%v", index), in.GetPosters()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.PromotionReq) trace1.Span {
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

func TraceMany(span trace1.Span, infos []*npool.PromotionReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
