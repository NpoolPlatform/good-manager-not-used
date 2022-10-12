package vendorlocation

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

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/vendorlocation"

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

var appDate = npool.VendorLocation{
	ID:       uuid.NewString(),
	Country:  uuid.NewString(),
	Province: uuid.NewString(),
	City:     uuid.NewString(),
	Address:  uuid.NewString(),
}

var (
	appInfo = npool.VendorLocationReq{
		ID:       &appDate.ID,
		Country:  &appDate.Country,
		Province: &appDate.Province,
		City:     &appDate.City,
		Address:  &appDate.Address,
	}
)

var info *npool.VendorLocation

func createVendorLocation(t *testing.T) {
	var err error
	info, err = CreateVendorLocation(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createVendorLocations(t *testing.T) {
	appDates := []npool.VendorLocation{
		{
			ID:       uuid.NewString(),
			Country:  uuid.NewString(),
			Province: uuid.NewString(),
			City:     uuid.NewString(),
			Address:  uuid.NewString(),
		},
		{
			ID:       uuid.NewString(),
			Country:  uuid.NewString(),
			Province: uuid.NewString(),
			City:     uuid.NewString(),
			Address:  uuid.NewString(),
		},
	}

	apps := []*npool.VendorLocationReq{}
	for key := range appDates {
		apps = append(apps, &npool.VendorLocationReq{
			ID:       &appDates[key].ID,
			Country:  &appDates[key].Country,
			Province: &appDates[key].Province,
			City:     &appDates[key].City,
			Address:  &appDates[key].Address,
		})
	}

	infos, err := CreateVendorLocations(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateVendorLocation(t *testing.T) {
	var err error
	info, err = UpdateVendorLocation(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getVendorLocation(t *testing.T) {
	var err error
	info, err = GetVendorLocation(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getVendorLocations(t *testing.T) {
	infos, total, err := GetVendorLocations(context.Background(),
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

func getVendorLocationOnly(t *testing.T) {
	var err error
	info, err = GetVendorLocationOnly(context.Background(),
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

func existVendorLocation(t *testing.T) {
	exist, err := ExistVendorLocation(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existVendorLocationConds(t *testing.T) {
	exist, err := ExistVendorLocationConds(context.Background(),
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

func deleteVendorLocation(t *testing.T) {
	info, err := DeleteVendorLocation(context.Background(), info.ID)
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

	t.Run("createVendorLocation", createVendorLocation)
	t.Run("createVendorLocations", createVendorLocations)
	t.Run("getVendorLocation", getVendorLocation)
	t.Run("getVendorLocations", getVendorLocations)
	t.Run("getVendorLocationOnly", getVendorLocationOnly)
	t.Run("updateVendorLocation", updateVendorLocation)
	t.Run("existVendorLocation", existVendorLocation)
	t.Run("existVendorLocationConds", existVendorLocationConds)
	t.Run("delete", deleteVendorLocation)
}
