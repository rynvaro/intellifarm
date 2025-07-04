// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/position"
	"cattleai/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PositionUpdate is the builder for updating Position entities.
type PositionUpdate struct {
	config
	hooks    []Hook
	mutation *PositionMutation
}

// Where appends a list predicates to the PositionUpdate builder.
func (pu *PositionUpdate) Where(ps ...predicate.Position) *PositionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PositionUpdate) SetName(s string) *PositionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetCode sets the "code" field.
func (pu *PositionUpdate) SetCode(s string) *PositionUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetTenantId sets the "tenantId" field.
func (pu *PositionUpdate) SetTenantId(i int64) *PositionUpdate {
	pu.mutation.ResetTenantId()
	pu.mutation.SetTenantId(i)
	return pu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableTenantId(i *int64) *PositionUpdate {
	if i != nil {
		pu.SetTenantId(*i)
	}
	return pu
}

// AddTenantId adds i to the "tenantId" field.
func (pu *PositionUpdate) AddTenantId(i int64) *PositionUpdate {
	pu.mutation.AddTenantId(i)
	return pu
}

// SetTenantName sets the "tenantName" field.
func (pu *PositionUpdate) SetTenantName(s string) *PositionUpdate {
	pu.mutation.SetTenantName(s)
	return pu
}

// SetNillableTenantName sets the "tenantName" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableTenantName(s *string) *PositionUpdate {
	if s != nil {
		pu.SetTenantName(*s)
	}
	return pu
}

// SetOrder sets the "order" field.
func (pu *PositionUpdate) SetOrder(i int) *PositionUpdate {
	pu.mutation.ResetOrder()
	pu.mutation.SetOrder(i)
	return pu
}

// AddOrder adds i to the "order" field.
func (pu *PositionUpdate) AddOrder(i int) *PositionUpdate {
	pu.mutation.AddOrder(i)
	return pu
}

// SetRemarks sets the "remarks" field.
func (pu *PositionUpdate) SetRemarks(s string) *PositionUpdate {
	pu.mutation.SetRemarks(s)
	return pu
}

// SetCreatedAt sets the "createdAt" field.
func (pu *PositionUpdate) SetCreatedAt(i int64) *PositionUpdate {
	pu.mutation.ResetCreatedAt()
	pu.mutation.SetCreatedAt(i)
	return pu
}

// AddCreatedAt adds i to the "createdAt" field.
func (pu *PositionUpdate) AddCreatedAt(i int64) *PositionUpdate {
	pu.mutation.AddCreatedAt(i)
	return pu
}

// SetUpdatedAt sets the "updatedAt" field.
func (pu *PositionUpdate) SetUpdatedAt(i int64) *PositionUpdate {
	pu.mutation.ResetUpdatedAt()
	pu.mutation.SetUpdatedAt(i)
	return pu
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (pu *PositionUpdate) AddUpdatedAt(i int64) *PositionUpdate {
	pu.mutation.AddUpdatedAt(i)
	return pu
}

// SetDeleted sets the "deleted" field.
func (pu *PositionUpdate) SetDeleted(i int) *PositionUpdate {
	pu.mutation.ResetDeleted()
	pu.mutation.SetDeleted(i)
	return pu
}

// AddDeleted adds i to the "deleted" field.
func (pu *PositionUpdate) AddDeleted(i int) *PositionUpdate {
	pu.mutation.AddDeleted(i)
	return pu
}

// Mutation returns the PositionMutation object of the builder.
func (pu *PositionUpdate) Mutation() *PositionMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PositionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PositionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PositionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PositionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PositionUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	return nil
}

func (pu *PositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   position.Table,
			Columns: position.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: position.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldName,
		})
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldCode,
		})
	}
	if value, ok := pu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldTenantId,
		})
	}
	if value, ok := pu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldTenantId,
		})
	}
	if value, ok := pu.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldTenantName,
		})
	}
	if value, ok := pu.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: position.FieldOrder,
		})
	}
	if value, ok := pu.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: position.FieldOrder,
		})
	}
	if value, ok := pu.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldRemarks,
		})
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: position.FieldDeleted,
		})
	}
	if value, ok := pu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: position.FieldDeleted,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PositionUpdateOne is the builder for updating a single Position entity.
type PositionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PositionMutation
}

// SetName sets the "name" field.
func (puo *PositionUpdateOne) SetName(s string) *PositionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetCode sets the "code" field.
func (puo *PositionUpdateOne) SetCode(s string) *PositionUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetTenantId sets the "tenantId" field.
func (puo *PositionUpdateOne) SetTenantId(i int64) *PositionUpdateOne {
	puo.mutation.ResetTenantId()
	puo.mutation.SetTenantId(i)
	return puo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableTenantId(i *int64) *PositionUpdateOne {
	if i != nil {
		puo.SetTenantId(*i)
	}
	return puo
}

// AddTenantId adds i to the "tenantId" field.
func (puo *PositionUpdateOne) AddTenantId(i int64) *PositionUpdateOne {
	puo.mutation.AddTenantId(i)
	return puo
}

// SetTenantName sets the "tenantName" field.
func (puo *PositionUpdateOne) SetTenantName(s string) *PositionUpdateOne {
	puo.mutation.SetTenantName(s)
	return puo
}

// SetNillableTenantName sets the "tenantName" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableTenantName(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetTenantName(*s)
	}
	return puo
}

// SetOrder sets the "order" field.
func (puo *PositionUpdateOne) SetOrder(i int) *PositionUpdateOne {
	puo.mutation.ResetOrder()
	puo.mutation.SetOrder(i)
	return puo
}

// AddOrder adds i to the "order" field.
func (puo *PositionUpdateOne) AddOrder(i int) *PositionUpdateOne {
	puo.mutation.AddOrder(i)
	return puo
}

// SetRemarks sets the "remarks" field.
func (puo *PositionUpdateOne) SetRemarks(s string) *PositionUpdateOne {
	puo.mutation.SetRemarks(s)
	return puo
}

// SetCreatedAt sets the "createdAt" field.
func (puo *PositionUpdateOne) SetCreatedAt(i int64) *PositionUpdateOne {
	puo.mutation.ResetCreatedAt()
	puo.mutation.SetCreatedAt(i)
	return puo
}

// AddCreatedAt adds i to the "createdAt" field.
func (puo *PositionUpdateOne) AddCreatedAt(i int64) *PositionUpdateOne {
	puo.mutation.AddCreatedAt(i)
	return puo
}

// SetUpdatedAt sets the "updatedAt" field.
func (puo *PositionUpdateOne) SetUpdatedAt(i int64) *PositionUpdateOne {
	puo.mutation.ResetUpdatedAt()
	puo.mutation.SetUpdatedAt(i)
	return puo
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (puo *PositionUpdateOne) AddUpdatedAt(i int64) *PositionUpdateOne {
	puo.mutation.AddUpdatedAt(i)
	return puo
}

// SetDeleted sets the "deleted" field.
func (puo *PositionUpdateOne) SetDeleted(i int) *PositionUpdateOne {
	puo.mutation.ResetDeleted()
	puo.mutation.SetDeleted(i)
	return puo
}

// AddDeleted adds i to the "deleted" field.
func (puo *PositionUpdateOne) AddDeleted(i int) *PositionUpdateOne {
	puo.mutation.AddDeleted(i)
	return puo
}

// Mutation returns the PositionMutation object of the builder.
func (puo *PositionUpdateOne) Mutation() *PositionMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PositionUpdateOne) Select(field string, fields ...string) *PositionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Position entity.
func (puo *PositionUpdateOne) Save(ctx context.Context) (*Position, error) {
	var (
		err  error
		node *Position
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Position)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PositionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PositionUpdateOne) SaveX(ctx context.Context) *Position {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PositionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PositionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PositionUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	return nil
}

func (puo *PositionUpdateOne) sqlSave(ctx context.Context) (_node *Position, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   position.Table,
			Columns: position.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: position.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Position.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, position.FieldID)
		for _, f := range fields {
			if !position.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != position.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldName,
		})
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldCode,
		})
	}
	if value, ok := puo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldTenantId,
		})
	}
	if value, ok := puo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldTenantId,
		})
	}
	if value, ok := puo.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldTenantName,
		})
	}
	if value, ok := puo.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: position.FieldOrder,
		})
	}
	if value, ok := puo.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: position.FieldOrder,
		})
	}
	if value, ok := puo.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldRemarks,
		})
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: position.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: position.FieldDeleted,
		})
	}
	if value, ok := puo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: position.FieldDeleted,
		})
	}
	_node = &Position{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
