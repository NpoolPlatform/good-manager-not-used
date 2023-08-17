package recommend

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

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/recommend"

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

var appDate = npool.Recommend{
	ID:             uuid.NewString(),
	AppID:          uuid.NewString(),
	GoodID:         uuid.NewString(),
	RecommenderID:  uuid.NewString(),
	Message:        uuid.NewString(),
	RecommendIndex: 10001,
}

var (
	appInfo = npool.RecommendReq{
		ID:             &appDate.ID,
		AppID:          &appDate.AppID,
		GoodID:         &appDate.GoodID,
		RecommenderID:  &appDate.RecommenderID,
		Message:        &appDate.Message,
		RecommendIndex: &appDate.RecommendIndex,
	}
)

var info *npool.Recommend

func createRecommend(t *testing.T) {
	var err error
	info, err = CreateRecommend(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createRecommends(t *testing.T) {
	appDates := []npool.Recommend{
		{
			ID:             uuid.NewString(),
			AppID:          uuid.NewString(),
			GoodID:         uuid.NewString(),
			RecommenderID:  uuid.NewString(),
			Message:        uuid.NewString(),
			RecommendIndex: 10001,
		},
		{
			ID:             uuid.NewString(),
			AppID:          uuid.NewString(),
			GoodID:         uuid.NewString(),
			RecommenderID:  uuid.NewString(),
			Message:        uuid.NewString(),
			RecommendIndex: 10001,
		},
	}

	apps := []*npool.RecommendReq{}
	for key := range appDates {
		apps = append(apps, &npool.RecommendReq{
			ID:             &appDates[key].ID,
			AppID:          &appDates[key].AppID,
			GoodID:         &appDates[key].GoodID,
			RecommenderID:  &appDates[key].RecommenderID,
			Message:        &appDates[key].Message,
			RecommendIndex: &appDates[key].RecommendIndex,
		})
	}

	infos, err := CreateRecommends(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateRecommend(t *testing.T) {
	var err error
	info, err = UpdateRecommend(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getRecommend(t *testing.T) {
	var err error
	info, err = GetRecommend(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getRecommends(t *testing.T) {
	infos, total, err := GetRecommends(context.Background(),
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

func getRecommendOnly(t *testing.T) {
	var err error
	info, err = GetRecommendOnly(context.Background(),
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

func existRecommend(t *testing.T) {
	exist, err := ExistRecommend(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existRecommendConds(t *testing.T) {
	exist, err := ExistRecommendConds(context.Background(),
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

func deleteRecommend(t *testing.T) {
	info, err := DeleteRecommend(context.Background(), info.ID)
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

	t.Run("createRecommend", createRecommend)
	t.Run("createRecommends", createRecommends)
	t.Run("getRecommend", getRecommend)
	t.Run("getRecommends", getRecommends)
	t.Run("getRecommendOnly", getRecommendOnly)
	t.Run("updateRecommend", updateRecommend)
	t.Run("existRecommend", existRecommend)
	t.Run("existRecommendConds", existRecommendConds)
	t.Run("delete", deleteRecommend)
}
