// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/predicate"
	"cattleai/ent/warehousesetting"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WarehouseSettingUpdate is the builder for updating WarehouseSetting entities.
type WarehouseSettingUpdate struct {
	config
	hooks    []Hook
	mutation *WarehouseSettingMutation
}

// Where appends a list predicates to the WarehouseSettingUpdate builder.
func (wsu *WarehouseSettingUpdate) Where(ps ...predicate.WarehouseSetting) *WarehouseSettingUpdate {
	wsu.mutation.Where(ps...)
	return wsu
}

// SetName sets the "name" field.
func (wsu *WarehouseSettingUpdate) SetName(s string) *WarehouseSettingUpdate {
	wsu.mutation.SetName(s)
	return wsu
}

// SetType sets the "type" field.
func (wsu *WarehouseSettingUpdate) SetType(s string) *WarehouseSettingUpdate {
	wsu.mutation.SetType(s)
	return wsu
}

// SetTenantId sets the "tenantId" field.
func (wsu *WarehouseSettingUpdate) SetTenantId(i int64) *WarehouseSettingUpdate {
	wsu.mutation.ResetTenantId()
	wsu.mutation.SetTenantId(i)
	return wsu
}

// AddTenantId adds i to the "tenantId" field.
func (wsu *WarehouseSettingUpdate) AddTenantId(i int64) *WarehouseSettingUpdate {
	wsu.mutation.AddTenantId(i)
	return wsu
}

// SetTenantName sets the "tenantName" field.
func (wsu *WarehouseSettingUpdate) SetTenantName(s string) *WarehouseSettingUpdate {
	wsu.mutation.SetTenantName(s)
	return wsu
}

// SetRemarks sets the "remarks" field.
func (wsu *WarehouseSettingUpdate) SetRemarks(s string) *WarehouseSettingUpdate {
	wsu.mutation.SetRemarks(s)
	return wsu
}

// SetCreatedAt sets the "createdAt" field.
func (wsu *WarehouseSettingUpdate) SetCreatedAt(i int64) *WarehouseSettingUpdate {
	wsu.mutation.ResetCreatedAt()
	wsu.mutation.SetCreatedAt(i)
	return wsu
}

// AddCreatedAt adds i to the "createdAt" field.
func (wsu *WarehouseSettingUpdate) AddCreatedAt(i int64) *WarehouseSettingUpdate {
	wsu.mutation.AddCreatedAt(i)
	return wsu
}

// SetUpdatedAt sets the "updatedAt" field.
func (wsu *WarehouseSettingUpdate) SetUpdatedAt(i int64) *WarehouseSettingUpdate {
	wsu.mutation.ResetUpdatedAt()
	wsu.mutation.SetUpdatedAt(i)
	return wsu
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (wsu *WarehouseSettingUpdate) AddUpdatedAt(i int64) *WarehouseSettingUpdate {
	wsu.mutation.AddUpdatedAt(i)
	return wsu
}

// SetDeleted sets the "deleted" field.
func (wsu *WarehouseSettingUpdate) SetDeleted(i int) *WarehouseSettingUpdate {
	wsu.mutation.ResetDeleted()
	wsu.mutation.SetDeleted(i)
	return wsu
}

// AddDeleted adds i to the "deleted" field.
func (wsu *WarehouseSettingUpdate) AddDeleted(i int) *WarehouseSettingUpdate {
	wsu.mutation.AddDeleted(i)
	return wsu
}

// Mutation returns the WarehouseSettingMutation object of the builder.
func (wsu *WarehouseSettingUpdate) Mutation() *WarehouseSettingMutation {
	return wsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wsu *WarehouseSettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wsu.hooks) == 0 {
		affected, err = wsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WarehouseSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wsu.mutation = mutation
			affected, err = wsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wsu.hooks) - 1; i >= 0; i-- {
			if wsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wsu *WarehouseSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := wsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wsu *WarehouseSettingUpdate) Exec(ctx context.Context) error {
	_, err := wsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsu *WarehouseSettingUpdate) ExecX(ctx context.Context) {
	if err := wsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wsu *WarehouseSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   warehousesetting.Table,
			Columns: warehousesetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: warehousesetting.FieldID,
			},
		},
	}
	if ps := wsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: warehousesetting.FieldName,
		})
	}
	if value, ok := wsu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: warehousesetting.FieldType,
		})
	}
	if value, ok := wsu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldTenantId,
		})
	}
	if value, ok := wsu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldTenantId,
		})
	}
	if value, ok := wsu.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: warehousesetting.FieldTenantName,
		})
	}
	if value, ok := wsu.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: warehousesetting.FieldRemarks,
		})
	}
	if value, ok := wsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldCreatedAt,
		})
	}
	if value, ok := wsu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldCreatedAt,
		})
	}
	if value, ok := wsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldUpdatedAt,
		})
	}
	if value, ok := wsu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldUpdatedAt,
		})
	}
	if value, ok := wsu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: warehousesetting.FieldDeleted,
		})
	}
	if value, ok := wsu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: warehousesetting.FieldDeleted,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{warehousesetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// WarehouseSettingUpdateOne is the builder for updating a single WarehouseSetting entity.
type WarehouseSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WarehouseSettingMutation
}

// SetName sets the "name" field.
func (wsuo *WarehouseSettingUpdateOne) SetName(s string) *WarehouseSettingUpdateOne {
	wsuo.mutation.SetName(s)
	return wsuo
}

// SetType sets the "type" field.
func (wsuo *WarehouseSettingUpdateOne) SetType(s string) *WarehouseSettingUpdateOne {
	wsuo.mutation.SetType(s)
	return wsuo
}

// SetTenantId sets the "tenantId" field.
func (wsuo *WarehouseSettingUpdateOne) SetTenantId(i int64) *WarehouseSettingUpdateOne {
	wsuo.mutation.ResetTenantId()
	wsuo.mutation.SetTenantId(i)
	return wsuo
}

// AddTenantId adds i to the "tenantId" field.
func (wsuo *WarehouseSettingUpdateOne) AddTenantId(i int64) *WarehouseSettingUpdateOne {
	wsuo.mutation.AddTenantId(i)
	return wsuo
}

// SetTenantName sets the "tenantName" field.
func (wsuo *WarehouseSettingUpdateOne) SetTenantName(s string) *WarehouseSettingUpdateOne {
	wsuo.mutation.SetTenantName(s)
	return wsuo
}

// SetRemarks sets the "remarks" field.
func (wsuo *WarehouseSettingUpdateOne) SetRemarks(s string) *WarehouseSettingUpdateOne {
	wsuo.mutation.SetRemarks(s)
	return wsuo
}

// SetCreatedAt sets the "createdAt" field.
func (wsuo *WarehouseSettingUpdateOne) SetCreatedAt(i int64) *WarehouseSettingUpdateOne {
	wsuo.mutation.ResetCreatedAt()
	wsuo.mutation.SetCreatedAt(i)
	return wsuo
}

// AddCreatedAt adds i to the "createdAt" field.
func (wsuo *WarehouseSettingUpdateOne) AddCreatedAt(i int64) *WarehouseSettingUpdateOne {
	wsuo.mutation.AddCreatedAt(i)
	return wsuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (wsuo *WarehouseSettingUpdateOne) SetUpdatedAt(i int64) *WarehouseSettingUpdateOne {
	wsuo.mutation.ResetUpdatedAt()
	wsuo.mutation.SetUpdatedAt(i)
	return wsuo
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (wsuo *WarehouseSettingUpdateOne) AddUpdatedAt(i int64) *WarehouseSettingUpdateOne {
	wsuo.mutation.AddUpdatedAt(i)
	return wsuo
}

// SetDeleted sets the "deleted" field.
func (wsuo *WarehouseSettingUpdateOne) SetDeleted(i int) *WarehouseSettingUpdateOne {
	wsuo.mutation.ResetDeleted()
	wsuo.mutation.SetDeleted(i)
	return wsuo
}

// AddDeleted adds i to the "deleted" field.
func (wsuo *WarehouseSettingUpdateOne) AddDeleted(i int) *WarehouseSettingUpdateOne {
	wsuo.mutation.AddDeleted(i)
	return wsuo
}

// Mutation returns the WarehouseSettingMutation object of the builder.
func (wsuo *WarehouseSettingUpdateOne) Mutation() *WarehouseSettingMutation {
	return wsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wsuo *WarehouseSettingUpdateOne) Select(field string, fields ...string) *WarehouseSettingUpdateOne {
	wsuo.fields = append([]string{field}, fields...)
	return wsuo
}

// Save executes the query and returns the updated WarehouseSetting entity.
func (wsuo *WarehouseSettingUpdateOne) Save(ctx context.Context) (*WarehouseSetting, error) {
	var (
		err  error
		node *WarehouseSetting
	)
	if len(wsuo.hooks) == 0 {
		node, err = wsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WarehouseSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wsuo.mutation = mutation
			node, err = wsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wsuo.hooks) - 1; i >= 0; i-- {
			if wsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wsuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*WarehouseSetting)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WarehouseSettingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wsuo *WarehouseSettingUpdateOne) SaveX(ctx context.Context) *WarehouseSetting {
	node, err := wsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wsuo *WarehouseSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := wsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsuo *WarehouseSettingUpdateOne) ExecX(ctx context.Context) {
	if err := wsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wsuo *WarehouseSettingUpdateOne) sqlSave(ctx context.Context) (_node *WarehouseSetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   warehousesetting.Table,
			Columns: warehousesetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: warehousesetting.FieldID,
			},
		},
	}
	id, ok := wsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WarehouseSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, warehousesetting.FieldID)
		for _, f := range fields {
			if !warehousesetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != warehousesetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: warehousesetting.FieldName,
		})
	}
	if value, ok := wsuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: warehousesetting.FieldType,
		})
	}
	if value, ok := wsuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldTenantId,
		})
	}
	if value, ok := wsuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldTenantId,
		})
	}
	if value, ok := wsuo.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: warehousesetting.FieldTenantName,
		})
	}
	if value, ok := wsuo.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: warehousesetting.FieldRemarks,
		})
	}
	if value, ok := wsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldCreatedAt,
		})
	}
	if value, ok := wsuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldCreatedAt,
		})
	}
	if value, ok := wsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldUpdatedAt,
		})
	}
	if value, ok := wsuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: warehousesetting.FieldUpdatedAt,
		})
	}
	if value, ok := wsuo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: warehousesetting.FieldDeleted,
		})
	}
	if value, ok := wsuo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: warehousesetting.FieldDeleted,
		})
	}
	_node = &WarehouseSetting{config: wsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{warehousesetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
