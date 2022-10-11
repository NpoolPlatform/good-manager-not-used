package recommend

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/recommend"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Recommend) *npool.Recommend {
	if row == nil {
		return nil
	}

	return &npool.Recommend{
		ID:             row.ID.String(),
		AppID:          row.AppID.String(),
		GoodID:         row.GoodID.String(),
		RecommenderID:  row.RecommenderID.String(),
		Message:        row.Message,
		RecommendIndex: float32(row.RecommendIndex),
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
		DeletedAt:      row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Recommend) []*npool.Recommend {
	infos := []*npool.Recommend{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
