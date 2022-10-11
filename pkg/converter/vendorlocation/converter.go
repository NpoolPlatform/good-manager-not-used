package vendorlocation

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/vendorlocation"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.VendorLocation) *npool.VendorLocation {
	if row == nil {
		return nil
	}

	return &npool.VendorLocation{
		ID:        row.ID.String(),
		Country:   row.Country,
		Province:  row.Province,
		City:      row.City,
		Address:   row.Address,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.VendorLocation) []*npool.VendorLocation {
	infos := []*npool.VendorLocation{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
