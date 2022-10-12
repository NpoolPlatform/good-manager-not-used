// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/deviceinfo"
	"github.com/google/uuid"
)

// DeviceInfoCreate is the builder for creating a DeviceInfo entity.
type DeviceInfoCreate struct {
	config
	mutation *DeviceInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dic *DeviceInfoCreate) SetCreatedAt(u uint32) *DeviceInfoCreate {
	dic.mutation.SetCreatedAt(u)
	return dic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableCreatedAt(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetCreatedAt(*u)
	}
	return dic
}

// SetUpdatedAt sets the "updated_at" field.
func (dic *DeviceInfoCreate) SetUpdatedAt(u uint32) *DeviceInfoCreate {
	dic.mutation.SetUpdatedAt(u)
	return dic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableUpdatedAt(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetUpdatedAt(*u)
	}
	return dic
}

// SetDeletedAt sets the "deleted_at" field.
func (dic *DeviceInfoCreate) SetDeletedAt(u uint32) *DeviceInfoCreate {
	dic.mutation.SetDeletedAt(u)
	return dic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableDeletedAt(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetDeletedAt(*u)
	}
	return dic
}

// SetType sets the "type" field.
func (dic *DeviceInfoCreate) SetType(s string) *DeviceInfoCreate {
	dic.mutation.SetType(s)
	return dic
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableType(s *string) *DeviceInfoCreate {
	if s != nil {
		dic.SetType(*s)
	}
	return dic
}

// SetManufacturer sets the "manufacturer" field.
func (dic *DeviceInfoCreate) SetManufacturer(s string) *DeviceInfoCreate {
	dic.mutation.SetManufacturer(s)
	return dic
}

// SetNillableManufacturer sets the "manufacturer" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableManufacturer(s *string) *DeviceInfoCreate {
	if s != nil {
		dic.SetManufacturer(*s)
	}
	return dic
}

// SetPowerComsuption sets the "power_comsuption" field.
func (dic *DeviceInfoCreate) SetPowerComsuption(u uint32) *DeviceInfoCreate {
	dic.mutation.SetPowerComsuption(u)
	return dic
}

// SetNillablePowerComsuption sets the "power_comsuption" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillablePowerComsuption(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetPowerComsuption(*u)
	}
	return dic
}

// SetShipmentAt sets the "shipment_at" field.
func (dic *DeviceInfoCreate) SetShipmentAt(u uint32) *DeviceInfoCreate {
	dic.mutation.SetShipmentAt(u)
	return dic
}

// SetNillableShipmentAt sets the "shipment_at" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableShipmentAt(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetShipmentAt(*u)
	}
	return dic
}

// SetPosters sets the "posters" field.
func (dic *DeviceInfoCreate) SetPosters(s []string) *DeviceInfoCreate {
	dic.mutation.SetPosters(s)
	return dic
}

// SetID sets the "id" field.
func (dic *DeviceInfoCreate) SetID(u uuid.UUID) *DeviceInfoCreate {
	dic.mutation.SetID(u)
	return dic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableID(u *uuid.UUID) *DeviceInfoCreate {
	if u != nil {
		dic.SetID(*u)
	}
	return dic
}

// Mutation returns the DeviceInfoMutation object of the builder.
func (dic *DeviceInfoCreate) Mutation() *DeviceInfoMutation {
	return dic.mutation
}

// Save creates the DeviceInfo in the database.
func (dic *DeviceInfoCreate) Save(ctx context.Context) (*DeviceInfo, error) {
	var (
		err  error
		node *DeviceInfo
	)
	if err := dic.defaults(); err != nil {
		return nil, err
	}
	if len(dic.hooks) == 0 {
		if err = dic.check(); err != nil {
			return nil, err
		}
		node, err = dic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dic.check(); err != nil {
				return nil, err
			}
			dic.mutation = mutation
			if node, err = dic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dic.hooks) - 1; i >= 0; i-- {
			if dic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DeviceInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DeviceInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dic *DeviceInfoCreate) SaveX(ctx context.Context) *DeviceInfo {
	v, err := dic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dic *DeviceInfoCreate) Exec(ctx context.Context) error {
	_, err := dic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dic *DeviceInfoCreate) ExecX(ctx context.Context) {
	if err := dic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dic *DeviceInfoCreate) defaults() error {
	if _, ok := dic.mutation.CreatedAt(); !ok {
		if deviceinfo.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized deviceinfo.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := deviceinfo.DefaultCreatedAt()
		dic.mutation.SetCreatedAt(v)
	}
	if _, ok := dic.mutation.UpdatedAt(); !ok {
		if deviceinfo.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized deviceinfo.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := deviceinfo.DefaultUpdatedAt()
		dic.mutation.SetUpdatedAt(v)
	}
	if _, ok := dic.mutation.DeletedAt(); !ok {
		if deviceinfo.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized deviceinfo.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := deviceinfo.DefaultDeletedAt()
		dic.mutation.SetDeletedAt(v)
	}
	if _, ok := dic.mutation.GetType(); !ok {
		v := deviceinfo.DefaultType
		dic.mutation.SetType(v)
	}
	if _, ok := dic.mutation.Manufacturer(); !ok {
		v := deviceinfo.DefaultManufacturer
		dic.mutation.SetManufacturer(v)
	}
	if _, ok := dic.mutation.PowerComsuption(); !ok {
		v := deviceinfo.DefaultPowerComsuption
		dic.mutation.SetPowerComsuption(v)
	}
	if _, ok := dic.mutation.ShipmentAt(); !ok {
		v := deviceinfo.DefaultShipmentAt
		dic.mutation.SetShipmentAt(v)
	}
	if _, ok := dic.mutation.Posters(); !ok {
		v := deviceinfo.DefaultPosters
		dic.mutation.SetPosters(v)
	}
	if _, ok := dic.mutation.ID(); !ok {
		if deviceinfo.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized deviceinfo.DefaultID (forgotten import ent/runtime?)")
		}
		v := deviceinfo.DefaultID()
		dic.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dic *DeviceInfoCreate) check() error {
	if _, ok := dic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeviceInfo.created_at"`)}
	}
	if _, ok := dic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DeviceInfo.updated_at"`)}
	}
	if _, ok := dic.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "DeviceInfo.deleted_at"`)}
	}
	if v, ok := dic.mutation.GetType(); ok {
		if err := deviceinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "DeviceInfo.type": %w`, err)}
		}
	}
	if v, ok := dic.mutation.Manufacturer(); ok {
		if err := deviceinfo.ManufacturerValidator(v); err != nil {
			return &ValidationError{Name: "manufacturer", err: fmt.Errorf(`ent: validator failed for field "DeviceInfo.manufacturer": %w`, err)}
		}
	}
	return nil
}

func (dic *DeviceInfoCreate) sqlSave(ctx context.Context) (*DeviceInfo, error) {
	_node, _spec := dic.createSpec()
	if err := sqlgraph.CreateNode(ctx, dic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (dic *DeviceInfoCreate) createSpec() (*DeviceInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceInfo{config: dic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deviceinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deviceinfo.FieldID,
			},
		}
	)
	_spec.OnConflict = dic.conflict
	if id, ok := dic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dic.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := dic.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deviceinfo.FieldType,
		})
		_node.Type = value
	}
	if value, ok := dic.mutation.Manufacturer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deviceinfo.FieldManufacturer,
		})
		_node.Manufacturer = value
	}
	if value, ok := dic.mutation.PowerComsuption(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldPowerComsuption,
		})
		_node.PowerComsuption = value
	}
	if value, ok := dic.mutation.ShipmentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldShipmentAt,
		})
		_node.ShipmentAt = value
	}
	if value, ok := dic.mutation.Posters(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: deviceinfo.FieldPosters,
		})
		_node.Posters = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceInfo.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceInfoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (dic *DeviceInfoCreate) OnConflict(opts ...sql.ConflictOption) *DeviceInfoUpsertOne {
	dic.conflict = opts
	return &DeviceInfoUpsertOne{
		create: dic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dic *DeviceInfoCreate) OnConflictColumns(columns ...string) *DeviceInfoUpsertOne {
	dic.conflict = append(dic.conflict, sql.ConflictColumns(columns...))
	return &DeviceInfoUpsertOne{
		create: dic,
	}
}

type (
	// DeviceInfoUpsertOne is the builder for "upsert"-ing
	//  one DeviceInfo node.
	DeviceInfoUpsertOne struct {
		create *DeviceInfoCreate
	}

	// DeviceInfoUpsert is the "OnConflict" setter.
	DeviceInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *DeviceInfoUpsert) SetCreatedAt(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateCreatedAt() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceInfoUpsert) AddCreatedAt(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceInfoUpsert) SetUpdatedAt(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateUpdatedAt() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceInfoUpsert) AddUpdatedAt(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceInfoUpsert) SetDeletedAt(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateDeletedAt() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceInfoUpsert) AddDeletedAt(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldDeletedAt, v)
	return u
}

// SetType sets the "type" field.
func (u *DeviceInfoUpsert) SetType(v string) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateType() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldType)
	return u
}

// ClearType clears the value of the "type" field.
func (u *DeviceInfoUpsert) ClearType() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldType)
	return u
}

// SetManufacturer sets the "manufacturer" field.
func (u *DeviceInfoUpsert) SetManufacturer(v string) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldManufacturer, v)
	return u
}

// UpdateManufacturer sets the "manufacturer" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateManufacturer() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldManufacturer)
	return u
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (u *DeviceInfoUpsert) ClearManufacturer() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldManufacturer)
	return u
}

// SetPowerComsuption sets the "power_comsuption" field.
func (u *DeviceInfoUpsert) SetPowerComsuption(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldPowerComsuption, v)
	return u
}

// UpdatePowerComsuption sets the "power_comsuption" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdatePowerComsuption() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldPowerComsuption)
	return u
}

// AddPowerComsuption adds v to the "power_comsuption" field.
func (u *DeviceInfoUpsert) AddPowerComsuption(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldPowerComsuption, v)
	return u
}

// ClearPowerComsuption clears the value of the "power_comsuption" field.
func (u *DeviceInfoUpsert) ClearPowerComsuption() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldPowerComsuption)
	return u
}

// SetShipmentAt sets the "shipment_at" field.
func (u *DeviceInfoUpsert) SetShipmentAt(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldShipmentAt, v)
	return u
}

// UpdateShipmentAt sets the "shipment_at" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateShipmentAt() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldShipmentAt)
	return u
}

// AddShipmentAt adds v to the "shipment_at" field.
func (u *DeviceInfoUpsert) AddShipmentAt(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldShipmentAt, v)
	return u
}

// ClearShipmentAt clears the value of the "shipment_at" field.
func (u *DeviceInfoUpsert) ClearShipmentAt() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldShipmentAt)
	return u
}

// SetPosters sets the "posters" field.
func (u *DeviceInfoUpsert) SetPosters(v []string) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldPosters, v)
	return u
}

// UpdatePosters sets the "posters" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdatePosters() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldPosters)
	return u
}

// ClearPosters clears the value of the "posters" field.
func (u *DeviceInfoUpsert) ClearPosters() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldPosters)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DeviceInfoUpsertOne) UpdateNewValues() *DeviceInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(deviceinfo.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.DeviceInfo.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DeviceInfoUpsertOne) Ignore() *DeviceInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceInfoUpsertOne) DoNothing() *DeviceInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceInfoCreate.OnConflict
// documentation for more info.
func (u *DeviceInfoUpsertOne) Update(set func(*DeviceInfoUpsert)) *DeviceInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DeviceInfoUpsertOne) SetCreatedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceInfoUpsertOne) AddCreatedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateCreatedAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceInfoUpsertOne) SetUpdatedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceInfoUpsertOne) AddUpdatedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateUpdatedAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceInfoUpsertOne) SetDeletedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceInfoUpsertOne) AddDeletedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateDeletedAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetType sets the "type" field.
func (u *DeviceInfoUpsertOne) SetType(v string) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateType() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *DeviceInfoUpsertOne) ClearType() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearType()
	})
}

// SetManufacturer sets the "manufacturer" field.
func (u *DeviceInfoUpsertOne) SetManufacturer(v string) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetManufacturer(v)
	})
}

// UpdateManufacturer sets the "manufacturer" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateManufacturer() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateManufacturer()
	})
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (u *DeviceInfoUpsertOne) ClearManufacturer() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearManufacturer()
	})
}

// SetPowerComsuption sets the "power_comsuption" field.
func (u *DeviceInfoUpsertOne) SetPowerComsuption(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetPowerComsuption(v)
	})
}

// AddPowerComsuption adds v to the "power_comsuption" field.
func (u *DeviceInfoUpsertOne) AddPowerComsuption(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddPowerComsuption(v)
	})
}

// UpdatePowerComsuption sets the "power_comsuption" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdatePowerComsuption() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdatePowerComsuption()
	})
}

// ClearPowerComsuption clears the value of the "power_comsuption" field.
func (u *DeviceInfoUpsertOne) ClearPowerComsuption() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearPowerComsuption()
	})
}

// SetShipmentAt sets the "shipment_at" field.
func (u *DeviceInfoUpsertOne) SetShipmentAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetShipmentAt(v)
	})
}

// AddShipmentAt adds v to the "shipment_at" field.
func (u *DeviceInfoUpsertOne) AddShipmentAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddShipmentAt(v)
	})
}

// UpdateShipmentAt sets the "shipment_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateShipmentAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateShipmentAt()
	})
}

// ClearShipmentAt clears the value of the "shipment_at" field.
func (u *DeviceInfoUpsertOne) ClearShipmentAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearShipmentAt()
	})
}

// SetPosters sets the "posters" field.
func (u *DeviceInfoUpsertOne) SetPosters(v []string) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetPosters(v)
	})
}

// UpdatePosters sets the "posters" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdatePosters() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdatePosters()
	})
}

// ClearPosters clears the value of the "posters" field.
func (u *DeviceInfoUpsertOne) ClearPosters() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearPosters()
	})
}

// Exec executes the query.
func (u *DeviceInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceInfoUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DeviceInfoUpsertOne.ID is not supported by MySQL driver. Use DeviceInfoUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceInfoUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceInfoCreateBulk is the builder for creating many DeviceInfo entities in bulk.
type DeviceInfoCreateBulk struct {
	config
	builders []*DeviceInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the DeviceInfo entities in the database.
func (dicb *DeviceInfoCreateBulk) Save(ctx context.Context) ([]*DeviceInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dicb.builders))
	nodes := make([]*DeviceInfo, len(dicb.builders))
	mutators := make([]Mutator, len(dicb.builders))
	for i := range dicb.builders {
		func(i int, root context.Context) {
			builder := dicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dicb *DeviceInfoCreateBulk) SaveX(ctx context.Context) []*DeviceInfo {
	v, err := dicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dicb *DeviceInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := dicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dicb *DeviceInfoCreateBulk) ExecX(ctx context.Context) {
	if err := dicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceInfoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (dicb *DeviceInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceInfoUpsertBulk {
	dicb.conflict = opts
	return &DeviceInfoUpsertBulk{
		create: dicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dicb *DeviceInfoCreateBulk) OnConflictColumns(columns ...string) *DeviceInfoUpsertBulk {
	dicb.conflict = append(dicb.conflict, sql.ConflictColumns(columns...))
	return &DeviceInfoUpsertBulk{
		create: dicb,
	}
}

// DeviceInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of DeviceInfo nodes.
type DeviceInfoUpsertBulk struct {
	create *DeviceInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DeviceInfoUpsertBulk) UpdateNewValues() *DeviceInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(deviceinfo.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DeviceInfoUpsertBulk) Ignore() *DeviceInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceInfoUpsertBulk) DoNothing() *DeviceInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceInfoCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceInfoUpsertBulk) Update(set func(*DeviceInfoUpsert)) *DeviceInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DeviceInfoUpsertBulk) SetCreatedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceInfoUpsertBulk) AddCreatedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateCreatedAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceInfoUpsertBulk) SetUpdatedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceInfoUpsertBulk) AddUpdatedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateUpdatedAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceInfoUpsertBulk) SetDeletedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceInfoUpsertBulk) AddDeletedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateDeletedAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetType sets the "type" field.
func (u *DeviceInfoUpsertBulk) SetType(v string) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateType() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *DeviceInfoUpsertBulk) ClearType() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearType()
	})
}

// SetManufacturer sets the "manufacturer" field.
func (u *DeviceInfoUpsertBulk) SetManufacturer(v string) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetManufacturer(v)
	})
}

// UpdateManufacturer sets the "manufacturer" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateManufacturer() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateManufacturer()
	})
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (u *DeviceInfoUpsertBulk) ClearManufacturer() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearManufacturer()
	})
}

// SetPowerComsuption sets the "power_comsuption" field.
func (u *DeviceInfoUpsertBulk) SetPowerComsuption(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetPowerComsuption(v)
	})
}

// AddPowerComsuption adds v to the "power_comsuption" field.
func (u *DeviceInfoUpsertBulk) AddPowerComsuption(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddPowerComsuption(v)
	})
}

// UpdatePowerComsuption sets the "power_comsuption" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdatePowerComsuption() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdatePowerComsuption()
	})
}

// ClearPowerComsuption clears the value of the "power_comsuption" field.
func (u *DeviceInfoUpsertBulk) ClearPowerComsuption() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearPowerComsuption()
	})
}

// SetShipmentAt sets the "shipment_at" field.
func (u *DeviceInfoUpsertBulk) SetShipmentAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetShipmentAt(v)
	})
}

// AddShipmentAt adds v to the "shipment_at" field.
func (u *DeviceInfoUpsertBulk) AddShipmentAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddShipmentAt(v)
	})
}

// UpdateShipmentAt sets the "shipment_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateShipmentAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateShipmentAt()
	})
}

// ClearShipmentAt clears the value of the "shipment_at" field.
func (u *DeviceInfoUpsertBulk) ClearShipmentAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearShipmentAt()
	})
}

// SetPosters sets the "posters" field.
func (u *DeviceInfoUpsertBulk) SetPosters(v []string) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetPosters(v)
	})
}

// UpdatePosters sets the "posters" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdatePosters() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdatePosters()
	})
}

// ClearPosters clears the value of the "posters" field.
func (u *DeviceInfoUpsertBulk) ClearPosters() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearPosters()
	})
}

// Exec executes the query.
func (u *DeviceInfoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeviceInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}