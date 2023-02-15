package stock

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.StockV1) *npool.Stock {
	if row == nil {
		return nil
	}

	return &npool.Stock{
		ID:        row.ID.String(),
		GoodID:    row.GoodID.String(),
		Total:     row.Total.String(),
		Locked:    row.Locked.String(),
		InService: row.InService.String(),
		WaitStart: row.WaitStart.String(),
		Sold:      row.Sold.String(),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.StockV1) []*npool.Stock {
	infos := []*npool.Stock{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
