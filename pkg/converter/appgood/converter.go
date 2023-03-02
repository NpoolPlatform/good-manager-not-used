package appgood

import (
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.AppGood) *npool.AppGood {
	if row == nil {
		return nil
	}

	return &npool.AppGood{
		ID:                     row.ID.String(),
		AppID:                  row.AppID.String(),
		GoodID:                 row.GoodID.String(),
		Online:                 row.Online,
		Visible:                row.Visible,
		GoodName:               row.GoodName,
		Price:                  row.Price.String(),
		DisplayIndex:           row.DisplayIndex,
		PurchaseLimit:          row.PurchaseLimit,
		CommissionPercent:      row.CommissionPercent,
		CreatedAt:              row.CreatedAt,
		UpdatedAt:              row.UpdatedAt,
		DeletedAt:              row.DeletedAt,
		SaleStartAt:            row.SaleStartAt,
		SaleEndAt:              row.SaleEndAt,
		ServiceStartAt:         row.ServiceStartAt,
		TechnicalFeeRatio:      row.TechnicalFeeRatio,
		ElectricityFeeRatio:    row.ElectricityFeeRatio,
		DailyRewardAmount:      row.DailyRewardAmount.String(),
		CommissionSettleType:   commmgrpb.SettleType(commmgrpb.SettleType_value[row.CommissionSettleType]),
		Descriptions:           row.Descriptions,
		GoodBanner:             row.GoodBanner,
		DisplayNames:           row.DisplayNames,
		EnablePurchase:         row.EnablePurchase,
		EnableProductPage:      row.EnableProductPage,
		CancelMode:             npool.CancelMode(npool.CancelMode_value[row.CancelMode]),
		UserPurchaseLimit:      row.UserPurchaseLimit.String(),
		DisplayColors:          row.DisplayColors,
		CancellableBeforeStart: row.CancellableBeforeStart,
		ProductPage:            row.ProductPage,
	}
}

func Ent2GrpcMany(rows []*ent.AppGood) []*npool.AppGood {
	infos := []*npool.AppGood{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
