package good

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/good"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	"github.com/google/uuid"
)

//nolint
func CreateSet(c *ent.GoodCreate, in *npool.GoodReq) (*ent.GoodCreate, error) {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.DeviceInfoID != nil {
		c.SetDeviceInfoID(uuid.MustParse(in.GetDeviceInfoID()))
	}
	if in.DurationDays != nil {
		c.SetDurationDays(in.GetDurationDays())
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.InheritFromGoodID != nil {
		c.SetInheritFromGoodID(uuid.MustParse(in.GetInheritFromGoodID()))
	}
	if in.VendorLocationID != nil {
		c.SetVendorLocationID(uuid.MustParse(in.GetVendorLocationID()))
	}
	if in.Price != nil {
		price, err := decimal.NewFromString(in.GetPrice())
		if err != nil {
			return nil, err
		}
		c.SetPrice(price)
	}
	if in.BenefitType != nil {
		c.SetBenefitType(in.GetBenefitType().String())
	}
	if in.GoodType != nil {
		c.SetGoodType(in.GetGoodType().String())
	}
	if in.Title != nil {
		c.SetTitle(in.GetTitle())
	}
	if in.Unit != nil {
		c.SetUnit(in.GetUnit())
	}
	if in.UnitAmount != nil {
		c.SetUnitAmount(in.GetUnitAmount())
	}
	if in.SupportCoinTypeIDs != nil {
		supportCoinTypeIDs := []uuid.UUID{}
		for _, val := range in.GetSupportCoinTypeIDs() {
			supportCoinTypeIDs = append(supportCoinTypeIDs, uuid.MustParse(val))
		}
		c.SetSupportCoinTypeIds(supportCoinTypeIDs)
	}
	if in.DeliveryAt != nil {
		c.SetDeliveryAt(in.GetDeliveryAt())
	}
	if in.StartAt != nil {
		c.SetStartAt(in.GetStartAt())
	}
	if in.TestOnly != nil {
		c.SetTestOnly(in.GetTestOnly())
	}
	if in.BenefitIntervalHours != nil {
		c.SetBenefitIntervalHours(in.GetBenefitIntervalHours())
	}
	c.SetBenefitState(npool.BenefitState_BenefitWait.String())
	c.SetLastBenefitAt(0)
	c.SetBenefitTids([]uuid.UUID{})
	return c, nil
}

func Create(ctx context.Context, in *npool.GoodReq) (*ent.Good, error) {
	var info *ent.Good
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.Good.Create()
		stm, err := CreateSet(c, in)
		if err != nil {
			return err
		}
		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.GoodReq) ([]*ent.Good, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Good{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.GoodCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.Good.Create()
			bulk[i], err = CreateSet(bulk[i], info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.Good.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//nolint
func UpdateSet(info *ent.Good, in *npool.GoodReq) (*ent.GoodUpdateOne, error) {
	u := info.Update()

	if in.DeviceInfoID != nil {
		u.SetDeviceInfoID(uuid.MustParse(in.GetDeviceInfoID()))
	}
	if in.DurationDays != nil {
		u.SetDurationDays(in.GetDurationDays())
	}
	if in.CoinTypeID != nil {
		u.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.InheritFromGoodID != nil {
		u.SetInheritFromGoodID(uuid.MustParse(in.GetInheritFromGoodID()))
	}
	if in.VendorLocationID != nil {
		u.SetVendorLocationID(uuid.MustParse(in.GetVendorLocationID()))
	}
	if in.Price != nil {
		price, err := decimal.NewFromString(in.GetPrice())
		if err != nil {
			return nil, err
		}
		u.SetPrice(price)
	}
	if in.Title != nil {
		u.SetTitle(in.GetTitle())
	}
	if in.Unit != nil {
		u.SetUnit(in.GetUnit())
	}
	if in.UnitAmount != nil {
		u.SetUnitAmount(in.GetUnitAmount())
	}
	if in.SupportCoinTypeIDs != nil {
		supportCoinTypeIDs := []uuid.UUID{}
		for _, val := range in.GetSupportCoinTypeIDs() {
			supportCoinTypeIDs = append(supportCoinTypeIDs, uuid.MustParse(val))
		}
		u.SetSupportCoinTypeIds(supportCoinTypeIDs)
	}
	if in.DeliveryAt != nil {
		u.SetDeliveryAt(in.GetDeliveryAt())
	}
	if in.StartAt != nil {
		u.SetStartAt(in.GetStartAt())
	}
	if in.TestOnly != nil {
		u.SetTestOnly(in.GetTestOnly())
	}
	if in.BenefitIntervalHours != nil {
		u.SetBenefitIntervalHours(in.GetBenefitIntervalHours())
	}
	if in.BenefitState != nil {
		u.SetBenefitState(in.GetBenefitState().String())
		if info.BenefitState != npool.BenefitState_BenefitWait.String() {
			if in.GetBenefitState() == npool.BenefitState_BenefitWait {
				u.SetLastBenefitAt(uint32(time.Now().Unix()))
			}
		}
	}
	if len(in.GetBenefitTIDs()) > 0 {
		ids := []uuid.UUID{}
		for _, id := range in.GetBenefitTIDs() {
			ids = append(ids, uuid.MustParse(id))
		}
		u.SetBenefitTids(ids)
	}
	return u, nil
}

func Update(ctx context.Context, in *npool.GoodReq) (*ent.Good, error) {
	var info *ent.Good
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Good.Query().Where(good.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := UpdateSet(info, in)
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Good, error) {
	var info *ent.Good
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Good.Query().Where(good.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.GoodQuery, error) {
	stm := cli.Good.Query()
	if conds == nil {
		return stm, nil
	}
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(good.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.DeviceInfoID != nil {
		switch conds.GetDeviceInfoID().GetOp() {
		case cruder.EQ:
			stm.Where(good.DeviceInfoID(uuid.MustParse(conds.GetDeviceInfoID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(good.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.VendorLocationID != nil {
		switch conds.GetVendorLocationID().GetOp() {
		case cruder.EQ:
			stm.Where(good.VendorLocationID(uuid.MustParse(conds.GetVendorLocationID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.BenefitType != nil {
		switch conds.GetBenefitType().GetOp() {
		case cruder.EQ:
			stm.Where(good.BenefitType(npool.BenefitType(conds.GetBenefitType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.GoodType != nil {
		switch conds.GetGoodType().GetOp() {
		case cruder.EQ:
			stm.Where(good.GoodType(npool.GoodType(conds.GetGoodType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.IDs != nil {
		switch conds.GetIDs().GetOp() {
		case cruder.IN:
			ids := []uuid.UUID{}
			for _, val := range conds.GetIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}

			stm.Where(good.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.BenefitState != nil {
		switch conds.GetBenefitState().GetOp() {
		case cruder.EQ:
			stm.Where(good.BenefitState(npool.BenefitState(conds.GetBenefitState().GetValue()).String()))
		case cruder.NEQ:
			stm.Where(good.BenefitState(npool.BenefitState(conds.GetBenefitState().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Good, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Good{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(good.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Good, error) {
	var info *ent.Good
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Good.Query().Where(good.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id string) (*ent.Good, error) {
	var info *ent.Good
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Good.UpdateOneID(uuid.MustParse(id)).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
