package extrainfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
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

var deviceInfo = ent.ExtraInfo{
	ID:        uuid.New(),
	GoodID:    uuid.New(),
	Posters:   []string{uuid.NewString()},
	Labels:    []string{uuid.NewString()},
	VoteCount: 100,
	Rating:    100,
}

var (
	id     = deviceInfo.ID.String()
	goodID = deviceInfo.GoodID.String()
	req    = npool.ExtraInfoReq{
		ID:        &id,
		GoodID:    &goodID,
		Posters:   deviceInfo.Posters,
		Labels:    deviceInfo.Labels,
		VoteCount: &deviceInfo.VoteCount,
		Rating:    &deviceInfo.Rating,
	}
)

var info *ent.ExtraInfo

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		deviceInfo.UpdatedAt = info.UpdatedAt
		deviceInfo.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), deviceInfo.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.ExtraInfo{
		{
			ID:        uuid.New(),
			GoodID:    uuid.New(),
			Posters:   []string{uuid.NewString()},
			Labels:    []string{uuid.NewString()},
			VoteCount: 100,
			Rating:    100,
		},
		{
			ID:        uuid.New(),
			GoodID:    uuid.New(),
			Posters:   []string{uuid.NewString()},
			Labels:    []string{uuid.NewString()},
			VoteCount: 100,
			Rating:    100,
		},
	}

	reqs := []*npool.ExtraInfoReq{}
	for _, _deviceInfo := range entities {
		_id := _deviceInfo.ID.String()
		_goodID := _deviceInfo.GoodID.String()
		reqs = append(reqs, &npool.ExtraInfoReq{
			ID:        &_id,
			GoodID:    &_goodID,
			Posters:   _deviceInfo.Posters,
			Labels:    _deviceInfo.Labels,
			VoteCount: &_deviceInfo.VoteCount,
			Rating:    &_deviceInfo.Rating,
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
		deviceInfo.UpdatedAt = info.UpdatedAt
		deviceInfo.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), deviceInfo.String())
	}
}
func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), deviceInfo.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), deviceInfo.String())
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
			assert.Equal(t, infos[0].String(), deviceInfo.String())
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
		assert.Equal(t, info.String(), deviceInfo.String())
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
	exist, err := Exist(context.Background(), deviceInfo.ID)
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
	info, err := Delete(context.Background(), deviceInfo.ID.String())
	if assert.Nil(t, err) {
		deviceInfo.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), deviceInfo.String())
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
