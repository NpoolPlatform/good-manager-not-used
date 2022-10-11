package good

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
)

func trace(span trace1.Span, in *npool.GoodReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("DeviceInfoID.%v", index), in.GetDeviceInfoID()),
		attribute.Int(fmt.Sprintf("DurationDays.%v", index), int(in.GetDurationDays())),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("InheritFromGoodID.%v", index), in.GetInheritFromGoodID()),
		attribute.String(fmt.Sprintf("VendorLocationID.%v", index), in.GetVendorLocationID()),
		attribute.String(fmt.Sprintf("Price.%v", index), in.GetPrice()),
		attribute.String(fmt.Sprintf("BenefitType.%v", index), in.GetBenefitType().String()),
		attribute.String(fmt.Sprintf("GoodType.%v", index), in.GetGoodType().String()),
		attribute.String(fmt.Sprintf("Title.%v", index), in.GetTitle()),
		attribute.String(fmt.Sprintf("Unit.%v", index), in.GetUnit()),
		attribute.Int(fmt.Sprintf("UnitAmount.%v", index), int(in.GetUnitAmount())),
		attribute.StringSlice(fmt.Sprintf("SupportCoinTypeIDs.%v", index), in.GetSupportCoinTypeIDs()),
		attribute.Int(fmt.Sprintf("DeliveryAt.%v", index), int(in.GetDeliveryAt())),
		attribute.Int(fmt.Sprintf("StartAt.%v", index), int(in.GetStartAt())),
		attribute.Bool(fmt.Sprintf("TestOnly.%v", index), in.GetTestOnly()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.GoodReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("DeviceInfoID.Op", in.GetDeviceInfoID().GetOp()),
		attribute.String("DeviceInfoID.Value", in.GetDeviceInfoID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("VendorLocationID.Op", in.GetVendorLocationID().GetOp()),
		attribute.String("VendorLocationID.Value", in.GetVendorLocationID().GetValue()),
		attribute.String("BenefitType.Op", in.GetBenefitType().GetOp()),
		attribute.Int("BenefitType.Value", int(in.GetBenefitType().GetValue())),
		attribute.String("GoodType.Op", in.GetGoodType().GetOp()),
		attribute.Int("GoodType.Value", int(in.GetGoodType().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.GoodReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
