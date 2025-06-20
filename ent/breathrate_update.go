// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/breathrate"
	"cattleai/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BreathRateUpdate is the builder for updating BreathRate entities.
type BreathRateUpdate struct {
	config
	hooks    []Hook
	mutation *BreathRateMutation
}

// Where appends a list predicates to the BreathRateUpdate builder.
func (bru *BreathRateUpdate) Where(ps ...predicate.BreathRate) *BreathRateUpdate {
	bru.mutation.Where(ps...)
	return bru
}

// SetName sets the "name" field.
func (bru *BreathRateUpdate) SetName(s string) *BreathRateUpdate {
	bru.mutation.SetName(s)
	return bru
}

// Mutation returns the BreathRateMutation object of the builder.
func (bru *BreathRateUpdate) Mutation() *BreathRateMutation {
	return bru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bru *BreathRateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bru.hooks) == 0 {
		if err = bru.check(); err != nil {
			return 0, err
		}
		affected, err = bru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BreathRateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bru.check(); err != nil {
				return 0, err
			}
			bru.mutation = mutation
			affected, err = bru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bru.hooks) - 1; i >= 0; i-- {
			if bru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bru *BreathRateUpdate) SaveX(ctx context.Context) int {
	affected, err := bru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bru *BreathRateUpdate) Exec(ctx context.Context) error {
	_, err := bru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bru *BreathRateUpdate) ExecX(ctx context.Context) {
	if err := bru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bru *BreathRateUpdate) check() error {
	if v, ok := bru.mutation.Name(); ok {
		if err := breathrate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "BreathRate.name": %w`, err)}
		}
	}
	return nil
}

func (bru *BreathRateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   breathrate.Table,
			Columns: breathrate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: breathrate.FieldID,
			},
		},
	}
	if ps := bru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: breathrate.FieldName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{breathrate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BreathRateUpdateOne is the builder for updating a single BreathRate entity.
type BreathRateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BreathRateMutation
}

// SetName sets the "name" field.
func (bruo *BreathRateUpdateOne) SetName(s string) *BreathRateUpdateOne {
	bruo.mutation.SetName(s)
	return bruo
}

// Mutation returns the BreathRateMutation object of the builder.
func (bruo *BreathRateUpdateOne) Mutation() *BreathRateMutation {
	return bruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bruo *BreathRateUpdateOne) Select(field string, fields ...string) *BreathRateUpdateOne {
	bruo.fields = append([]string{field}, fields...)
	return bruo
}

// Save executes the query and returns the updated BreathRate entity.
func (bruo *BreathRateUpdateOne) Save(ctx context.Context) (*BreathRate, error) {
	var (
		err  error
		node *BreathRate
	)
	if len(bruo.hooks) == 0 {
		if err = bruo.check(); err != nil {
			return nil, err
		}
		node, err = bruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BreathRateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bruo.check(); err != nil {
				return nil, err
			}
			bruo.mutation = mutation
			node, err = bruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bruo.hooks) - 1; i >= 0; i-- {
			if bruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*BreathRate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BreathRateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bruo *BreathRateUpdateOne) SaveX(ctx context.Context) *BreathRate {
	node, err := bruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bruo *BreathRateUpdateOne) Exec(ctx context.Context) error {
	_, err := bruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bruo *BreathRateUpdateOne) ExecX(ctx context.Context) {
	if err := bruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bruo *BreathRateUpdateOne) check() error {
	if v, ok := bruo.mutation.Name(); ok {
		if err := breathrate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "BreathRate.name": %w`, err)}
		}
	}
	return nil
}

func (bruo *BreathRateUpdateOne) sqlSave(ctx context.Context) (_node *BreathRate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   breathrate.Table,
			Columns: breathrate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: breathrate.FieldID,
			},
		},
	}
	id, ok := bruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BreathRate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, breathrate.FieldID)
		for _, f := range fields {
			if !breathrate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != breathrate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: breathrate.FieldName,
		})
	}
	_node = &BreathRate{config: bruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{breathrate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
