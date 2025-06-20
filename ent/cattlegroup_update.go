// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/cattlegroup"
	"cattleai/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CattleGroupUpdate is the builder for updating CattleGroup entities.
type CattleGroupUpdate struct {
	config
	hooks    []Hook
	mutation *CattleGroupMutation
}

// Where appends a list predicates to the CattleGroupUpdate builder.
func (cgu *CattleGroupUpdate) Where(ps ...predicate.CattleGroup) *CattleGroupUpdate {
	cgu.mutation.Where(ps...)
	return cgu
}

// SetEarNumber sets the "earNumber" field.
func (cgu *CattleGroupUpdate) SetEarNumber(s string) *CattleGroupUpdate {
	cgu.mutation.SetEarNumber(s)
	return cgu
}

// SetDate sets the "date" field.
func (cgu *CattleGroupUpdate) SetDate(i int64) *CattleGroupUpdate {
	cgu.mutation.ResetDate()
	cgu.mutation.SetDate(i)
	return cgu
}

// AddDate adds i to the "date" field.
func (cgu *CattleGroupUpdate) AddDate(i int64) *CattleGroupUpdate {
	cgu.mutation.AddDate(i)
	return cgu
}

// SetToShed sets the "toShed" field.
func (cgu *CattleGroupUpdate) SetToShed(s string) *CattleGroupUpdate {
	cgu.mutation.SetToShed(s)
	return cgu
}

// SetTenantId sets the "tenantId" field.
func (cgu *CattleGroupUpdate) SetTenantId(i int64) *CattleGroupUpdate {
	cgu.mutation.ResetTenantId()
	cgu.mutation.SetTenantId(i)
	return cgu
}

// AddTenantId adds i to the "tenantId" field.
func (cgu *CattleGroupUpdate) AddTenantId(i int64) *CattleGroupUpdate {
	cgu.mutation.AddTenantId(i)
	return cgu
}

// SetTenantName sets the "tenantName" field.
func (cgu *CattleGroupUpdate) SetTenantName(s string) *CattleGroupUpdate {
	cgu.mutation.SetTenantName(s)
	return cgu
}

// SetRemarks sets the "remarks" field.
func (cgu *CattleGroupUpdate) SetRemarks(s string) *CattleGroupUpdate {
	cgu.mutation.SetRemarks(s)
	return cgu
}

// SetCreatedAt sets the "createdAt" field.
func (cgu *CattleGroupUpdate) SetCreatedAt(i int64) *CattleGroupUpdate {
	cgu.mutation.ResetCreatedAt()
	cgu.mutation.SetCreatedAt(i)
	return cgu
}

// AddCreatedAt adds i to the "createdAt" field.
func (cgu *CattleGroupUpdate) AddCreatedAt(i int64) *CattleGroupUpdate {
	cgu.mutation.AddCreatedAt(i)
	return cgu
}

// SetUpdatedAt sets the "updatedAt" field.
func (cgu *CattleGroupUpdate) SetUpdatedAt(i int64) *CattleGroupUpdate {
	cgu.mutation.ResetUpdatedAt()
	cgu.mutation.SetUpdatedAt(i)
	return cgu
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (cgu *CattleGroupUpdate) AddUpdatedAt(i int64) *CattleGroupUpdate {
	cgu.mutation.AddUpdatedAt(i)
	return cgu
}

// SetDeleted sets the "deleted" field.
func (cgu *CattleGroupUpdate) SetDeleted(i int) *CattleGroupUpdate {
	cgu.mutation.ResetDeleted()
	cgu.mutation.SetDeleted(i)
	return cgu
}

// AddDeleted adds i to the "deleted" field.
func (cgu *CattleGroupUpdate) AddDeleted(i int) *CattleGroupUpdate {
	cgu.mutation.AddDeleted(i)
	return cgu
}

// Mutation returns the CattleGroupMutation object of the builder.
func (cgu *CattleGroupUpdate) Mutation() *CattleGroupMutation {
	return cgu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cgu *CattleGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cgu.hooks) == 0 {
		affected, err = cgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CattleGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cgu.mutation = mutation
			affected, err = cgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cgu.hooks) - 1; i >= 0; i-- {
			if cgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cgu *CattleGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := cgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cgu *CattleGroupUpdate) Exec(ctx context.Context) error {
	_, err := cgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cgu *CattleGroupUpdate) ExecX(ctx context.Context) {
	if err := cgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cgu *CattleGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cattlegroup.Table,
			Columns: cattlegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cattlegroup.FieldID,
			},
		},
	}
	if ps := cgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cgu.mutation.EarNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlegroup.FieldEarNumber,
		})
	}
	if value, ok := cgu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldDate,
		})
	}
	if value, ok := cgu.mutation.AddedDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldDate,
		})
	}
	if value, ok := cgu.mutation.ToShed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlegroup.FieldToShed,
		})
	}
	if value, ok := cgu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldTenantId,
		})
	}
	if value, ok := cgu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldTenantId,
		})
	}
	if value, ok := cgu.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlegroup.FieldTenantName,
		})
	}
	if value, ok := cgu.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlegroup.FieldRemarks,
		})
	}
	if value, ok := cgu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldCreatedAt,
		})
	}
	if value, ok := cgu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldCreatedAt,
		})
	}
	if value, ok := cgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldUpdatedAt,
		})
	}
	if value, ok := cgu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldUpdatedAt,
		})
	}
	if value, ok := cgu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cattlegroup.FieldDeleted,
		})
	}
	if value, ok := cgu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cattlegroup.FieldDeleted,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cattlegroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CattleGroupUpdateOne is the builder for updating a single CattleGroup entity.
type CattleGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CattleGroupMutation
}

// SetEarNumber sets the "earNumber" field.
func (cguo *CattleGroupUpdateOne) SetEarNumber(s string) *CattleGroupUpdateOne {
	cguo.mutation.SetEarNumber(s)
	return cguo
}

// SetDate sets the "date" field.
func (cguo *CattleGroupUpdateOne) SetDate(i int64) *CattleGroupUpdateOne {
	cguo.mutation.ResetDate()
	cguo.mutation.SetDate(i)
	return cguo
}

// AddDate adds i to the "date" field.
func (cguo *CattleGroupUpdateOne) AddDate(i int64) *CattleGroupUpdateOne {
	cguo.mutation.AddDate(i)
	return cguo
}

// SetToShed sets the "toShed" field.
func (cguo *CattleGroupUpdateOne) SetToShed(s string) *CattleGroupUpdateOne {
	cguo.mutation.SetToShed(s)
	return cguo
}

// SetTenantId sets the "tenantId" field.
func (cguo *CattleGroupUpdateOne) SetTenantId(i int64) *CattleGroupUpdateOne {
	cguo.mutation.ResetTenantId()
	cguo.mutation.SetTenantId(i)
	return cguo
}

// AddTenantId adds i to the "tenantId" field.
func (cguo *CattleGroupUpdateOne) AddTenantId(i int64) *CattleGroupUpdateOne {
	cguo.mutation.AddTenantId(i)
	return cguo
}

// SetTenantName sets the "tenantName" field.
func (cguo *CattleGroupUpdateOne) SetTenantName(s string) *CattleGroupUpdateOne {
	cguo.mutation.SetTenantName(s)
	return cguo
}

// SetRemarks sets the "remarks" field.
func (cguo *CattleGroupUpdateOne) SetRemarks(s string) *CattleGroupUpdateOne {
	cguo.mutation.SetRemarks(s)
	return cguo
}

// SetCreatedAt sets the "createdAt" field.
func (cguo *CattleGroupUpdateOne) SetCreatedAt(i int64) *CattleGroupUpdateOne {
	cguo.mutation.ResetCreatedAt()
	cguo.mutation.SetCreatedAt(i)
	return cguo
}

// AddCreatedAt adds i to the "createdAt" field.
func (cguo *CattleGroupUpdateOne) AddCreatedAt(i int64) *CattleGroupUpdateOne {
	cguo.mutation.AddCreatedAt(i)
	return cguo
}

// SetUpdatedAt sets the "updatedAt" field.
func (cguo *CattleGroupUpdateOne) SetUpdatedAt(i int64) *CattleGroupUpdateOne {
	cguo.mutation.ResetUpdatedAt()
	cguo.mutation.SetUpdatedAt(i)
	return cguo
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (cguo *CattleGroupUpdateOne) AddUpdatedAt(i int64) *CattleGroupUpdateOne {
	cguo.mutation.AddUpdatedAt(i)
	return cguo
}

// SetDeleted sets the "deleted" field.
func (cguo *CattleGroupUpdateOne) SetDeleted(i int) *CattleGroupUpdateOne {
	cguo.mutation.ResetDeleted()
	cguo.mutation.SetDeleted(i)
	return cguo
}

// AddDeleted adds i to the "deleted" field.
func (cguo *CattleGroupUpdateOne) AddDeleted(i int) *CattleGroupUpdateOne {
	cguo.mutation.AddDeleted(i)
	return cguo
}

// Mutation returns the CattleGroupMutation object of the builder.
func (cguo *CattleGroupUpdateOne) Mutation() *CattleGroupMutation {
	return cguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cguo *CattleGroupUpdateOne) Select(field string, fields ...string) *CattleGroupUpdateOne {
	cguo.fields = append([]string{field}, fields...)
	return cguo
}

// Save executes the query and returns the updated CattleGroup entity.
func (cguo *CattleGroupUpdateOne) Save(ctx context.Context) (*CattleGroup, error) {
	var (
		err  error
		node *CattleGroup
	)
	if len(cguo.hooks) == 0 {
		node, err = cguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CattleGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cguo.mutation = mutation
			node, err = cguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cguo.hooks) - 1; i >= 0; i-- {
			if cguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CattleGroup)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CattleGroupMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cguo *CattleGroupUpdateOne) SaveX(ctx context.Context) *CattleGroup {
	node, err := cguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cguo *CattleGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := cguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cguo *CattleGroupUpdateOne) ExecX(ctx context.Context) {
	if err := cguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cguo *CattleGroupUpdateOne) sqlSave(ctx context.Context) (_node *CattleGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cattlegroup.Table,
			Columns: cattlegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cattlegroup.FieldID,
			},
		},
	}
	id, ok := cguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CattleGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cattlegroup.FieldID)
		for _, f := range fields {
			if !cattlegroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cattlegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cguo.mutation.EarNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlegroup.FieldEarNumber,
		})
	}
	if value, ok := cguo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldDate,
		})
	}
	if value, ok := cguo.mutation.AddedDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldDate,
		})
	}
	if value, ok := cguo.mutation.ToShed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlegroup.FieldToShed,
		})
	}
	if value, ok := cguo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldTenantId,
		})
	}
	if value, ok := cguo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldTenantId,
		})
	}
	if value, ok := cguo.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlegroup.FieldTenantName,
		})
	}
	if value, ok := cguo.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlegroup.FieldRemarks,
		})
	}
	if value, ok := cguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldCreatedAt,
		})
	}
	if value, ok := cguo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldCreatedAt,
		})
	}
	if value, ok := cguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldUpdatedAt,
		})
	}
	if value, ok := cguo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlegroup.FieldUpdatedAt,
		})
	}
	if value, ok := cguo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cattlegroup.FieldDeleted,
		})
	}
	if value, ok := cguo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cattlegroup.FieldDeleted,
		})
	}
	_node = &CattleGroup{config: cguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cattlegroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
