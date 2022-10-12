package good

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

	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

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

var appDate = npool.Good{
	ID:                 uuid.NewString(),
	DeviceInfoID:       uuid.NewString(),
	DurationDays:       1001,
	CoinTypeID:         uuid.NewString(),
	InheritFromGoodID:  uuid.NewString(),
	VendorLocationID:   uuid.NewString(),
	Price:              "9999999999999999999.999999999999999999",
	BenefitType:        npool.BenefitType_BenefitTypePlatform,
	GoodType:           npool.GoodType_GoodTypeElectricityFee,
	Title:              uuid.NewString(),
	Unit:               uuid.NewString(),
	UnitAmount:         1004,
	SupportCoinTypeIDs: []string{uuid.NewString()},
	DeliveryAt:         1005,
	StartAt:            1006,
	TestOnly:           true,
}

var (
	appInfo = npool.GoodReq{
		ID:                 &appDate.ID,
		DeviceInfoID:       &appDate.DeviceInfoID,
		DurationDays:       &appDate.DurationDays,
		CoinTypeID:         &appDate.CoinTypeID,
		InheritFromGoodID:  &appDate.InheritFromGoodID,
		VendorLocationID:   &appDate.VendorLocationID,
		Price:              &appDate.Price,
		BenefitType:        &appDate.BenefitType,
		GoodType:           &appDate.GoodType,
		Title:              &appDate.Title,
		Unit:               &appDate.Unit,
		UnitAmount:         &appDate.UnitAmount,
		SupportCoinTypeIDs: appDate.SupportCoinTypeIDs,
		DeliveryAt:         &appDate.DeliveryAt,
		StartAt:            &appDate.StartAt,
		TestOnly:           &appDate.TestOnly,
	}
)

var info *npool.Good

func createGood(t *testing.T) {
	var err error
	info, err = CreateGood(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createGoods(t *testing.T) {
	appDates := []npool.Good{
		{
			ID:                 uuid.NewString(),
			DeviceInfoID:       uuid.NewString(),
			DurationDays:       1001,
			CoinTypeID:         uuid.NewString(),
			InheritFromGoodID:  uuid.NewString(),
			VendorLocationID:   uuid.NewString(),
			Price:              "9999999999999999999.999999999999999999",
			BenefitType:        npool.BenefitType_BenefitTypePlatform,
			GoodType:           npool.GoodType_GoodTypeElectricityFee,
			Title:              uuid.NewString(),
			Unit:               uuid.NewString(),
			UnitAmount:         1004,
			SupportCoinTypeIDs: []string{uuid.NewString()},
			DeliveryAt:         1005,
			StartAt:            1006,
			TestOnly:           true,
		},
		{
			ID:                 uuid.NewString(),
			DeviceInfoID:       uuid.NewString(),
			DurationDays:       1001,
			CoinTypeID:         uuid.NewString(),
			InheritFromGoodID:  uuid.NewString(),
			VendorLocationID:   uuid.NewString(),
			Price:              "9999999999999999999.999999999999999999",
			BenefitType:        npool.BenefitType_BenefitTypePlatform,
			GoodType:           npool.GoodType_GoodTypeElectricityFee,
			Title:              uuid.NewString(),
			Unit:               uuid.NewString(),
			UnitAmount:         1004,
			SupportCoinTypeIDs: []string{uuid.NewString()},
			DeliveryAt:         1005,
			StartAt:            1006,
			TestOnly:           true,
		},
	}

	apps := []*npool.GoodReq{}
	for key := range appDates {
		apps = append(apps, &npool.GoodReq{
			ID:                 &appDates[key].ID,
			DeviceInfoID:       &appDates[key].DeviceInfoID,
			DurationDays:       &appDates[key].DurationDays,
			CoinTypeID:         &appDates[key].CoinTypeID,
			InheritFromGoodID:  &appDates[key].InheritFromGoodID,
			VendorLocationID:   &appDates[key].VendorLocationID,
			Price:              &appDates[key].Price,
			BenefitType:        &appDates[key].BenefitType,
			GoodType:           &appDates[key].GoodType,
			Title:              &appDates[key].Title,
			Unit:               &appDates[key].Unit,
			UnitAmount:         &appDates[key].UnitAmount,
			SupportCoinTypeIDs: appDates[key].SupportCoinTypeIDs,
			DeliveryAt:         &appDates[key].DeliveryAt,
			StartAt:            &appDates[key].StartAt,
			TestOnly:           &appDates[key].TestOnly,
		})
	}

	infos, err := CreateGoods(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateGood(t *testing.T) {
	var err error
	info, err = UpdateGood(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getGood(t *testing.T) {
	var err error
	info, err = GetGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getGoods(t *testing.T) {
	infos, total, err := GetGoods(context.Background(),
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

func getGoodOnly(t *testing.T) {
	var err error
	info, err = GetGoodOnly(context.Background(),
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

func existGood(t *testing.T) {
	exist, err := ExistGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existGoodConds(t *testing.T) {
	exist, err := ExistGoodConds(context.Background(),
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

func deleteGood(t *testing.T) {
	info, err := DeleteGood(context.Background(), info.ID)
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

	t.Run("createGood", createGood)
	t.Run("createGoods", createGoods)
	t.Run("getGood", getGood)
	t.Run("getGoods", getGoods)
	t.Run("getGoodOnly", getGoodOnly)
	t.Run("updateGood", updateGood)
	t.Run("existGood", existGood)
	t.Run("existGoodConds", existGoodConds)
	t.Run("delete", deleteGood)
}
