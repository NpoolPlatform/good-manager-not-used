package promotion

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

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/promotion"

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

var appDate = npool.Promotion{
	ID:      uuid.NewString(),
	AppID:   uuid.NewString(),
	GoodID:  uuid.NewString(),
	Message: uuid.NewString(),
	StartAt: 1001,
	EndAt:   1002,
	Price:   "9999999999999999999.999999999999999999",
	Posters: []string{uuid.NewString()},
}

var (
	appInfo = npool.PromotionReq{
		ID:      &appDate.ID,
		AppID:   &appDate.AppID,
		GoodID:  &appDate.GoodID,
		Message: &appDate.Message,
		StartAt: &appDate.StartAt,
		EndAt:   &appDate.EndAt,
		Price:   &appDate.Price,
		Posters: appDate.Posters,
	}
)

var info *npool.Promotion

func createPromotion(t *testing.T) {
	var err error
	info, err = CreatePromotion(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createPromotions(t *testing.T) {
	appID := uuid.NewString()

	promotions := []npool.Promotion{
		{
			ID:      uuid.NewString(),
			AppID:   appID,
			GoodID:  uuid.NewString(),
			Message: uuid.NewString(),
			StartAt: 1001,
			EndAt:   1002,
			Price:   "9999999999999999999.999999999999999999",
			Posters: []string{uuid.NewString()},
		},
		{
			ID:      uuid.NewString(),
			AppID:   appID,
			GoodID:  uuid.NewString(),
			Message: uuid.NewString(),
			StartAt: 1001,
			EndAt:   1002,
			Price:   "9999999999999999999.999999999999999999",
			Posters: []string{uuid.NewString()},
		},
	}

	apps := []*npool.PromotionReq{}
	for key := range promotions {
		apps = append(apps, &npool.PromotionReq{
			ID:      &promotions[key].ID,
			AppID:   &promotions[key].AppID,
			GoodID:  &promotions[key].GoodID,
			Message: &promotions[key].Message,
			StartAt: &promotions[key].StartAt,
			EndAt:   &promotions[key].EndAt,
			Price:   &promotions[key].Price,
			Posters: promotions[key].Posters,
		})
	}

	infos, err := CreatePromotions(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updatePromotion(t *testing.T) {
	var err error
	info, err = UpdatePromotion(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getPromotion(t *testing.T) {
	var err error
	info, err = GetPromotion(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getPromotions(t *testing.T) {
	infos, total, err := GetPromotions(context.Background(),
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

func getPromotionOnly(t *testing.T) {
	var err error
	info, err = GetPromotionOnly(context.Background(),
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

func existPromotion(t *testing.T) {
	exist, err := ExistPromotion(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existPromotionConds(t *testing.T) {
	exist, err := ExistPromotionConds(context.Background(),
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

func deletePromotion(t *testing.T) {
	info, err := DeletePromotion(context.Background(), info.ID)
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

	t.Run("createPromotion", createPromotion)
	t.Run("createPromotions", createPromotions)
	t.Run("getPromotion", getPromotion)
	t.Run("getPromotions", getPromotions)
	t.Run("getPromotionOnly", getPromotionOnly)
	t.Run("updatePromotion", updatePromotion)
	t.Run("existPromotion", existPromotion)
	t.Run("existPromotionConds", existPromotionConds)
	t.Run("delete", deletePromotion)
}
