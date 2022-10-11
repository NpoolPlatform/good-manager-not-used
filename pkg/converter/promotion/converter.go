package promotion

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/promotion"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Promotion) *npool.Promotion {
	if row == nil {
		return nil
	}

	return &npool.Promotion{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		GoodID:    row.GoodID.String(),
		Message:   row.Message,
		StartAt:   row.StartAt,
		EndAt:     row.EndAt,
		Price:     row.Price.String(),
		Posters:   row.Posters,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Promotion) []*npool.Promotion {
	infos := []*npool.Promotion{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
