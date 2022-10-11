package extrainfo

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/extrainfo"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.ExtraInfo) *npool.ExtraInfo {
	if row == nil {
		return nil
	}

	return &npool.ExtraInfo{
		ID:        row.ID.String(),
		GoodID:    row.GoodID.String(),
		Posters:   row.Posters,
		Labels:    row.Labels,
		VoteCount: row.VoteCount,
		Rating:    row.Rating,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.ExtraInfo) []*npool.ExtraInfo {
	infos := []*npool.ExtraInfo{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
