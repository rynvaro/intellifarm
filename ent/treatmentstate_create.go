// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/treatmentstate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TreatmentStateCreate is the builder for creating a TreatmentState entity.
type TreatmentStateCreate struct {
	config
	mutation *TreatmentStateMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tsc *TreatmentStateCreate) SetName(s string) *TreatmentStateCreate {
	tsc.mutation.SetName(s)
	return tsc
}

// Mutation returns the TreatmentStateMutation object of the builder.
func (tsc *TreatmentStateCreate) Mutation() *TreatmentStateMutation {
	return tsc.mutation
}

// Save creates the TreatmentState in the database.
func (tsc *TreatmentStateCreate) Save(ctx context.Context) (*TreatmentState, error) {
	var (
		err  error
		node *TreatmentState
	)
	if len(tsc.hooks) == 0 {
		if err = tsc.check(); err != nil {
			return nil, err
		}
		node, err = tsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TreatmentStateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tsc.check(); err != nil {
				return nil, err
			}
			tsc.mutation = mutation
			if node, err = tsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tsc.hooks) - 1; i >= 0; i-- {
			if tsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tsc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tsc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TreatmentState)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TreatmentStateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tsc *TreatmentStateCreate) SaveX(ctx context.Context) *TreatmentState {
	v, err := tsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tsc *TreatmentStateCreate) Exec(ctx context.Context) error {
	_, err := tsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsc *TreatmentStateCreate) ExecX(ctx context.Context) {
	if err := tsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsc *TreatmentStateCreate) check() error {
	if _, ok := tsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TreatmentState.name"`)}
	}
	if v, ok := tsc.mutation.Name(); ok {
		if err := treatmentstate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TreatmentState.name": %w`, err)}
		}
	}
	return nil
}

func (tsc *TreatmentStateCreate) sqlSave(ctx context.Context) (*TreatmentState, error) {
	_node, _spec := tsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tsc *TreatmentStateCreate) createSpec() (*TreatmentState, *sqlgraph.CreateSpec) {
	var (
		_node = &TreatmentState{config: tsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: treatmentstate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: treatmentstate.FieldID,
			},
		}
	)
	if value, ok := tsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatmentstate.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// TreatmentStateCreateBulk is the builder for creating many TreatmentState entities in bulk.
type TreatmentStateCreateBulk struct {
	config
	builders []*TreatmentStateCreate
}

// Save creates the TreatmentState entities in the database.
func (tscb *TreatmentStateCreateBulk) Save(ctx context.Context) ([]*TreatmentState, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tscb.builders))
	nodes := make([]*TreatmentState, len(tscb.builders))
	mutators := make([]Mutator, len(tscb.builders))
	for i := range tscb.builders {
		func(i int, root context.Context) {
			builder := tscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TreatmentStateMutation)
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
					_, err = mutators[i+1].Mutate(root, tscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tscb *TreatmentStateCreateBulk) SaveX(ctx context.Context) []*TreatmentState {
	v, err := tscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tscb *TreatmentStateCreateBulk) Exec(ctx context.Context) error {
	_, err := tscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tscb *TreatmentStateCreateBulk) ExecX(ctx context.Context) {
	if err := tscb.Exec(ctx); err != nil {
		panic(err)
	}
}
