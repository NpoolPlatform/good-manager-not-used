package comment

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/comment"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Comment) *npool.Comment {
	if row == nil {
		return nil
	}

	return &npool.Comment{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		UserID:    row.UserID.String(),
		GoodID:    row.GoodID.String(),
		OrderID:   row.OrderID.String(),
		Content:   row.Content,
		ReplyToID: row.ReplyToID.String(),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Comment) []*npool.Comment {
	infos := []*npool.Comment{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
