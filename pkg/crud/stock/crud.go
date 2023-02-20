package stock

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent/stock"
	"github.com/shopspring/decimal"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/good-manager/pkg/tracer/stock"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"

	"github.com/google/uuid"
)

func CreateSet(c *ent.StockCreate, in *npool.StockReq) (*ent.StockCreate, error) {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.GoodID != nil {
		c.SetGoodID(uuid.MustParse(in.GetGoodID()))
	}
	if in.Total != nil {
		val, err := decimal.NewFromString(in.GetTotal())
		if err != nil {
			return nil, err
		}
		c.SetTotal(val)
	}
	v := decimal.NewFromInt(0)
	c.SetLocked(v)
	c.SetInService(v)
	c.SetWaitStart(v)
	c.SetSold(v)
	return c, nil
}

func Create(ctx context.Context, in *npool.StockReq) (*ent.Stock, error) {
	var info *ent.Stock
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
		c := cli.Stock.Create()
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

func CreateBulk(ctx context.Context, in []*npool.StockReq) ([]*ent.Stock, error) {
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

	rows := []*ent.Stock{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.StockCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.Stock.Create()
			bulk[i], err = CreateSet(bulk[i], info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.Stock.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(u *ent.StockUpdateOne, in *npool.StockReq) (*ent.StockUpdateOne, error) {
	if in.Total != nil {
		val, err := decimal.NewFromString(in.GetTotal())
		if err != nil {
			return nil, err
		}
		u.SetTotal(val)
	}
	return u, nil
}

func Update(ctx context.Context, in *npool.StockReq) (*ent.Stock, error) {
	var info *ent.Stock
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
		info, err = tx.Stock.Query().Where(stock.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		total, err := decimal.NewFromString(in.GetTotal())
		if err != nil {
			return err
		}

		if total.Cmp(info.Locked.Add(info.InService).Add(info.WaitStart)) < 0 {
			return fmt.Errorf("stock insufficient")
		}

		stm, err := UpdateSet(info.Update(), in)
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

//nolint:gocyclo
func AddFieldSet(info *ent.Stock, in *npool.StockReq) (*ent.StockUpdateOne, error) {
	locked := info.Locked
	if in.Locked != nil {
		val, err := decimal.NewFromString(in.GetLocked())
		if err != nil {
			return nil, err
		}
		locked = locked.Add(val)
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("locked stock exhausted")
	}

	inService := info.InService
	if in.InService != nil {
		val, err := decimal.NewFromString(in.GetInService())
		if err != nil {
			return nil, err
		}
		inService = inService.Add(val)
	}
	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("in service stock exhausted")
	}

	waitStart := info.WaitStart
	if in.WaitStart != nil {
		val, err := decimal.NewFromString(in.GetWaitStart())
		if err != nil {
			return nil, err
		}
		waitStart = waitStart.Add(val)
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("in service stock exhausted")
	}

	if info.Total.Cmp(locked.Add(inService).Add(waitStart)) < 0 {
		return nil, fmt.Errorf("stock exhausted")
	}
	u := info.Update()
	if in.Locked != nil {
		u.SetLocked(locked)
	}
	if in.InService != nil {
		u.SetInService(inService)
	}
	if in.WaitStart != nil {
		u.SetWaitStart(waitStart)
		waitStart1, err := decimal.NewFromString(in.GetWaitStart())
		if err != nil {
			return nil, err
		}
		if waitStart1.Cmp(decimal.NewFromInt(0)) > 0 {
			u.SetSold(info.Sold.Add(waitStart1))
		}
	}
	return u, nil
}

func AddFields(ctx context.Context, in *npool.StockReq) (*ent.Stock, error) {
	var info *ent.Stock
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "AddFields")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Stock.Query().Where(stock.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query stock: %v", err)
		}

		stm, err := AddFieldSet(info, in)
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update stock: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update stock: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Stock, error) {
	var info *ent.Stock
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
		info, err = cli.Stock.Query().Where(stock.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.StockQuery, error) {
	stm := cli.Stock.Query()
	if conds == nil {
		return stm, nil
	}
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(stock.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid stock field")
		}
	}
	if conds.GoodID != nil {
		switch conds.GetGoodID().GetOp() {
		case cruder.EQ:
			stm.Where(stock.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid stock field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Stock, int, error) {
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

	rows := []*ent.Stock{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(stock.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Stock, error) {
	var info *ent.Stock
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
		stm, err := setQueryConds(conds, cli)
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
		stm, err := setQueryConds(conds, cli)
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
		exist, err = cli.Stock.Query().Where(stock.ID(id)).Exist(_ctx)
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
		stm, err := setQueryConds(conds, cli)
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

func Delete(ctx context.Context, id string) (*ent.Stock, error) {
	var info *ent.Stock
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
		info, err = cli.Stock.UpdateOneID(uuid.MustParse(id)).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
