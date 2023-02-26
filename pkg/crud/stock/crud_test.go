package stock

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"

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
	Total:  decimal.NewFromInt(1005),
}

var (
	id        = _stock.ID.String()
	goodID    = _stock.GoodID.String()
	total     = _stock.Total.String()
	locked    = _stock.Locked.String()
	inService = _stock.InService.String()
	req       = npool.StockReq{
		ID:        &id,
		GoodID:    &goodID,
		Total:     &total,
		Locked:    &locked,
		InService: &inService,
		Sold:      &inService,
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
			ID:     uuid.New(),
			GoodID: uuid.New(),
			Total:  decimal.NewFromInt(1005),
		},
		{
			ID:     uuid.New(),
			GoodID: uuid.New(),
			Total:  decimal.NewFromInt(1005),
		},
	}

	reqs := []*npool.StockReq{}
	for _, _stock1 := range entities {
		_id := _stock1.ID.String()
		_goodID := _stock1.GoodID.String()
		_total := _stock1.Total.String()
		_locked := _stock1.Locked.String()
		_inService := _stock1.InService.String()
		_sold := _stock1.Sold.String()
		reqs = append(reqs, &npool.StockReq{
			ID:        &_id,
			GoodID:    &_goodID,
			Total:     &_total,
			Locked:    &_locked,
			InService: &_inService,
			Sold:      &_sold,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	var err error

	total := decimal.NewFromInt(2000)
	_total := total.String()
	req.Total = &_total
	_stock.Total = total

	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		_stock.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), _stock.String())
	}
}

func addFields(t *testing.T) {
	locked := decimal.NewFromInt(10)
	lockedStr := locked.String()
	waitStart := decimal.NewFromInt(10)
	waitStartStr := waitStart.String()
	inService := decimal.NewFromInt(20)
	inServiceStr := inService.String()
	sold := decimal.NewFromInt(20)
	soldStr := sold.String()

	req.Locked = &lockedStr
	req.WaitStart = &waitStartStr
	req.InService = &inServiceStr
	req.Sold = &soldStr

	_stock.Locked = _stock.Locked.Add(locked)
	_stock.WaitStart = _stock.WaitStart.Add(waitStart)
	_stock.InService = _stock.InService.Add(inService)
	_stock.Sold = _stock.Sold.Add(waitStart)

	info, err := AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		_stock.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), _stock.String())
	}

	locked = decimal.NewFromInt(-5)
	lockedStr = locked.String()
	inService = decimal.NewFromInt(-3)
	inServiceStr = inService.String()

	_stock.Locked = _stock.Locked.Add(locked)
	_stock.InService = _stock.InService.Add(inService)

	req.Locked = &lockedStr
	req.InService = &inServiceStr
	req.WaitStart = nil

	info, err = AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		_stock.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), _stock.String())
	}

	locked = decimal.NewFromInt(3000)
	lockedStr = locked.String()
	req.Locked = &lockedStr
	_, err = AddFields(context.Background(), &req)
	assert.NotNil(t, err)

	inService = decimal.NewFromInt(3000)
	inServiceStr = inService.String()
	req.InService = &inServiceStr
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
