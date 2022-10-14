package stock

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
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

var _stock = ent.Stock{
	ID:     uuid.New(),
	GoodID: uuid.New(),
	Total:  1005,
}

var (
	id        = _stock.ID.String()
	goodID    = _stock.GoodID.String()
	locked    = int32(_stock.Locked)
	inService = int32(_stock.InService)
	req       = npool.StockReq{
		ID:        &id,
		GoodID:    &goodID,
		Total:     &_stock.Total,
		Locked:    &locked,
		InService: &inService,
		Sold:      &_stock.Sold,
	}
)

var info *ent.Stock

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		_stock.UpdatedAt = info.UpdatedAt
		_stock.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), _stock.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.Stock{
		{
			ID:        uuid.New(),
			GoodID:    uuid.New(),
			Total:     1005,
			Locked:    0,
			InService: 0,
			Sold:      0,
		},
		{
			ID:        uuid.New(),
			GoodID:    uuid.New(),
			Total:     1005,
			Locked:    0,
			InService: 0,
			Sold:      0,
		},
	}

	reqs := []*npool.StockReq{}
	for _, _stock1 := range entities {
		_id := _stock1.ID.String()
		_goodID := _stock1.GoodID.String()
		_locked := int32(_stock1.Locked)
		_inService := int32(_stock1.InService)
		reqs = append(reqs, &npool.StockReq{
			ID:        &_id,
			GoodID:    &_goodID,
			Total:     &_stock1.Total,
			Locked:    &_locked,
			InService: &_inService,
			Sold:      &_stock1.Sold,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	var err error

	total := uint32(2000)
	req.Total = &total
	_stock.Total = total

	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		_stock.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), _stock.String())
	}
}

func addFields(t *testing.T) {
	locked := int32(10)
	inService := int32(20)
	sold := uint32(20)

	req.Locked = &locked
	req.InService = &inService
	req.Sold = &sold

	_stock.Locked = uint32(int32(_stock.Locked) + locked)
	_stock.InService = uint32(int32(_stock.InService) + inService)
	_stock.Sold += sold

	info, err := AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		_stock.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), _stock.String())
	}

	locked = -5
	inService = -3

	_stock.Locked = uint32(int32(_stock.Locked) + locked)
	_stock.InService = uint32(int32(_stock.InService) + inService)

	req.Locked = &locked
	req.InService = &inService

	info, err = AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		_stock.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), _stock.String())
	}

	locked = 3000
	req.Locked = &locked
	_, err = AddFields(context.Background(), &req)
	assert.NotNil(t, err)

	inService = 3000
	req.InService = &inService
	_, err = AddFields(context.Background(), &req)
	assert.NotNil(t, err)
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), _stock.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), _stock.String())
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
			assert.Equal(t, infos[0].String(), _stock.String())
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
		assert.Equal(t, info.String(), _stock.String())
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
	exist, err := Exist(context.Background(), _stock.ID)
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
	info, err := Delete(context.Background(), _stock.ID.String())
	if assert.Nil(t, err) {
		_stock.DeletedAt = info.DeletedAt
		_stock.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), _stock.String())
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("update", update)
	t.Run("addFields", addFields)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
