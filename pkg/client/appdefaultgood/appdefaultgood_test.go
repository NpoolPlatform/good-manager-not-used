package appdefaultgood

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

	testinit "github.com/NpoolPlatform/good-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appdefaultgood"
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

var appDate = npool.AppDefaultGood{
	ID:         uuid.NewString(),
	AppID:      uuid.NewString(),
	GoodID:     uuid.NewString(),
	CoinTypeID: uuid.NewString(),
}

var (
	appInfo = npool.AppDefaultGoodReq{
		ID:         &appDate.ID,
		AppID:      &appDate.AppID,
		GoodID:     &appDate.GoodID,
		CoinTypeID: &appDate.CoinTypeID,
	}
)

var info *npool.AppDefaultGood

func createAppDefaultGood(t *testing.T) {
	var err error
	info, err = CreateAppDefaultGood(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createAppDefaultGoods(t *testing.T) {
	appID := uuid.NewString()
	appDates := []npool.AppDefaultGood{
		{
			ID:         uuid.NewString(),
			AppID:      appID,
			GoodID:     uuid.NewString(),
			CoinTypeID: uuid.NewString(),
		},
		{
			ID:         uuid.NewString(),
			AppID:      appID,
			GoodID:     uuid.NewString(),
			CoinTypeID: uuid.NewString(),
		},
	}

	apps := []*npool.AppDefaultGoodReq{}
	for key := range appDates {
		apps = append(apps, &npool.AppDefaultGoodReq{
			ID:         &appDates[key].ID,
			AppID:      &appDates[key].AppID,
			GoodID:     &appDates[key].GoodID,
			CoinTypeID: &appDates[key].CoinTypeID,
		})
	}

	infos, err := CreateAppDefaultGoods(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppDefaultGood(t *testing.T) {
	var err error
	info, err = UpdateAppDefaultGood(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getAppDefaultGood(t *testing.T) {
	var err error
	info, err = GetAppDefaultGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getAppDefaultGoods(t *testing.T) {
	infos, total, err := GetAppDefaultGoods(context.Background(),
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

func getAppDefaultGoodOnly(t *testing.T) {
	var err error
	info, err = GetAppDefaultGoodOnly(context.Background(),
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

func existAppDefaultGood(t *testing.T) {
	exist, err := ExistAppDefaultGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppDefaultGoodConds(t *testing.T) {
	exist, err := ExistAppDefaultGoodConds(context.Background(),
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

func deleteAppDefaultGood(t *testing.T) {
	info, err := DeleteAppDefaultGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
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

	t.Run("createAppDefaultGood", createAppDefaultGood)
	t.Run("createAppDefaultGoods", createAppDefaultGoods)
	t.Run("getAppDefaultGood", getAppDefaultGood)
	t.Run("getAppDefaultGoods", getAppDefaultGoods)
	t.Run("getAppDefaultGoodOnly", getAppDefaultGoodOnly)
	t.Run("updateAppDefaultGood", updateAppDefaultGood)
	t.Run("existAppDefaultGood", existAppDefaultGood)
	t.Run("existAppDefaultGoodConds", existAppDefaultGoodConds)
	t.Run("delete", deleteAppDefaultGood)
}
