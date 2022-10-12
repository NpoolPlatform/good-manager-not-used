package stock

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

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"

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

var appDate = npool.Stock{
	ID:        uuid.NewString(),
	GoodID:    uuid.NewString(),
	Total:     1005,
	Locked:    12,
	InService: 13,
	Sold:      14,
}

var (
	appInfo = npool.StockReq{
		ID:        &appDate.ID,
		GoodID:    &appDate.GoodID,
		Total:     &appDate.Total,
		Locked:    &appDate.Locked,
		InService: &appDate.InService,
		Sold:      &appDate.Sold,
	}
)

var info *npool.Stock

func createStock(t *testing.T) {
	var err error
	info, err = CreateStock(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createStocks(t *testing.T) {
	appDates := []npool.Stock{
		{
			ID:        uuid.NewString(),
			GoodID:    uuid.NewString(),
			Total:     1005,
			Locked:    12,
			InService: 13,
			Sold:      14,
		},
		{
			ID:        uuid.NewString(),
			GoodID:    uuid.NewString(),
			Total:     1005,
			Locked:    12,
			InService: 13,
			Sold:      14,
		},
	}

	apps := []*npool.StockReq{}
	for key := range appDates {
		apps = append(apps, &npool.StockReq{
			ID:        &appDates[key].ID,
			GoodID:    &appDates[key].GoodID,
			Total:     &appDates[key].Total,
			Locked:    &appDates[key].Locked,
			InService: &appDates[key].InService,
			Sold:      &appDates[key].Sold,
		})
	}

	infos, err := CreateStocks(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateStock(t *testing.T) {
	var err error
	info, err = UpdateStock(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getStock(t *testing.T) {
	var err error
	info, err = GetStock(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getStocks(t *testing.T) {
	infos, total, err := GetStocks(context.Background(),
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

func getStockOnly(t *testing.T) {
	var err error
	info, err = GetStockOnly(context.Background(),
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

func existStock(t *testing.T) {
	exist, err := ExistStock(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existStockConds(t *testing.T) {
	exist, err := ExistStockConds(context.Background(),
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

func deleteStock(t *testing.T) {
	info, err := DeleteStock(context.Background(), info.ID)
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

	t.Run("createStock", createStock)
	t.Run("createStocks", createStocks)
	t.Run("getStock", getStock)
	t.Run("getStocks", getStocks)
	t.Run("getStockOnly", getStockOnly)
	t.Run("updateStock", updateStock)
	t.Run("existStock", existStock)
	t.Run("existStockConds", existStockConds)
	t.Run("delete", deleteStock)
}
