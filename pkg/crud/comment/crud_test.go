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

var commentInfo = ent.Comment{
	ID:        uuid.New(),
	AppID:     uuid.New(),
	UserID:    uuid.New(),
	GoodID:    uuid.New(),
	OrderID:   uuid.New(),
	Content:   uuid.NewString(),
	ReplyToID: uuid.New(),
}

var (
	id        = commentInfo.ID.String()
	appID     = commentInfo.AppID.String()
	goodID    = commentInfo.GoodID.String()
	userID    = commentInfo.UserID.String()
	orderID   = commentInfo.OrderID.String()
	replyToID = commentInfo.ReplyToID.String()
	req       = npool.CommentReq{
		ID:        &id,
		AppID:     &appID,
		UserID:    &userID,
		GoodID:    &goodID,
		OrderID:   &orderID,
		Content:   &commentInfo.Content,
		ReplyToID: &replyToID,
	}
)

var info *ent.Comment

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		commentInfo.UpdatedAt = info.UpdatedAt
		commentInfo.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), commentInfo.String())
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
	for _, _commentInfo := range entities {
		_id := _commentInfo.ID.String()
		_appID := _commentInfo.AppID.String()
		_goodID := _commentInfo.GoodID.String()
		_userID := _commentInfo.UserID.String()
		_orderID := _commentInfo.OrderID.String()
		_replyToID := _commentInfo.ReplyToID.String()
		reqs = append(reqs, &npool.CommentReq{
			ID:        &_id,
			AppID:     &_appID,
			UserID:    &_userID,
			GoodID:    &_goodID,
			OrderID:   &_orderID,
			Content:   &_commentInfo.Content,
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
		commentInfo.UpdatedAt = info.UpdatedAt
		commentInfo.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), commentInfo.String())
	}
}
func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), commentInfo.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), commentInfo.String())
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
			assert.Equal(t, infos[0].String(), commentInfo.String())
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
		assert.Equal(t, info.String(), commentInfo.String())
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
	exist, err := Exist(context.Background(), commentInfo.ID)
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
	info, err := Delete(context.Background(), commentInfo.ID.String())
	if assert.Nil(t, err) {
		commentInfo.DeletedAt = info.DeletedAt
		commentInfo.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), commentInfo.String())
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
