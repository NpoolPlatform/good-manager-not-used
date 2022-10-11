package appgood

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.AppGood) *npool.AppGood {
	if row == nil {
		return nil
	}

	return &npool.AppGood{
		ID:                row.ID.String(),
		AppID:             row.AppID.String(),
		GoodID:            row.GoodID.String(),
		Online:            row.Online,
		Visible:           row.Visible,
		GoodName:          row.GoodName,
		Price:             row.Price.String(),
		DisplayIndex:      row.DisplayIndex,
		PurchaseLimit:     row.PurchaseLimit,
		CommissionPercent: row.CommissionPercent,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
		DeletedAt:         row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.AppGood) []*npool.AppGood {
	infos := []*npool.AppGood{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
