package deviceinfo

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

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"

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

var appDate = npool.DeviceInfo{
	ID:              uuid.NewString(),
	Type:            uuid.NewString(),
	Manufacturer:    uuid.NewString(),
	PowerComsuption: 1001,
	ShipmentAt:      1002,
	Posters:         []string{uuid.NewString()},
}

var (
	appInfo = npool.DeviceInfoReq{
		ID:              &appDate.ID,
		Type:            &appDate.Type,
		Manufacturer:    &appDate.Manufacturer,
		PowerComsuption: &appDate.PowerComsuption,
		ShipmentAt:      &appDate.ShipmentAt,
		Posters:         appDate.Posters,
	}
)

var info *npool.DeviceInfo

func createDeviceInfo(t *testing.T) {
	var err error
	info, err = CreateDeviceInfo(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createDeviceInfos(t *testing.T) {
	appDates := []npool.DeviceInfo{
		{
			ID:              uuid.NewString(),
			Type:            uuid.NewString(),
			Manufacturer:    uuid.NewString(),
			PowerComsuption: 1001,
			ShipmentAt:      1002,
			Posters:         []string{uuid.NewString()},
		},
		{
			ID:              uuid.NewString(),
			Type:            uuid.NewString(),
			Manufacturer:    uuid.NewString(),
			PowerComsuption: 1001,
			ShipmentAt:      1002,
			Posters:         []string{uuid.NewString()},
		},
	}

	apps := []*npool.DeviceInfoReq{}
	for key := range appDates {
		apps = append(apps, &npool.DeviceInfoReq{
			ID:              &appDates[key].ID,
			Type:            &appDates[key].Type,
			Manufacturer:    &appDates[key].Manufacturer,
			PowerComsuption: &appDates[key].PowerComsuption,
			ShipmentAt:      &appDates[key].ShipmentAt,
			Posters:         appDates[key].Posters,
		})
	}

	infos, err := CreateDeviceInfos(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateDeviceInfo(t *testing.T) {
	var err error
	info, err = UpdateDeviceInfo(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getDeviceInfo(t *testing.T) {
	var err error
	info, err = GetDeviceInfo(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getDeviceInfos(t *testing.T) {
	infos, total, err := GetDeviceInfos(context.Background(),
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

func getDeviceInfoOnly(t *testing.T) {
	var err error
	info, err = GetDeviceInfoOnly(context.Background(),
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

func existDeviceInfo(t *testing.T) {
	exist, err := ExistDeviceInfo(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existDeviceInfoConds(t *testing.T) {
	exist, err := ExistDeviceInfoConds(context.Background(),
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

func deleteDeviceInfo(t *testing.T) {
	info, err := DeleteDeviceInfo(context.Background(), info.ID)
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

	t.Run("createDeviceInfo", createDeviceInfo)
	t.Run("createDeviceInfos", createDeviceInfos)
	t.Run("getDeviceInfo", getDeviceInfo)
	t.Run("getDeviceInfos", getDeviceInfos)
	t.Run("getDeviceInfoOnly", getDeviceInfoOnly)
	t.Run("updateDeviceInfo", updateDeviceInfo)
	t.Run("existDeviceInfo", existDeviceInfo)
	t.Run("existDeviceInfoConds", existDeviceInfoConds)
	t.Run("delete", deleteDeviceInfo)
}
