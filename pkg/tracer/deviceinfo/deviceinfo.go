package deviceinfo

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"
)

func trace(span trace1.Span, in *npool.DeviceInfoReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("Type.%v", index), in.GetType()),
		attribute.String(fmt.Sprintf("Manufacturer.%v", index), in.GetManufacturer()),
		attribute.Int(fmt.Sprintf("PowerComsuption.%v", index), int(in.GetPowerComsuption())),
		attribute.Int(fmt.Sprintf("ShipmentAt.%v", index), int(in.GetShipmentAt())),
		attribute.StringSlice(fmt.Sprintf("Posters.%v", index), in.GetPosters()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.DeviceInfoReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("Type.Op", in.GetType().GetOp()),
		attribute.String("Type.Value", in.GetType().GetValue()),
		attribute.String("Manufacturer.Op", in.GetManufacturer().GetOp()),
		attribute.String("Manufacturer.Value", in.GetManufacturer().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.DeviceInfoReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
