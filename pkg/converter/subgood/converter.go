package subgood

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.SubGood) *npool.SubGood {
	if row == nil {
		return nil
	}

	return &npool.SubGood{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		MainGoodID: row.MainGoodID.String(),
		SubGoodID:  row.SubGoodID.String(),
		Must:       row.Must,
		Commission: row.Commission,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
		DeletedAt:  row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.SubGood) []*npool.SubGood {
	infos := []*npool.SubGood{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
