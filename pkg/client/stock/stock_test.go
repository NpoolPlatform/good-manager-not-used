package stock

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"

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

var ret = npool.Stock{
	ID:        uuid.NewString(),
	GoodID:    uuid.NewString(),
	Total:     decimal.NewFromInt(1005).String(),
	Locked:    decimal.NewFromInt(0).String(),
	InService: decimal.NewFromInt(0).String(),
	WaitStart: decimal.NewFromInt(0).String(),
	Sold:      decimal.NewFromInt(0).String(),
}

var (
	req = npool.StockReq{
		ID:     &ret.ID,
		GoodID: &ret.GoodID,
		Total:  &ret.Total,
	}
)

var info *npool.Stock

func createStock(t *testing.T) {
	var err error
	info, err = CreateStock(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func createStocks(t *testing.T) {
	rets := []npool.Stock{
		{
			ID:     uuid.NewString(),
			GoodID: uuid.NewString(),
			Total:  decimal.NewFromInt(1005).String(),
		},
		{
			ID:     uuid.NewString(),
			GoodID: uuid.NewString(),
			Total:  decimal.NewFromInt(1005).String(),
		},
	}

	apps := []*npool.StockReq{}
	for key := range rets {
		apps = append(apps, &npool.StockReq{
			ID:     &rets[key].ID,
			GoodID: &rets[key].GoodID,
			Total:  &rets[key].Total,
		})
	}

	infos, err := CreateStocks(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateStock(t *testing.T) {
	var err error
	info, err = UpdateStock(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getStock(t *testing.T) {
	var err error
	info, err = GetStock(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
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
		assert.Equal(t, infos[0], &ret)
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
		assert.Equal(t, info, &ret)
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
		ret.DeletedAt = info.DeletedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func TestStock(t *testing.T) {
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
