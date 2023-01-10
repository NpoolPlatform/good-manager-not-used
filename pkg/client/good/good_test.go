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

var ret = npool.Good{
	ID:                     uuid.NewString(),
	DeviceInfoID:           uuid.NewString(),
	DurationDays:           1001,
	CoinTypeID:             uuid.NewString(),
	InheritFromGoodID:      uuid.NewString(),
	VendorLocationID:       uuid.NewString(),
	Price:                  "9999999999999999999.999999999999999999",
	BenefitType:            npool.BenefitType_BenefitTypePlatform,
	GoodType:               npool.GoodType_GoodTypeElectricityFee,
	Title:                  uuid.NewString(),
	Unit:                   uuid.NewString(),
	UnitAmount:             1004,
	SupportCoinTypeIDs:     []string{uuid.NewString()},
	DeliveryAt:             1005,
	StartAt:                1006,
	TestOnly:               true,
	BenefitIntervalHours:   24,
	BenefitState:           npool.BenefitState_BenefitWait,
	NextBenefitStartAmount: "0",
}

var (
	req = npool.GoodReq{
		ID:                 &ret.ID,
		DeviceInfoID:       &ret.DeviceInfoID,
		DurationDays:       &ret.DurationDays,
		CoinTypeID:         &ret.CoinTypeID,
		InheritFromGoodID:  &ret.InheritFromGoodID,
		VendorLocationID:   &ret.VendorLocationID,
		Price:              &ret.Price,
		BenefitType:        &ret.BenefitType,
		GoodType:           &ret.GoodType,
		Title:              &ret.Title,
		Unit:               &ret.Unit,
		UnitAmount:         &ret.UnitAmount,
		SupportCoinTypeIDs: ret.SupportCoinTypeIDs,
		DeliveryAt:         &ret.DeliveryAt,
		StartAt:            &ret.StartAt,
		TestOnly:           &ret.TestOnly,
	}
)

var info *npool.Good

func createGood(t *testing.T) {
	var err error
	info, err = CreateGood(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func createGoods(t *testing.T) {
	rets := []npool.Good{
		{
			ID:                     uuid.NewString(),
			DeviceInfoID:           uuid.NewString(),
			DurationDays:           1001,
			CoinTypeID:             uuid.NewString(),
			InheritFromGoodID:      uuid.NewString(),
			VendorLocationID:       uuid.NewString(),
			Price:                  "9999999999999999999.999999999999999999",
			BenefitType:            npool.BenefitType_BenefitTypePlatform,
			GoodType:               npool.GoodType_GoodTypeElectricityFee,
			Title:                  uuid.NewString(),
			Unit:                   uuid.NewString(),
			UnitAmount:             1004,
			SupportCoinTypeIDs:     []string{uuid.NewString()},
			DeliveryAt:             1005,
			StartAt:                1006,
			TestOnly:               true,
			BenefitIntervalHours:   24,
			BenefitState:           npool.BenefitState_BenefitWait,
			NextBenefitStartAmount: "0",
		},
		{
			ID:                     uuid.NewString(),
			DeviceInfoID:           uuid.NewString(),
			DurationDays:           1001,
			CoinTypeID:             uuid.NewString(),
			InheritFromGoodID:      uuid.NewString(),
			VendorLocationID:       uuid.NewString(),
			Price:                  "9999999999999999999.999999999999999999",
			BenefitType:            npool.BenefitType_BenefitTypePlatform,
			GoodType:               npool.GoodType_GoodTypeElectricityFee,
			Title:                  uuid.NewString(),
			Unit:                   uuid.NewString(),
			UnitAmount:             1004,
			SupportCoinTypeIDs:     []string{uuid.NewString()},
			DeliveryAt:             1005,
			StartAt:                1006,
			TestOnly:               true,
			BenefitIntervalHours:   24,
			BenefitState:           npool.BenefitState_BenefitWait,
			NextBenefitStartAmount: "0",
		},
	}

	apps := []*npool.GoodReq{}
	for key := range rets {
		apps = append(apps, &npool.GoodReq{
			ID:                 &rets[key].ID,
			DeviceInfoID:       &rets[key].DeviceInfoID,
			DurationDays:       &rets[key].DurationDays,
			CoinTypeID:         &rets[key].CoinTypeID,
			InheritFromGoodID:  &rets[key].InheritFromGoodID,
			VendorLocationID:   &rets[key].VendorLocationID,
			Price:              &rets[key].Price,
			BenefitType:        &rets[key].BenefitType,
			GoodType:           &rets[key].GoodType,
			Title:              &rets[key].Title,
			Unit:               &rets[key].Unit,
			UnitAmount:         &rets[key].UnitAmount,
			SupportCoinTypeIDs: rets[key].SupportCoinTypeIDs,
			DeliveryAt:         &rets[key].DeliveryAt,
			StartAt:            &rets[key].StartAt,
			TestOnly:           &rets[key].TestOnly,
		})
	}

	infos, err := CreateGoods(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateGood(t *testing.T) {
	var err error
	hours := uint32(23)

	req.BenefitIntervalHours = &hours
	ret.BenefitIntervalHours = hours

	info, err = UpdateGood(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getGood(t *testing.T) {
	var err error
	info, err = GetGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
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
		assert.Equal(t, infos[0], &ret)
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
		assert.Equal(t, info, &ret)
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
