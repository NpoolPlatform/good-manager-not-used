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
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/extrainfo"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ExtraInfoQuery is the builder for querying ExtraInfo entities.
type ExtraInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ExtraInfo
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExtraInfoQuery builder.
func (eiq *ExtraInfoQuery) Where(ps ...predicate.ExtraInfo) *ExtraInfoQuery {
	eiq.predicates = append(eiq.predicates, ps...)
	return eiq
}

// Limit adds a limit step to the query.
func (eiq *ExtraInfoQuery) Limit(limit int) *ExtraInfoQuery {
	eiq.limit = &limit
	return eiq
}

// Offset adds an offset step to the query.
func (eiq *ExtraInfoQuery) Offset(offset int) *ExtraInfoQuery {
	eiq.offset = &offset
	return eiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eiq *ExtraInfoQuery) Unique(unique bool) *ExtraInfoQuery {
	eiq.unique = &unique
	return eiq
}

// Order adds an order step to the query.
func (eiq *ExtraInfoQuery) Order(o ...OrderFunc) *ExtraInfoQuery {
	eiq.order = append(eiq.order, o...)
	return eiq
}

// First returns the first ExtraInfo entity from the query.
// Returns a *NotFoundError when no ExtraInfo was found.
func (eiq *ExtraInfoQuery) First(ctx context.Context) (*ExtraInfo, error) {
	nodes, err := eiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{extrainfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eiq *ExtraInfoQuery) FirstX(ctx context.Context) *ExtraInfo {
	node, err := eiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExtraInfo ID from the query.
// Returns a *NotFoundError when no ExtraInfo ID was found.
func (eiq *ExtraInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{extrainfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eiq *ExtraInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := eiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExtraInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExtraInfo entity is found.
// Returns a *NotFoundError when no ExtraInfo entities are found.
func (eiq *ExtraInfoQuery) Only(ctx context.Context) (*ExtraInfo, error) {
	nodes, err := eiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{extrainfo.Label}
	default:
		return nil, &NotSingularError{extrainfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eiq *ExtraInfoQuery) OnlyX(ctx context.Context) *ExtraInfo {
	node, err := eiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExtraInfo ID in the query.
// Returns a *NotSingularError when more than one ExtraInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (eiq *ExtraInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{extrainfo.Label}
	default:
		err = &NotSingularError{extrainfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eiq *ExtraInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := eiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExtraInfos.
func (eiq *ExtraInfoQuery) All(ctx context.Context) ([]*ExtraInfo, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eiq *ExtraInfoQuery) AllX(ctx context.Context) []*ExtraInfo {
	nodes, err := eiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExtraInfo IDs.
func (eiq *ExtraInfoQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := eiq.Select(extrainfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eiq *ExtraInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := eiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eiq *ExtraInfoQuery) Count(ctx context.Context) (int, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eiq *ExtraInfoQuery) CountX(ctx context.Context) int {
	count, err := eiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eiq *ExtraInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eiq *ExtraInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := eiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExtraInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eiq *ExtraInfoQuery) Clone() *ExtraInfoQuery {
	if eiq == nil {
		return nil
	}
	return &ExtraInfoQuery{
		config:     eiq.config,
		limit:      eiq.limit,
		offset:     eiq.offset,
		order:      append([]OrderFunc{}, eiq.order...),
		predicates: append([]predicate.ExtraInfo{}, eiq.predicates...),
		// clone intermediate query.
		sql:    eiq.sql.Clone(),
		path:   eiq.path,
		unique: eiq.unique,
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
//	client.ExtraInfo.Query().
//		GroupBy(extrainfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (eiq *ExtraInfoQuery) GroupBy(field string, fields ...string) *ExtraInfoGroupBy {
	grbuild := &ExtraInfoGroupBy{config: eiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eiq.sqlQuery(ctx), nil
	}
	grbuild.label = extrainfo.Label
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
//	client.ExtraInfo.Query().
//		Select(extrainfo.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (eiq *ExtraInfoQuery) Select(fields ...string) *ExtraInfoSelect {
	eiq.fields = append(eiq.fields, fields...)
	selbuild := &ExtraInfoSelect{ExtraInfoQuery: eiq}
	selbuild.label = extrainfo.Label
	selbuild.flds, selbuild.scan = &eiq.fields, selbuild.Scan
	return selbuild
}

func (eiq *ExtraInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eiq.fields {
		if !extrainfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eiq.path != nil {
		prev, err := eiq.path(ctx)
		if err != nil {
			return err
		}
		eiq.sql = prev
	}
	if extrainfo.Policy == nil {
		return errors.New("ent: uninitialized extrainfo.Policy (forgotten import ent/runtime?)")
	}
	if err := extrainfo.Policy.EvalQuery(ctx, eiq); err != nil {
		return err
	}
	return nil
}

func (eiq *ExtraInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExtraInfo, error) {
	var (
		nodes = []*ExtraInfo{}
		_spec = eiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ExtraInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ExtraInfo{config: eiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(eiq.modifiers) > 0 {
		_spec.Modifiers = eiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (eiq *ExtraInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eiq.querySpec()
	if len(eiq.modifiers) > 0 {
		_spec.Modifiers = eiq.modifiers
	}
	_spec.Node.Columns = eiq.fields
	if len(eiq.fields) > 0 {
		_spec.Unique = eiq.unique != nil && *eiq.unique
	}
	return sqlgraph.CountNodes(ctx, eiq.driver, _spec)
}

func (eiq *ExtraInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := eiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (eiq *ExtraInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   extrainfo.Table,
			Columns: extrainfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: extrainfo.FieldID,
			},
		},
		From:   eiq.sql,
		Unique: true,
	}
	if unique := eiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := eiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, extrainfo.FieldID)
		for i := range fields {
			if fields[i] != extrainfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eiq *ExtraInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eiq.driver.Dialect())
	t1 := builder.Table(extrainfo.Table)
	columns := eiq.fields
	if len(columns) == 0 {
		columns = extrainfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eiq.sql != nil {
		selector = eiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eiq.unique != nil && *eiq.unique {
		selector.Distinct()
	}
	for _, m := range eiq.modifiers {
		m(selector)
	}
	for _, p := range eiq.predicates {
		p(selector)
	}
	for _, p := range eiq.order {
		p(selector)
	}
	if offset := eiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (eiq *ExtraInfoQuery) ForUpdate(opts ...sql.LockOption) *ExtraInfoQuery {
	if eiq.driver.Dialect() == dialect.Postgres {
		eiq.Unique(false)
	}
	eiq.modifiers = append(eiq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return eiq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (eiq *ExtraInfoQuery) ForShare(opts ...sql.LockOption) *ExtraInfoQuery {
	if eiq.driver.Dialect() == dialect.Postgres {
		eiq.Unique(false)
	}
	eiq.modifiers = append(eiq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return eiq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (eiq *ExtraInfoQuery) Modify(modifiers ...func(s *sql.Selector)) *ExtraInfoSelect {
	eiq.modifiers = append(eiq.modifiers, modifiers...)
	return eiq.Select()
}

// ExtraInfoGroupBy is the group-by builder for ExtraInfo entities.
type ExtraInfoGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eigb *ExtraInfoGroupBy) Aggregate(fns ...AggregateFunc) *ExtraInfoGroupBy {
	eigb.fns = append(eigb.fns, fns...)
	return eigb
}

// Scan applies the group-by query and scans the result into the given value.
func (eigb *ExtraInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := eigb.path(ctx)
	if err != nil {
		return err
	}
	eigb.sql = query
	return eigb.sqlScan(ctx, v)
}

func (eigb *ExtraInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range eigb.fields {
		if !extrainfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := eigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (eigb *ExtraInfoGroupBy) sqlQuery() *sql.Selector {
	selector := eigb.sql.Select()
	aggregation := make([]string, 0, len(eigb.fns))
	for _, fn := range eigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(eigb.fields)+len(eigb.fns))
		for _, f := range eigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(eigb.fields...)...)
}

// ExtraInfoSelect is the builder for selecting fields of ExtraInfo entities.
type ExtraInfoSelect struct {
	*ExtraInfoQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (eis *ExtraInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := eis.prepareQuery(ctx); err != nil {
		return err
	}
	eis.sql = eis.ExtraInfoQuery.sqlQuery(ctx)
	return eis.sqlScan(ctx, v)
}

func (eis *ExtraInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := eis.sql.Query()
	if err := eis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (eis *ExtraInfoSelect) Modify(modifiers ...func(s *sql.Selector)) *ExtraInfoSelect {
	eis.modifiers = append(eis.modifiers, modifiers...)
	return eis
}