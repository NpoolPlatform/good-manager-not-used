package extrainfo

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

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/extrainfo"

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

var appDate = npool.ExtraInfo{
	ID:        uuid.NewString(),
	GoodID:    uuid.NewString(),
	Posters:   []string{uuid.NewString()},
	Labels:    []string{uuid.NewString()},
	VoteCount: 1001,
	Rating:    1002,
}

var (
	appInfo = npool.ExtraInfoReq{
		ID:        &appDate.ID,
		GoodID:    &appDate.GoodID,
		Posters:   appDate.Posters,
		Labels:    appDate.Labels,
		VoteCount: &appDate.VoteCount,
		Rating:    &appDate.Rating,
	}
)

var info *npool.ExtraInfo

func createExtraInfo(t *testing.T) {
	var err error
	info, err = CreateExtraInfo(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createExtraInfos(t *testing.T) {
	appDates := []npool.ExtraInfo{
		{
			ID:        uuid.NewString(),
			GoodID:    uuid.NewString(),
			Posters:   []string{uuid.NewString()},
			Labels:    []string{uuid.NewString()},
			VoteCount: 1001,
			Rating:    1002,
		},
		{
			ID:        uuid.NewString(),
			GoodID:    uuid.NewString(),
			Posters:   []string{uuid.NewString()},
			Labels:    []string{uuid.NewString()},
			VoteCount: 1001,
			Rating:    1002,
		},
	}

	apps := []*npool.ExtraInfoReq{}
	for key := range appDates {
		apps = append(apps, &npool.ExtraInfoReq{
			ID:        &appDates[key].ID,
			GoodID:    &appDates[key].GoodID,
			Posters:   appDates[key].Posters,
			Labels:    appDates[key].Labels,
			VoteCount: &appDates[key].VoteCount,
			Rating:    &appDates[key].Rating,
		})
	}

	infos, err := CreateExtraInfos(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateExtraInfo(t *testing.T) {
	var err error
	info, err = UpdateExtraInfo(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getExtraInfo(t *testing.T) {
	var err error
	info, err = GetExtraInfo(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getExtraInfos(t *testing.T) {
	infos, total, err := GetExtraInfos(context.Background(),
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

func getExtraInfoOnly(t *testing.T) {
	var err error
	info, err = GetExtraInfoOnly(context.Background(),
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

func existExtraInfo(t *testing.T) {
	exist, err := ExistExtraInfo(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existExtraInfoConds(t *testing.T) {
	exist, err := ExistExtraInfoConds(context.Background(),
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

func deleteExtraInfo(t *testing.T) {
	info, err := DeleteExtraInfo(context.Background(), info.ID)
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

	t.Run("createExtraInfo", createExtraInfo)
	t.Run("createExtraInfos", createExtraInfos)
	t.Run("getExtraInfo", getExtraInfo)
	t.Run("getExtraInfos", getExtraInfos)
	t.Run("getExtraInfoOnly", getExtraInfoOnly)
	t.Run("updateExtraInfo", updateExtraInfo)
	t.Run("existExtraInfo", existExtraInfo)
	t.Run("existExtraInfoConds", existExtraInfoConds)
	t.Run("delete", deleteExtraInfo)
}
