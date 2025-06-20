// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/predicate"
	"cattleai/ent/winddirection"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WindDirectionUpdate is the builder for updating WindDirection entities.
type WindDirectionUpdate struct {
	config
	hooks    []Hook
	mutation *WindDirectionMutation
}

// Where appends a list predicates to the WindDirectionUpdate builder.
func (wdu *WindDirectionUpdate) Where(ps ...predicate.WindDirection) *WindDirectionUpdate {
	wdu.mutation.Where(ps...)
	return wdu
}

// SetName sets the "name" field.
func (wdu *WindDirectionUpdate) SetName(s string) *WindDirectionUpdate {
	wdu.mutation.SetName(s)
	return wdu
}

// Mutation returns the WindDirectionMutation object of the builder.
func (wdu *WindDirectionUpdate) Mutation() *WindDirectionMutation {
	return wdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wdu *WindDirectionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wdu.hooks) == 0 {
		if err = wdu.check(); err != nil {
			return 0, err
		}
		affected, err = wdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WindDirectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wdu.check(); err != nil {
				return 0, err
			}
			wdu.mutation = mutation
			affected, err = wdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wdu.hooks) - 1; i >= 0; i-- {
			if wdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wdu *WindDirectionUpdate) SaveX(ctx context.Context) int {
	affected, err := wdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wdu *WindDirectionUpdate) Exec(ctx context.Context) error {
	_, err := wdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdu *WindDirectionUpdate) ExecX(ctx context.Context) {
	if err := wdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wdu *WindDirectionUpdate) check() error {
	if v, ok := wdu.mutation.Name(); ok {
		if err := winddirection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WindDirection.name": %w`, err)}
		}
	}
	return nil
}

func (wdu *WindDirectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   winddirection.Table,
			Columns: winddirection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: winddirection.FieldID,
			},
		},
	}
	if ps := wdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wdu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: winddirection.FieldName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{winddirection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// WindDirectionUpdateOne is the builder for updating a single WindDirection entity.
type WindDirectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WindDirectionMutation
}

// SetName sets the "name" field.
func (wduo *WindDirectionUpdateOne) SetName(s string) *WindDirectionUpdateOne {
	wduo.mutation.SetName(s)
	return wduo
}

// Mutation returns the WindDirectionMutation object of the builder.
func (wduo *WindDirectionUpdateOne) Mutation() *WindDirectionMutation {
	return wduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wduo *WindDirectionUpdateOne) Select(field string, fields ...string) *WindDirectionUpdateOne {
	wduo.fields = append([]string{field}, fields...)
	return wduo
}

// Save executes the query and returns the updated WindDirection entity.
func (wduo *WindDirectionUpdateOne) Save(ctx context.Context) (*WindDirection, error) {
	var (
		err  error
		node *WindDirection
	)
	if len(wduo.hooks) == 0 {
		if err = wduo.check(); err != nil {
			return nil, err
		}
		node, err = wduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WindDirectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wduo.check(); err != nil {
				return nil, err
			}
			wduo.mutation = mutation
			node, err = wduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wduo.hooks) - 1; i >= 0; i-- {
			if wduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*WindDirection)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WindDirectionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wduo *WindDirectionUpdateOne) SaveX(ctx context.Context) *WindDirection {
	node, err := wduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wduo *WindDirectionUpdateOne) Exec(ctx context.Context) error {
	_, err := wduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wduo *WindDirectionUpdateOne) ExecX(ctx context.Context) {
	if err := wduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wduo *WindDirectionUpdateOne) check() error {
	if v, ok := wduo.mutation.Name(); ok {
		if err := winddirection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WindDirection.name": %w`, err)}
		}
	}
	return nil
}

func (wduo *WindDirectionUpdateOne) sqlSave(ctx context.Context) (_node *WindDirection, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   winddirection.Table,
			Columns: winddirection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: winddirection.FieldID,
			},
		},
	}
	id, ok := wduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WindDirection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, winddirection.FieldID)
		for _, f := range fields {
			if !winddirection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != winddirection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: winddirection.FieldName,
		})
	}
	_node = &WindDirection{config: wduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{winddirection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
