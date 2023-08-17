package subgood

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

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

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

var appDate = npool.SubGood{
	ID:         uuid.NewString(),
	AppID:      uuid.NewString(),
	MainGoodID: uuid.NewString(),
	SubGoodID:  uuid.NewString(),
	Must:       true,
	Commission: true,
}

var (
	appInfo = npool.SubGoodReq{
		ID:         &appDate.ID,
		AppID:      &appDate.AppID,
		MainGoodID: &appDate.MainGoodID,
		SubGoodID:  &appDate.SubGoodID,
		Must:       &appDate.Must,
		Commission: &appDate.Commission,
		CreatedAt:  &appDate.CreatedAt,
	}
)

var info *npool.SubGood

func createSubGood(t *testing.T) {
	var err error
	info, err = CreateSubGood(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createSubGoods(t *testing.T) {
	appDates := []npool.SubGood{
		{
			ID:         uuid.NewString(),
			AppID:      uuid.NewString(),
			MainGoodID: uuid.NewString(),
			SubGoodID:  uuid.NewString(),
			Must:       true,
			Commission: true,
		},
		{
			ID:         uuid.NewString(),
			AppID:      uuid.NewString(),
			MainGoodID: uuid.NewString(),
			SubGoodID:  uuid.NewString(),
			Must:       true,
			Commission: true,
		},
	}

	apps := []*npool.SubGoodReq{}
	for key := range appDates {
		apps = append(apps, &npool.SubGoodReq{
			ID:         &appDates[key].ID,
			AppID:      &appDates[key].AppID,
			MainGoodID: &appDates[key].MainGoodID,
			SubGoodID:  &appDates[key].SubGoodID,
			Must:       &appDates[key].Must,
			Commission: &appDates[key].Commission,
			CreatedAt:  &appDates[key].CreatedAt,
		})
	}

	infos, err := CreateSubGoods(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateSubGood(t *testing.T) {
	var err error
	info, err = UpdateSubGood(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getSubGood(t *testing.T) {
	var err error
	info, err = GetSubGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getSubGoods(t *testing.T) {
	infos, total, err := GetSubGoods(context.Background(),
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

func getSubGoodOnly(t *testing.T) {
	var err error
	info, err = GetSubGoodOnly(context.Background(),
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

func existSubGood(t *testing.T) {
	exist, err := ExistSubGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existSubGoodConds(t *testing.T) {
	exist, err := ExistSubGoodConds(context.Background(),
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

func deleteSubGood(t *testing.T) {
	info, err := DeleteSubGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.DeletedAt = info.DeletedAt
		appDate.UpdatedAt = info.UpdatedAt
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

	t.Run("createSubGood", createSubGood)
	t.Run("createSubGoods", createSubGoods)
	t.Run("getSubGood", getSubGood)
	t.Run("getSubGoods", getSubGoods)
	t.Run("getSubGoodOnly", getSubGoodOnly)
	t.Run("updateSubGood", updateSubGood)
	t.Run("existSubGood", existSubGood)
	t.Run("existSubGoodConds", existSubGoodConds)
	t.Run("delete", deleteSubGood)
}
