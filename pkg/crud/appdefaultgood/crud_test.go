package appdefaultgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/good-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
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

var (
	appGood = ent.AppDefaultGood{
		ID:         uuid.New(),
		AppID:      uuid.New(),
		GoodID:     uuid.New(),
		CoinTypeID: uuid.New(),
	}
)

var (
	id         = appGood.ID.String()
	appID      = appGood.AppID.String()
	goodID     = appGood.GoodID.String()
	coinTypeID = appGood.CoinTypeID.String()
	req        = npool.AppDefaultGoodReq{
		ID:         &id,
		AppID:      &appID,
		GoodID:     &goodID,
		CoinTypeID: &coinTypeID,
	}
)

var info *ent.AppDefaultGood

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		appGood.UpdatedAt = info.UpdatedAt
		appGood.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.AppDefaultGood{
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			GoodID:     uuid.New(),
			CoinTypeID: uuid.New(),
		},
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			GoodID:     uuid.New(),
			CoinTypeID: uuid.New(),
		},
	}

	reqs := []*npool.AppDefaultGoodReq{}
	for _, _appGood := range entities {
		_id := _appGood.ID.String()
		_appID := _appGood.AppID.String()
		_goodID := _appGood.GoodID.String()
		_coinTypeID := _appGood.CoinTypeID.String()
		reqs = append(reqs, &npool.AppDefaultGoodReq{
			ID:         &_id,
			AppID:      &_appID,
			GoodID:     &_goodID,
			CoinTypeID: &_coinTypeID,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		appGood.UpdatedAt = info.UpdatedAt
		appGood.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}
func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), appGood.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), appGood.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), appGood.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), appGood.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), appGood.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), appGood.ID.String())
	if assert.Nil(t, err) {
		appGood.DeletedAt = info.DeletedAt
		appGood.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("update", update)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
