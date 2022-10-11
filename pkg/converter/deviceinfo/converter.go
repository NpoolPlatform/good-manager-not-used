package deviceinfo

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.DeviceInfo) *npool.DeviceInfo {
	if row == nil {
		return nil
	}

	return &npool.DeviceInfo{
		ID:              row.ID.String(),
		Type:            row.Type,
		Manufacturer:    row.Manufacturer,
		PowerComsuption: row.PowerComsuption,
		ShipmentAt:      row.ShipmentAt,
		Posters:         row.Posters,
		CreatedAt:       row.CreatedAt,
		UpdatedAt:       row.UpdatedAt,
		DeletedAt:       row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.DeviceInfo) []*npool.DeviceInfo {
	infos := []*npool.DeviceInfo{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
