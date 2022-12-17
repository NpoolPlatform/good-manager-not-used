package good

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Good) *npool.Good {
	if row == nil {
		return nil
	}

	supportCoinTypeIDs := []string{}
	for _, val := range row.SupportCoinTypeIds {
		supportCoinTypeIDs = append(supportCoinTypeIDs, val.String())
	}

	return &npool.Good{
		ID:                   row.ID.String(),
		DeviceInfoID:         row.DeviceInfoID.String(),
		DurationDays:         row.DurationDays,
		CoinTypeID:           row.CoinTypeID.String(),
		InheritFromGoodID:    row.InheritFromGoodID.String(),
		VendorLocationID:     row.VendorLocationID.String(),
		Price:                row.Price.String(),
		BenefitType:          npool.BenefitType(npool.BenefitType_value[row.BenefitType]),
		GoodType:             npool.GoodType(npool.GoodType_value[row.GoodType]),
		Title:                row.Title,
		Unit:                 row.Unit,
		UnitAmount:           row.UnitAmount,
		SupportCoinTypeIDs:   supportCoinTypeIDs,
		DeliveryAt:           row.DeliveryAt,
		StartAt:              row.StartAt,
		TestOnly:             row.TestOnly,
		BenefitIntervalHours: row.BenefitIntervalHours,
		CreatedAt:            row.CreatedAt,
		UpdatedAt:            row.UpdatedAt,
		DeletedAt:            row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Good) []*npool.Good {
	infos := []*npool.Good{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
