package comment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/comment"

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

var comment = ent.Comment{
	ID:        uuid.New(),
	AppID:     uuid.New(),
	UserID:    uuid.New(),
	GoodID:    uuid.New(),
	OrderID:   uuid.New(),
	Content:   uuid.NewString(),
	ReplyToID: uuid.New(),
}

var (
	id        = comment.ID.String()
	appID     = comment.AppID.String()
	goodID    = comment.GoodID.String()
	userID    = comment.UserID.String()
	orderID   = comment.OrderID.String()
	replyToID = comment.ReplyToID.String()
	req       = npool.CommentReq{
		ID:        &id,
		AppID:     &appID,
		UserID:    &userID,
		GoodID:    &goodID,
		OrderID:   &orderID,
		Content:   &comment.Content,
		ReplyToID: &replyToID,
	}
)

var info *ent.Comment

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		comment.UpdatedAt = info.UpdatedAt
		comment.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), comment.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.Comment{
		{
			ID:        uuid.New(),
			AppID:     uuid.New(),
			UserID:    uuid.New(),
			GoodID:    uuid.New(),
			OrderID:   uuid.New(),
			Content:   uuid.NewString(),
			ReplyToID: uuid.New(),
		},
		{
			ID:        uuid.New(),
			AppID:     uuid.New(),
			UserID:    uuid.New(),
			GoodID:    uuid.New(),
			OrderID:   uuid.New(),
			Content:   uuid.NewString(),
			ReplyToID: uuid.New(),
		},
	}

	reqs := []*npool.CommentReq{}
	for _, _comment := range entities {
		_id := _comment.ID.String()
		_appID := _comment.AppID.String()
		_goodID := _comment.GoodID.String()
		_userID := _comment.UserID.String()
		_orderID := _comment.OrderID.String()
		_replyToID := _comment.ReplyToID.String()
		reqs = append(reqs, &npool.CommentReq{
			ID:        &_id,
			AppID:     &_appID,
			UserID:    &_userID,
			GoodID:    &_goodID,
			OrderID:   &_orderID,
			Content:   &_comment.Content,
			ReplyToID: &_replyToID,
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
		comment.UpdatedAt = info.UpdatedAt
		comment.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), comment.String())
	}
}
func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), comment.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), comment.String())
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
			assert.Equal(t, infos[0].String(), comment.String())
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
		assert.Equal(t, info.String(), comment.String())
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
	exist, err := Exist(context.Background(), comment.ID)
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
	info, err := Delete(context.Background(), comment.ID.String())
	if assert.Nil(t, err) {
		comment.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), comment.String())
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
