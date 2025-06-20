// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/predicate"
	"cattleai/ent/shedtype"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShedTypeUpdate is the builder for updating ShedType entities.
type ShedTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ShedTypeMutation
}

// Where appends a list predicates to the ShedTypeUpdate builder.
func (stu *ShedTypeUpdate) Where(ps ...predicate.ShedType) *ShedTypeUpdate {
	stu.mutation.Where(ps...)
	return stu
}

// SetName sets the "name" field.
func (stu *ShedTypeUpdate) SetName(s string) *ShedTypeUpdate {
	stu.mutation.SetName(s)
	return stu
}

// SetCode sets the "code" field.
func (stu *ShedTypeUpdate) SetCode(s string) *ShedTypeUpdate {
	stu.mutation.SetCode(s)
	return stu
}

// SetTenantId sets the "tenantId" field.
func (stu *ShedTypeUpdate) SetTenantId(i int64) *ShedTypeUpdate {
	stu.mutation.ResetTenantId()
	stu.mutation.SetTenantId(i)
	return stu
}

// AddTenantId adds i to the "tenantId" field.
func (stu *ShedTypeUpdate) AddTenantId(i int64) *ShedTypeUpdate {
	stu.mutation.AddTenantId(i)
	return stu
}

// SetTenantName sets the "tenantName" field.
func (stu *ShedTypeUpdate) SetTenantName(s string) *ShedTypeUpdate {
	stu.mutation.SetTenantName(s)
	return stu
}

// SetOrder sets the "order" field.
func (stu *ShedTypeUpdate) SetOrder(i int) *ShedTypeUpdate {
	stu.mutation.ResetOrder()
	stu.mutation.SetOrder(i)
	return stu
}

// AddOrder adds i to the "order" field.
func (stu *ShedTypeUpdate) AddOrder(i int) *ShedTypeUpdate {
	stu.mutation.AddOrder(i)
	return stu
}

// SetRemarks sets the "remarks" field.
func (stu *ShedTypeUpdate) SetRemarks(s string) *ShedTypeUpdate {
	stu.mutation.SetRemarks(s)
	return stu
}

// SetCreatedAt sets the "createdAt" field.
func (stu *ShedTypeUpdate) SetCreatedAt(i int64) *ShedTypeUpdate {
	stu.mutation.ResetCreatedAt()
	stu.mutation.SetCreatedAt(i)
	return stu
}

// AddCreatedAt adds i to the "createdAt" field.
func (stu *ShedTypeUpdate) AddCreatedAt(i int64) *ShedTypeUpdate {
	stu.mutation.AddCreatedAt(i)
	return stu
}

// SetUpdatedAt sets the "updatedAt" field.
func (stu *ShedTypeUpdate) SetUpdatedAt(i int64) *ShedTypeUpdate {
	stu.mutation.ResetUpdatedAt()
	stu.mutation.SetUpdatedAt(i)
	return stu
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (stu *ShedTypeUpdate) AddUpdatedAt(i int64) *ShedTypeUpdate {
	stu.mutation.AddUpdatedAt(i)
	return stu
}

// SetDeleted sets the "deleted" field.
func (stu *ShedTypeUpdate) SetDeleted(i int) *ShedTypeUpdate {
	stu.mutation.ResetDeleted()
	stu.mutation.SetDeleted(i)
	return stu
}

// AddDeleted adds i to the "deleted" field.
func (stu *ShedTypeUpdate) AddDeleted(i int) *ShedTypeUpdate {
	stu.mutation.AddDeleted(i)
	return stu
}

// Mutation returns the ShedTypeMutation object of the builder.
func (stu *ShedTypeUpdate) Mutation() *ShedTypeMutation {
	return stu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *ShedTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(stu.hooks) == 0 {
		affected, err = stu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShedTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stu.mutation = mutation
			affected, err = stu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(stu.hooks) - 1; i >= 0; i-- {
			if stu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (stu *ShedTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *ShedTypeUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *ShedTypeUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (stu *ShedTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shedtype.Table,
			Columns: shedtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shedtype.FieldID,
			},
		},
	}
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shedtype.FieldName,
		})
	}
	if value, ok := stu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shedtype.FieldCode,
		})
	}
	if value, ok := stu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldTenantId,
		})
	}
	if value, ok := stu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldTenantId,
		})
	}
	if value, ok := stu.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shedtype.FieldTenantName,
		})
	}
	if value, ok := stu.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shedtype.FieldOrder,
		})
	}
	if value, ok := stu.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shedtype.FieldOrder,
		})
	}
	if value, ok := stu.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shedtype.FieldRemarks,
		})
	}
	if value, ok := stu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldCreatedAt,
		})
	}
	if value, ok := stu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldCreatedAt,
		})
	}
	if value, ok := stu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldUpdatedAt,
		})
	}
	if value, ok := stu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldUpdatedAt,
		})
	}
	if value, ok := stu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shedtype.FieldDeleted,
		})
	}
	if value, ok := stu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shedtype.FieldDeleted,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shedtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ShedTypeUpdateOne is the builder for updating a single ShedType entity.
type ShedTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShedTypeMutation
}

// SetName sets the "name" field.
func (stuo *ShedTypeUpdateOne) SetName(s string) *ShedTypeUpdateOne {
	stuo.mutation.SetName(s)
	return stuo
}

// SetCode sets the "code" field.
func (stuo *ShedTypeUpdateOne) SetCode(s string) *ShedTypeUpdateOne {
	stuo.mutation.SetCode(s)
	return stuo
}

// SetTenantId sets the "tenantId" field.
func (stuo *ShedTypeUpdateOne) SetTenantId(i int64) *ShedTypeUpdateOne {
	stuo.mutation.ResetTenantId()
	stuo.mutation.SetTenantId(i)
	return stuo
}

// AddTenantId adds i to the "tenantId" field.
func (stuo *ShedTypeUpdateOne) AddTenantId(i int64) *ShedTypeUpdateOne {
	stuo.mutation.AddTenantId(i)
	return stuo
}

// SetTenantName sets the "tenantName" field.
func (stuo *ShedTypeUpdateOne) SetTenantName(s string) *ShedTypeUpdateOne {
	stuo.mutation.SetTenantName(s)
	return stuo
}

// SetOrder sets the "order" field.
func (stuo *ShedTypeUpdateOne) SetOrder(i int) *ShedTypeUpdateOne {
	stuo.mutation.ResetOrder()
	stuo.mutation.SetOrder(i)
	return stuo
}

// AddOrder adds i to the "order" field.
func (stuo *ShedTypeUpdateOne) AddOrder(i int) *ShedTypeUpdateOne {
	stuo.mutation.AddOrder(i)
	return stuo
}

// SetRemarks sets the "remarks" field.
func (stuo *ShedTypeUpdateOne) SetRemarks(s string) *ShedTypeUpdateOne {
	stuo.mutation.SetRemarks(s)
	return stuo
}

// SetCreatedAt sets the "createdAt" field.
func (stuo *ShedTypeUpdateOne) SetCreatedAt(i int64) *ShedTypeUpdateOne {
	stuo.mutation.ResetCreatedAt()
	stuo.mutation.SetCreatedAt(i)
	return stuo
}

// AddCreatedAt adds i to the "createdAt" field.
func (stuo *ShedTypeUpdateOne) AddCreatedAt(i int64) *ShedTypeUpdateOne {
	stuo.mutation.AddCreatedAt(i)
	return stuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (stuo *ShedTypeUpdateOne) SetUpdatedAt(i int64) *ShedTypeUpdateOne {
	stuo.mutation.ResetUpdatedAt()
	stuo.mutation.SetUpdatedAt(i)
	return stuo
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (stuo *ShedTypeUpdateOne) AddUpdatedAt(i int64) *ShedTypeUpdateOne {
	stuo.mutation.AddUpdatedAt(i)
	return stuo
}

// SetDeleted sets the "deleted" field.
func (stuo *ShedTypeUpdateOne) SetDeleted(i int) *ShedTypeUpdateOne {
	stuo.mutation.ResetDeleted()
	stuo.mutation.SetDeleted(i)
	return stuo
}

// AddDeleted adds i to the "deleted" field.
func (stuo *ShedTypeUpdateOne) AddDeleted(i int) *ShedTypeUpdateOne {
	stuo.mutation.AddDeleted(i)
	return stuo
}

// Mutation returns the ShedTypeMutation object of the builder.
func (stuo *ShedTypeUpdateOne) Mutation() *ShedTypeMutation {
	return stuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *ShedTypeUpdateOne) Select(field string, fields ...string) *ShedTypeUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated ShedType entity.
func (stuo *ShedTypeUpdateOne) Save(ctx context.Context) (*ShedType, error) {
	var (
		err  error
		node *ShedType
	)
	if len(stuo.hooks) == 0 {
		node, err = stuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShedTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stuo.mutation = mutation
			node, err = stuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(stuo.hooks) - 1; i >= 0; i-- {
			if stuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, stuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ShedType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ShedTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *ShedTypeUpdateOne) SaveX(ctx context.Context) *ShedType {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *ShedTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *ShedTypeUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (stuo *ShedTypeUpdateOne) sqlSave(ctx context.Context) (_node *ShedType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shedtype.Table,
			Columns: shedtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shedtype.FieldID,
			},
		},
	}
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ShedType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shedtype.FieldID)
		for _, f := range fields {
			if !shedtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shedtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shedtype.FieldName,
		})
	}
	if value, ok := stuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shedtype.FieldCode,
		})
	}
	if value, ok := stuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldTenantId,
		})
	}
	if value, ok := stuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldTenantId,
		})
	}
	if value, ok := stuo.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shedtype.FieldTenantName,
		})
	}
	if value, ok := stuo.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shedtype.FieldOrder,
		})
	}
	if value, ok := stuo.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shedtype.FieldOrder,
		})
	}
	if value, ok := stuo.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shedtype.FieldRemarks,
		})
	}
	if value, ok := stuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldCreatedAt,
		})
	}
	if value, ok := stuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldCreatedAt,
		})
	}
	if value, ok := stuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldUpdatedAt,
		})
	}
	if value, ok := stuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shedtype.FieldUpdatedAt,
		})
	}
	if value, ok := stuo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shedtype.FieldDeleted,
		})
	}
	if value, ok := stuo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shedtype.FieldDeleted,
		})
	}
	_node = &ShedType{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shedtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
