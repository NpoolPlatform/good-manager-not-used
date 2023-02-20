package comment

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

var ret = npool.Comment{
	ID:        uuid.NewString(),
	AppID:     uuid.NewString(),
	UserID:    uuid.NewString(),
	GoodID:    uuid.NewString(),
	OrderID:   uuid.NewString(),
	Content:   uuid.NewString(),
	ReplyToID: uuid.NewString(),
}

var (
	req = npool.CommentReq{
		ID:        &ret.ID,
		AppID:     &ret.AppID,
		UserID:    &ret.UserID,
		GoodID:    &ret.GoodID,
		OrderID:   &ret.OrderID,
		Content:   &ret.Content,
		ReplyToID: &ret.ReplyToID,
	}
)

var info *npool.Comment

func createComment(t *testing.T) {
	var err error
	info, err = CreateComment(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func createComments(t *testing.T) {
	rets := []npool.Comment{
		{
			ID:        uuid.NewString(),
			AppID:     uuid.NewString(),
			UserID:    uuid.NewString(),
			GoodID:    uuid.NewString(),
			OrderID:   uuid.NewString(),
			Content:   uuid.NewString(),
			ReplyToID: uuid.NewString(),
		},
		{
			ID:        uuid.NewString(),
			AppID:     uuid.NewString(),
			UserID:    uuid.NewString(),
			GoodID:    uuid.NewString(),
			OrderID:   uuid.NewString(),
			Content:   uuid.NewString(),
			ReplyToID: uuid.NewString(),
		},
	}

	apps := []*npool.CommentReq{}
	for key := range rets {
		apps = append(apps, &npool.CommentReq{
			ID:        &rets[key].ID,
			AppID:     &rets[key].AppID,
			UserID:    &rets[key].UserID,
			GoodID:    &rets[key].GoodID,
			OrderID:   &rets[key].OrderID,
			Content:   &rets[key].Content,
			ReplyToID: &rets[key].ReplyToID,
		})
	}

	infos, err := CreateComments(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateComment(t *testing.T) {
	var err error
	info, err = UpdateComment(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getComment(t *testing.T) {
	var err error
	info, err = GetComment(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getComments(t *testing.T) {
	infos, total, err := GetComments(context.Background(),
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

func getCommentOnly(t *testing.T) {
	var err error
	info, err = GetCommentOnly(context.Background(),
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

func existComment(t *testing.T) {
	exist, err := ExistComment(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existCommentConds(t *testing.T) {
	exist, err := ExistCommentConds(context.Background(),
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

func deleteComment(t *testing.T) {
	info, err := DeleteComment(context.Background(), info.ID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.DeletedAt = info.DeletedAt
		assert.Equal(t, info, &ret)
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

	t.Run("createComment", createComment)
	t.Run("createComments", createComments)
	t.Run("getComment", getComment)
	t.Run("getComments", getComments)
	t.Run("getCommentOnly", getCommentOnly)
	t.Run("updateComment", updateComment)
	t.Run("existComment", existComment)
	t.Run("existCommentConds", existCommentConds)
	t.Run("delete", deleteComment)
}
