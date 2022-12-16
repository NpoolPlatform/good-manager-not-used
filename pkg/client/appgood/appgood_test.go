package appgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	testinit "github.com/NpoolPlatform/good-manager/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appDate = npool.AppGood{
	ID:                uuid.NewString(),
	AppID:             uuid.NewString(),
	GoodID:            uuid.NewString(),
	Online:            true,
	Visible:           true,
	GoodName:          uuid.NewString(),
	Price:             "999",
	DisplayIndex:      100,
	PurchaseLimit:     101,
	CommissionPercent: 102,
	DailyRewardAmount: "999",
}

var (
	appInfo = npool.AppGoodReq{
		ID:                &appDate.ID,
		AppID:             &appDate.AppID,
		GoodID:            &appDate.GoodID,
		Online:            &appDate.Online,
		Visible:           &appDate.Visible,
		GoodName:          &appDate.GoodName,
		Price:             &appDate.Price,
		DisplayIndex:      &appDate.DisplayIndex,
		PurchaseLimit:     &appDate.PurchaseLimit,
		CommissionPercent: &appDate.CommissionPercent,
		DailyRewardAmount: &appDate.DailyRewardAmount,
	}
)

var info *npool.AppGood

func createAppGood(t *testing.T) {
	var err error
	info, err = CreateAppGood(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createAppGoods(t *testing.T) {
	appID := uuid.NewString()
	appDates := []npool.AppGood{
		{
			ID:                uuid.NewString(),
			AppID:             appID,
			GoodID:            uuid.NewString(),
			Online:            true,
			Visible:           true,
			GoodName:          uuid.NewString(),
			Price:             "999",
			DisplayIndex:      100,
			PurchaseLimit:     101,
			CommissionPercent: 102,
			DailyRewardAmount: "999",
		},
		{
			ID:                uuid.NewString(),
			AppID:             appID,
			GoodID:            uuid.NewString(),
			Online:            true,
			Visible:           true,
			GoodName:          uuid.NewString(),
			Price:             "999",
			DisplayIndex:      100,
			PurchaseLimit:     101,
			CommissionPercent: 102,
			DailyRewardAmount: "999",
		},
	}

	apps := []*npool.AppGoodReq{}
	for key := range appDates {
		apps = append(apps, &npool.AppGoodReq{
			ID:                &appDates[key].ID,
			AppID:             &appDates[key].AppID,
			GoodID:            &appDates[key].GoodID,
			Online:            &appDates[key].Online,
			Visible:           &appDates[key].Visible,
			GoodName:          &appDates[key].GoodName,
			Price:             &appDates[key].Price,
			DisplayIndex:      &appDates[key].DisplayIndex,
			PurchaseLimit:     &appDates[key].PurchaseLimit,
			CommissionPercent: &appDates[key].CommissionPercent,
			DailyRewardAmount: &appDates[key].DailyRewardAmount,
		})
	}

	infos, err := CreateAppGoods(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppGood(t *testing.T) {
	var err error
	info, err = UpdateAppGood(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getAppGood(t *testing.T) {
	var err error
	info, err = GetAppGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getAppGoods(t *testing.T) {
	infos, total, err := GetAppGoods(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appDate)
	}
}

func getAppGoodOnly(t *testing.T) {
	var err error
	info, err = GetAppGoodOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func existAppGood(t *testing.T) {
	exist, err := ExistAppGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistAppGoodConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteAppGood(t *testing.T) {
	info, err := DeleteAppGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.DeletedAt = info.DeletedAt
		assert.Equal(t, info, &appDate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAppGood", createAppGood)
	t.Run("createAppGoods", createAppGoods)
	t.Run("getAppGood", getAppGood)
	t.Run("getAppGoods", getAppGoods)
	t.Run("getAppGoodOnly", getAppGoodOnly)
	t.Run("updateAppGood", updateAppGood)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("delete", deleteAppGood)
}
