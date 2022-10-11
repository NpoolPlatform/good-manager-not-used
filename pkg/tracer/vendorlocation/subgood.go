package vendorlocation

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/vendorlocation"
)

func trace(span trace1.Span, in *npool.VendorLocationReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("Country.%v", index), in.GetCountry()),
		attribute.String(fmt.Sprintf("Province.%v", index), in.GetProvince()),
		attribute.String(fmt.Sprintf("City.%v", index), in.GetCity()),
		attribute.String(fmt.Sprintf("Address.%v", index), in.GetAddress()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.VendorLocationReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("Country.Op", in.GetCountry().GetOp()),
		attribute.String("Country.Value", in.GetCountry().GetValue()),
		attribute.String("Province.Op", in.GetProvince().GetOp()),
		attribute.String("Province.Value", in.GetProvince().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.VendorLocationReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
