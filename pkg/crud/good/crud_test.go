package good

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

var ret = ent.Good{
	ID:                   uuid.New(),
	DeviceInfoID:         uuid.New(),
	DurationDays:         100,
	CoinTypeID:           uuid.New(),
	InheritFromGoodID:    uuid.New(),
	VendorLocationID:     uuid.New(),
	Price:                decimal.RequireFromString("9999999999999999999.999999999999999999"),
	BenefitType:          npool.BenefitType_BenefitTypePlatform.String(),
	GoodType:             npool.GoodType_GoodTypeClassicMining.String(),
	Title:                uuid.NewString(),
	Unit:                 uuid.NewString(),
	UnitAmount:           100,
	SupportCoinTypeIds:   []uuid.UUID{uuid.New()},
	DeliveryAt:           100,
	StartAt:              100,
	TestOnly:             true,
	BenefitIntervalHours: 24,
	BenefitState:         npool.BenefitState_BenefitWait.String(),
}

var (
	id                 = ret.ID.String()
	retID              = ret.DeviceInfoID.String()
	coinTypeID         = ret.CoinTypeID.String()
	inheritFromGoodID  = ret.InheritFromGoodID.String()
	vendorLocationID   = ret.VendorLocationID.String()
	price              = ret.Price.String()
	benefitType        = npool.BenefitType_BenefitTypePlatform
	goodType           = npool.GoodType_GoodTypeClassicMining
	supportCoinTypeIDs = []string{ret.SupportCoinTypeIds[0].String()}
	req                = npool.GoodReq{
		ID:                 &id,
		DeviceInfoID:       &retID,
		DurationDays:       &ret.DurationDays,
		CoinTypeID:         &coinTypeID,
		InheritFromGoodID:  &inheritFromGoodID,
		VendorLocationID:   &vendorLocationID,
		Price:              &price,
		BenefitType:        &benefitType,
		GoodType:           &goodType,
		Title:              &ret.Title,
		Unit:               &ret.Unit,
		UnitAmount:         &ret.UnitAmount,
		SupportCoinTypeIDs: supportCoinTypeIDs,
		DeliveryAt:         &ret.DeliveryAt,
		StartAt:            &ret.StartAt,
		TestOnly:           &ret.TestOnly,
	}
)

var info *ent.Good

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.Good{
		{
			ID:                   uuid.New(),
			DeviceInfoID:         uuid.New(),
			DurationDays:         100,
			CoinTypeID:           uuid.New(),
			InheritFromGoodID:    uuid.New(),
			VendorLocationID:     uuid.New(),
			Price:                decimal.RequireFromString("9999999999999999999.999999999999999999"),
			BenefitType:          npool.BenefitType_BenefitTypePlatform.String(),
			GoodType:             npool.GoodType_GoodTypeClassicMining.String(),
			Title:                uuid.NewString(),
			Unit:                 uuid.NewString(),
			UnitAmount:           100,
			SupportCoinTypeIds:   []uuid.UUID{uuid.New()},
			DeliveryAt:           100,
			StartAt:              100,
			TestOnly:             true,
			BenefitIntervalHours: 24,
			BenefitState:         npool.BenefitState_BenefitWait.String(),
		},
		{
			ID:                   uuid.New(),
			DeviceInfoID:         uuid.New(),
			DurationDays:         100,
			CoinTypeID:           uuid.New(),
			InheritFromGoodID:    uuid.New(),
			VendorLocationID:     uuid.New(),
			Price:                decimal.RequireFromString("9999999999999999999.999999999999999999"),
			BenefitType:          npool.BenefitType_BenefitTypePlatform.String(),
			GoodType:             npool.GoodType_GoodTypeClassicMining.String(),
			Title:                uuid.NewString(),
			Unit:                 uuid.NewString(),
			UnitAmount:           100,
			SupportCoinTypeIds:   []uuid.UUID{uuid.New()},
			DeliveryAt:           100,
			StartAt:              100,
			TestOnly:             true,
			BenefitIntervalHours: 24,
			BenefitState:         npool.BenefitState_BenefitWait.String(),
		},
	}

	reqs := []*npool.GoodReq{}
	for _, _ret := range entities {
		_id := _ret.ID.String()
		_retID := _ret.DeviceInfoID.String()
		_coinTypeID := _ret.CoinTypeID.String()
		_inheritFromGoodID := _ret.InheritFromGoodID.String()
		_vendorLocationID := _ret.VendorLocationID.String()
		_price := _ret.Price.String()
		_benefitType := npool.BenefitType_BenefitTypePlatform
		_goodType := npool.GoodType_GoodTypeClassicMining
		_supportCoinTypeIDs := []string{uuid.NewString()}
		reqs = append(reqs, &npool.GoodReq{
			ID:                 &_id,
			DeviceInfoID:       &_retID,
			DurationDays:       &ret.DurationDays,
			CoinTypeID:         &_coinTypeID,
			InheritFromGoodID:  &_inheritFromGoodID,
			VendorLocationID:   &_vendorLocationID,
			Price:              &_price,
			BenefitType:        &_benefitType,
			GoodType:           &_goodType,
			Title:              &ret.Title,
			Unit:               &ret.Unit,
			UnitAmount:         &ret.UnitAmount,
			SupportCoinTypeIDs: _supportCoinTypeIDs,
			DeliveryAt:         &ret.DeliveryAt,
			StartAt:            &ret.StartAt,
			TestOnly:           &ret.TestOnly,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	var err error
	hours := uint32(2)

	req.BenefitIntervalHours = &hours
	ret.BenefitIntervalHours = hours

	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}

	state := npool.BenefitState_BenefitWait
	req.BenefitState = &state

	info, err = Update(context.Background(), &req)
	assert.NotNil(t, err)

	state = npool.BenefitState_BenefitTransferring
	req.BenefitState = &state
	ret.BenefitState = state.String()

	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}

	state = npool.BenefitState_BenefitBookKeeping
	req.BenefitState = &state

	info, err = Update(context.Background(), &req)
	assert.NotNil(t, err)
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
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
			assert.Equal(t, infos[0].String(), ret.String())
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
		assert.Equal(t, info.String(), ret.String())
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
	exist, err := Exist(context.Background(), ret.ID)
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
	info, err := Delete(context.Background(), ret.ID.String())
	if assert.Nil(t, err) {
		ret.DeletedAt = info.DeletedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
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
