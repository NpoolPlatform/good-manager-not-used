package appdefaultgood

import (
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appdefaultgood"
)

func Ent2Grpc(row *ent.AppDefaultGood) *npool.AppDefaultGood {
	if row == nil {
		return nil
	}

	return &npool.AppDefaultGood{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		GoodID:     row.GoodID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.AppDefaultGood) []*npool.AppDefaultGood {
	infos := []*npool.AppDefaultGood{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
