// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/subgood"
	"github.com/google/uuid"
)

// SubGoodUpdate is the builder for updating SubGood entities.
type SubGoodUpdate struct {
	config
	hooks     []Hook
	mutation  *SubGoodMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SubGoodUpdate builder.
func (sgu *SubGoodUpdate) Where(ps ...predicate.SubGood) *SubGoodUpdate {
	sgu.mutation.Where(ps...)
	return sgu
}

// SetCreatedAt sets the "created_at" field.
func (sgu *SubGoodUpdate) SetCreatedAt(u uint32) *SubGoodUpdate {
	sgu.mutation.ResetCreatedAt()
	sgu.mutation.SetCreatedAt(u)
	return sgu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sgu *SubGoodUpdate) SetNillableCreatedAt(u *uint32) *SubGoodUpdate {
	if u != nil {
		sgu.SetCreatedAt(*u)
	}
	return sgu
}

// AddCreatedAt adds u to the "created_at" field.
func (sgu *SubGoodUpdate) AddCreatedAt(u int32) *SubGoodUpdate {
	sgu.mutation.AddCreatedAt(u)
	return sgu
}

// SetUpdatedAt sets the "updated_at" field.
func (sgu *SubGoodUpdate) SetUpdatedAt(u uint32) *SubGoodUpdate {
	sgu.mutation.ResetUpdatedAt()
	sgu.mutation.SetUpdatedAt(u)
	return sgu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (sgu *SubGoodUpdate) AddUpdatedAt(u int32) *SubGoodUpdate {
	sgu.mutation.AddUpdatedAt(u)
	return sgu
}

// SetDeletedAt sets the "deleted_at" field.
func (sgu *SubGoodUpdate) SetDeletedAt(u uint32) *SubGoodUpdate {
	sgu.mutation.ResetDeletedAt()
	sgu.mutation.SetDeletedAt(u)
	return sgu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sgu *SubGoodUpdate) SetNillableDeletedAt(u *uint32) *SubGoodUpdate {
	if u != nil {
		sgu.SetDeletedAt(*u)
	}
	return sgu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (sgu *SubGoodUpdate) AddDeletedAt(u int32) *SubGoodUpdate {
	sgu.mutation.AddDeletedAt(u)
	return sgu
}

// SetAppID sets the "app_id" field.
func (sgu *SubGoodUpdate) SetAppID(u uuid.UUID) *SubGoodUpdate {
	sgu.mutation.SetAppID(u)
	return sgu
}

// SetMainGoodID sets the "main_good_id" field.
func (sgu *SubGoodUpdate) SetMainGoodID(u uuid.UUID) *SubGoodUpdate {
	sgu.mutation.SetMainGoodID(u)
	return sgu
}

// SetSubGoodID sets the "sub_good_id" field.
func (sgu *SubGoodUpdate) SetSubGoodID(u uuid.UUID) *SubGoodUpdate {
	sgu.mutation.SetSubGoodID(u)
	return sgu
}

// SetMust sets the "must" field.
func (sgu *SubGoodUpdate) SetMust(b bool) *SubGoodUpdate {
	sgu.mutation.SetMust(b)
	return sgu
}

// SetNillableMust sets the "must" field if the given value is not nil.
func (sgu *SubGoodUpdate) SetNillableMust(b *bool) *SubGoodUpdate {
	if b != nil {
		sgu.SetMust(*b)
	}
	return sgu
}

// ClearMust clears the value of the "must" field.
func (sgu *SubGoodUpdate) ClearMust() *SubGoodUpdate {
	sgu.mutation.ClearMust()
	return sgu
}

// SetCommission sets the "commission" field.
func (sgu *SubGoodUpdate) SetCommission(b bool) *SubGoodUpdate {
	sgu.mutation.SetCommission(b)
	return sgu
}

// SetNillableCommission sets the "commission" field if the given value is not nil.
func (sgu *SubGoodUpdate) SetNillableCommission(b *bool) *SubGoodUpdate {
	if b != nil {
		sgu.SetCommission(*b)
	}
	return sgu
}

// ClearCommission clears the value of the "commission" field.
func (sgu *SubGoodUpdate) ClearCommission() *SubGoodUpdate {
	sgu.mutation.ClearCommission()
	return sgu
}

// Mutation returns the SubGoodMutation object of the builder.
func (sgu *SubGoodUpdate) Mutation() *SubGoodMutation {
	return sgu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sgu *SubGoodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := sgu.defaults(); err != nil {
		return 0, err
	}
	if len(sgu.hooks) == 0 {
		affected, err = sgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sgu.mutation = mutation
			affected, err = sgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sgu.hooks) - 1; i >= 0; i-- {
			if sgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sgu *SubGoodUpdate) SaveX(ctx context.Context) int {
	affected, err := sgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sgu *SubGoodUpdate) Exec(ctx context.Context) error {
	_, err := sgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sgu *SubGoodUpdate) ExecX(ctx context.Context) {
	if err := sgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sgu *SubGoodUpdate) defaults() error {
	if _, ok := sgu.mutation.UpdatedAt(); !ok {
		if subgood.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized subgood.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := subgood.UpdateDefaultUpdatedAt()
		sgu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sgu *SubGoodUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SubGoodUpdate {
	sgu.modifiers = append(sgu.modifiers, modifiers...)
	return sgu
}

func (sgu *SubGoodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subgood.Table,
			Columns: subgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subgood.FieldID,
			},
		},
	}
	if ps := sgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sgu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldCreatedAt,
		})
	}
	if value, ok := sgu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldCreatedAt,
		})
	}
	if value, ok := sgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldUpdatedAt,
		})
	}
	if value, ok := sgu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldUpdatedAt,
		})
	}
	if value, ok := sgu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldDeletedAt,
		})
	}
	if value, ok := sgu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldDeletedAt,
		})
	}
	if value, ok := sgu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subgood.FieldAppID,
		})
	}
	if value, ok := sgu.mutation.MainGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subgood.FieldMainGoodID,
		})
	}
	if value, ok := sgu.mutation.SubGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subgood.FieldSubGoodID,
		})
	}
	if value, ok := sgu.mutation.Must(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subgood.FieldMust,
		})
	}
	if sgu.mutation.MustCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: subgood.FieldMust,
		})
	}
	if value, ok := sgu.mutation.Commission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subgood.FieldCommission,
		})
	}
	if sgu.mutation.CommissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: subgood.FieldCommission,
		})
	}
	_spec.Modifiers = sgu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, sgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subgood.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SubGoodUpdateOne is the builder for updating a single SubGood entity.
type SubGoodUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SubGoodMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (sguo *SubGoodUpdateOne) SetCreatedAt(u uint32) *SubGoodUpdateOne {
	sguo.mutation.ResetCreatedAt()
	sguo.mutation.SetCreatedAt(u)
	return sguo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sguo *SubGoodUpdateOne) SetNillableCreatedAt(u *uint32) *SubGoodUpdateOne {
	if u != nil {
		sguo.SetCreatedAt(*u)
	}
	return sguo
}

// AddCreatedAt adds u to the "created_at" field.
func (sguo *SubGoodUpdateOne) AddCreatedAt(u int32) *SubGoodUpdateOne {
	sguo.mutation.AddCreatedAt(u)
	return sguo
}

// SetUpdatedAt sets the "updated_at" field.
func (sguo *SubGoodUpdateOne) SetUpdatedAt(u uint32) *SubGoodUpdateOne {
	sguo.mutation.ResetUpdatedAt()
	sguo.mutation.SetUpdatedAt(u)
	return sguo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (sguo *SubGoodUpdateOne) AddUpdatedAt(u int32) *SubGoodUpdateOne {
	sguo.mutation.AddUpdatedAt(u)
	return sguo
}

// SetDeletedAt sets the "deleted_at" field.
func (sguo *SubGoodUpdateOne) SetDeletedAt(u uint32) *SubGoodUpdateOne {
	sguo.mutation.ResetDeletedAt()
	sguo.mutation.SetDeletedAt(u)
	return sguo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sguo *SubGoodUpdateOne) SetNillableDeletedAt(u *uint32) *SubGoodUpdateOne {
	if u != nil {
		sguo.SetDeletedAt(*u)
	}
	return sguo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (sguo *SubGoodUpdateOne) AddDeletedAt(u int32) *SubGoodUpdateOne {
	sguo.mutation.AddDeletedAt(u)
	return sguo
}

// SetAppID sets the "app_id" field.
func (sguo *SubGoodUpdateOne) SetAppID(u uuid.UUID) *SubGoodUpdateOne {
	sguo.mutation.SetAppID(u)
	return sguo
}

// SetMainGoodID sets the "main_good_id" field.
func (sguo *SubGoodUpdateOne) SetMainGoodID(u uuid.UUID) *SubGoodUpdateOne {
	sguo.mutation.SetMainGoodID(u)
	return sguo
}

// SetSubGoodID sets the "sub_good_id" field.
func (sguo *SubGoodUpdateOne) SetSubGoodID(u uuid.UUID) *SubGoodUpdateOne {
	sguo.mutation.SetSubGoodID(u)
	return sguo
}

// SetMust sets the "must" field.
func (sguo *SubGoodUpdateOne) SetMust(b bool) *SubGoodUpdateOne {
	sguo.mutation.SetMust(b)
	return sguo
}

// SetNillableMust sets the "must" field if the given value is not nil.
func (sguo *SubGoodUpdateOne) SetNillableMust(b *bool) *SubGoodUpdateOne {
	if b != nil {
		sguo.SetMust(*b)
	}
	return sguo
}

// ClearMust clears the value of the "must" field.
func (sguo *SubGoodUpdateOne) ClearMust() *SubGoodUpdateOne {
	sguo.mutation.ClearMust()
	return sguo
}

// SetCommission sets the "commission" field.
func (sguo *SubGoodUpdateOne) SetCommission(b bool) *SubGoodUpdateOne {
	sguo.mutation.SetCommission(b)
	return sguo
}

// SetNillableCommission sets the "commission" field if the given value is not nil.
func (sguo *SubGoodUpdateOne) SetNillableCommission(b *bool) *SubGoodUpdateOne {
	if b != nil {
		sguo.SetCommission(*b)
	}
	return sguo
}

// ClearCommission clears the value of the "commission" field.
func (sguo *SubGoodUpdateOne) ClearCommission() *SubGoodUpdateOne {
	sguo.mutation.ClearCommission()
	return sguo
}

// Mutation returns the SubGoodMutation object of the builder.
func (sguo *SubGoodUpdateOne) Mutation() *SubGoodMutation {
	return sguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sguo *SubGoodUpdateOne) Select(field string, fields ...string) *SubGoodUpdateOne {
	sguo.fields = append([]string{field}, fields...)
	return sguo
}

// Save executes the query and returns the updated SubGood entity.
func (sguo *SubGoodUpdateOne) Save(ctx context.Context) (*SubGood, error) {
	var (
		err  error
		node *SubGood
	)
	if err := sguo.defaults(); err != nil {
		return nil, err
	}
	if len(sguo.hooks) == 0 {
		node, err = sguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sguo.mutation = mutation
			node, err = sguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sguo.hooks) - 1; i >= 0; i-- {
			if sguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SubGood)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SubGoodMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sguo *SubGoodUpdateOne) SaveX(ctx context.Context) *SubGood {
	node, err := sguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sguo *SubGoodUpdateOne) Exec(ctx context.Context) error {
	_, err := sguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sguo *SubGoodUpdateOne) ExecX(ctx context.Context) {
	if err := sguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sguo *SubGoodUpdateOne) defaults() error {
	if _, ok := sguo.mutation.UpdatedAt(); !ok {
		if subgood.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized subgood.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := subgood.UpdateDefaultUpdatedAt()
		sguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sguo *SubGoodUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SubGoodUpdateOne {
	sguo.modifiers = append(sguo.modifiers, modifiers...)
	return sguo
}

func (sguo *SubGoodUpdateOne) sqlSave(ctx context.Context) (_node *SubGood, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subgood.Table,
			Columns: subgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subgood.FieldID,
			},
		},
	}
	id, ok := sguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SubGood.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subgood.FieldID)
		for _, f := range fields {
			if !subgood.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subgood.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldCreatedAt,
		})
	}
	if value, ok := sguo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldCreatedAt,
		})
	}
	if value, ok := sguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldUpdatedAt,
		})
	}
	if value, ok := sguo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldUpdatedAt,
		})
	}
	if value, ok := sguo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldDeletedAt,
		})
	}
	if value, ok := sguo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subgood.FieldDeletedAt,
		})
	}
	if value, ok := sguo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subgood.FieldAppID,
		})
	}
	if value, ok := sguo.mutation.MainGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subgood.FieldMainGoodID,
		})
	}
	if value, ok := sguo.mutation.SubGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subgood.FieldSubGoodID,
		})
	}
	if value, ok := sguo.mutation.Must(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subgood.FieldMust,
		})
	}
	if sguo.mutation.MustCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: subgood.FieldMust,
		})
	}
	if value, ok := sguo.mutation.Commission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subgood.FieldCommission,
		})
	}
	if sguo.mutation.CommissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: subgood.FieldCommission,
		})
	}
	_spec.Modifiers = sguo.modifiers
	_node = &SubGood{config: sguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subgood.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}