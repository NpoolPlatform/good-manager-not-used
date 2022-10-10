// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/deviceinfo"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// DeviceInfoQuery is the builder for querying DeviceInfo entities.
type DeviceInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeviceInfo
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceInfoQuery builder.
func (diq *DeviceInfoQuery) Where(ps ...predicate.DeviceInfo) *DeviceInfoQuery {
	diq.predicates = append(diq.predicates, ps...)
	return diq
}

// Limit adds a limit step to the query.
func (diq *DeviceInfoQuery) Limit(limit int) *DeviceInfoQuery {
	diq.limit = &limit
	return diq
}

// Offset adds an offset step to the query.
func (diq *DeviceInfoQuery) Offset(offset int) *DeviceInfoQuery {
	diq.offset = &offset
	return diq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (diq *DeviceInfoQuery) Unique(unique bool) *DeviceInfoQuery {
	diq.unique = &unique
	return diq
}

// Order adds an order step to the query.
func (diq *DeviceInfoQuery) Order(o ...OrderFunc) *DeviceInfoQuery {
	diq.order = append(diq.order, o...)
	return diq
}

// First returns the first DeviceInfo entity from the query.
// Returns a *NotFoundError when no DeviceInfo was found.
func (diq *DeviceInfoQuery) First(ctx context.Context) (*DeviceInfo, error) {
	nodes, err := diq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deviceinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (diq *DeviceInfoQuery) FirstX(ctx context.Context) *DeviceInfo {
	node, err := diq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeviceInfo ID from the query.
// Returns a *NotFoundError when no DeviceInfo ID was found.
func (diq *DeviceInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = diq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deviceinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (diq *DeviceInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := diq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeviceInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeviceInfo entity is found.
// Returns a *NotFoundError when no DeviceInfo entities are found.
func (diq *DeviceInfoQuery) Only(ctx context.Context) (*DeviceInfo, error) {
	nodes, err := diq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deviceinfo.Label}
	default:
		return nil, &NotSingularError{deviceinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (diq *DeviceInfoQuery) OnlyX(ctx context.Context) *DeviceInfo {
	node, err := diq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeviceInfo ID in the query.
// Returns a *NotSingularError when more than one DeviceInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (diq *DeviceInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = diq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deviceinfo.Label}
	default:
		err = &NotSingularError{deviceinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (diq *DeviceInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := diq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeviceInfos.
func (diq *DeviceInfoQuery) All(ctx context.Context) ([]*DeviceInfo, error) {
	if err := diq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return diq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (diq *DeviceInfoQuery) AllX(ctx context.Context) []*DeviceInfo {
	nodes, err := diq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeviceInfo IDs.
func (diq *DeviceInfoQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := diq.Select(deviceinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (diq *DeviceInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := diq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (diq *DeviceInfoQuery) Count(ctx context.Context) (int, error) {
	if err := diq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return diq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (diq *DeviceInfoQuery) CountX(ctx context.Context) int {
	count, err := diq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (diq *DeviceInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := diq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return diq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (diq *DeviceInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := diq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (diq *DeviceInfoQuery) Clone() *DeviceInfoQuery {
	if diq == nil {
		return nil
	}
	return &DeviceInfoQuery{
		config:     diq.config,
		limit:      diq.limit,
		offset:     diq.offset,
		order:      append([]OrderFunc{}, diq.order...),
		predicates: append([]predicate.DeviceInfo{}, diq.predicates...),
		// clone intermediate query.
		sql:    diq.sql.Clone(),
		path:   diq.path,
		unique: diq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeviceInfo.Query().
//		GroupBy(deviceinfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (diq *DeviceInfoQuery) GroupBy(field string, fields ...string) *DeviceInfoGroupBy {
	grbuild := &DeviceInfoGroupBy{config: diq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := diq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return diq.sqlQuery(ctx), nil
	}
	grbuild.label = deviceinfo.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.DeviceInfo.Query().
//		Select(deviceinfo.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (diq *DeviceInfoQuery) Select(fields ...string) *DeviceInfoSelect {
	diq.fields = append(diq.fields, fields...)
	selbuild := &DeviceInfoSelect{DeviceInfoQuery: diq}
	selbuild.label = deviceinfo.Label
	selbuild.flds, selbuild.scan = &diq.fields, selbuild.Scan
	return selbuild
}

func (diq *DeviceInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range diq.fields {
		if !deviceinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if diq.path != nil {
		prev, err := diq.path(ctx)
		if err != nil {
			return err
		}
		diq.sql = prev
	}
	if deviceinfo.Policy == nil {
		return errors.New("ent: uninitialized deviceinfo.Policy (forgotten import ent/runtime?)")
	}
	if err := deviceinfo.Policy.EvalQuery(ctx, diq); err != nil {
		return err
	}
	return nil
}

func (diq *DeviceInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DeviceInfo, error) {
	var (
		nodes = []*DeviceInfo{}
		_spec = diq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DeviceInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DeviceInfo{config: diq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(diq.modifiers) > 0 {
		_spec.Modifiers = diq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, diq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (diq *DeviceInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := diq.querySpec()
	if len(diq.modifiers) > 0 {
		_spec.Modifiers = diq.modifiers
	}
	_spec.Node.Columns = diq.fields
	if len(diq.fields) > 0 {
		_spec.Unique = diq.unique != nil && *diq.unique
	}
	return sqlgraph.CountNodes(ctx, diq.driver, _spec)
}

func (diq *DeviceInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := diq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (diq *DeviceInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deviceinfo.Table,
			Columns: deviceinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deviceinfo.FieldID,
			},
		},
		From:   diq.sql,
		Unique: true,
	}
	if unique := diq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := diq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deviceinfo.FieldID)
		for i := range fields {
			if fields[i] != deviceinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := diq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := diq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := diq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := diq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (diq *DeviceInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(diq.driver.Dialect())
	t1 := builder.Table(deviceinfo.Table)
	columns := diq.fields
	if len(columns) == 0 {
		columns = deviceinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if diq.sql != nil {
		selector = diq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if diq.unique != nil && *diq.unique {
		selector.Distinct()
	}
	for _, m := range diq.modifiers {
		m(selector)
	}
	for _, p := range diq.predicates {
		p(selector)
	}
	for _, p := range diq.order {
		p(selector)
	}
	if offset := diq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := diq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (diq *DeviceInfoQuery) ForUpdate(opts ...sql.LockOption) *DeviceInfoQuery {
	if diq.driver.Dialect() == dialect.Postgres {
		diq.Unique(false)
	}
	diq.modifiers = append(diq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return diq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (diq *DeviceInfoQuery) ForShare(opts ...sql.LockOption) *DeviceInfoQuery {
	if diq.driver.Dialect() == dialect.Postgres {
		diq.Unique(false)
	}
	diq.modifiers = append(diq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return diq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (diq *DeviceInfoQuery) Modify(modifiers ...func(s *sql.Selector)) *DeviceInfoSelect {
	diq.modifiers = append(diq.modifiers, modifiers...)
	return diq.Select()
}

// DeviceInfoGroupBy is the group-by builder for DeviceInfo entities.
type DeviceInfoGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (digb *DeviceInfoGroupBy) Aggregate(fns ...AggregateFunc) *DeviceInfoGroupBy {
	digb.fns = append(digb.fns, fns...)
	return digb
}

// Scan applies the group-by query and scans the result into the given value.
func (digb *DeviceInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := digb.path(ctx)
	if err != nil {
		return err
	}
	digb.sql = query
	return digb.sqlScan(ctx, v)
}

func (digb *DeviceInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range digb.fields {
		if !deviceinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := digb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := digb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (digb *DeviceInfoGroupBy) sqlQuery() *sql.Selector {
	selector := digb.sql.Select()
	aggregation := make([]string, 0, len(digb.fns))
	for _, fn := range digb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(digb.fields)+len(digb.fns))
		for _, f := range digb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(digb.fields...)...)
}

// DeviceInfoSelect is the builder for selecting fields of DeviceInfo entities.
type DeviceInfoSelect struct {
	*DeviceInfoQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dis *DeviceInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dis.prepareQuery(ctx); err != nil {
		return err
	}
	dis.sql = dis.DeviceInfoQuery.sqlQuery(ctx)
	return dis.sqlScan(ctx, v)
}

func (dis *DeviceInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dis.sql.Query()
	if err := dis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dis *DeviceInfoSelect) Modify(modifiers ...func(s *sql.Selector)) *DeviceInfoSelect {
	dis.modifiers = append(dis.modifiers, modifiers...)
	return dis
}
