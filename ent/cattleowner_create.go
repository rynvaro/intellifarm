// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/cattleowner"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CattleOwnerCreate is the builder for creating a CattleOwner entity.
type CattleOwnerCreate struct {
	config
	mutation *CattleOwnerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (coc *CattleOwnerCreate) SetName(s string) *CattleOwnerCreate {
	coc.mutation.SetName(s)
	return coc
}

// Mutation returns the CattleOwnerMutation object of the builder.
func (coc *CattleOwnerCreate) Mutation() *CattleOwnerMutation {
	return coc.mutation
}

// Save creates the CattleOwner in the database.
func (coc *CattleOwnerCreate) Save(ctx context.Context) (*CattleOwner, error) {
	var (
		err  error
		node *CattleOwner
	)
	if len(coc.hooks) == 0 {
		if err = coc.check(); err != nil {
			return nil, err
		}
		node, err = coc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CattleOwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = coc.check(); err != nil {
				return nil, err
			}
			coc.mutation = mutation
			if node, err = coc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(coc.hooks) - 1; i >= 0; i-- {
			if coc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = coc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, coc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CattleOwner)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CattleOwnerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (coc *CattleOwnerCreate) SaveX(ctx context.Context) *CattleOwner {
	v, err := coc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (coc *CattleOwnerCreate) Exec(ctx context.Context) error {
	_, err := coc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (coc *CattleOwnerCreate) ExecX(ctx context.Context) {
	if err := coc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (coc *CattleOwnerCreate) check() error {
	if _, ok := coc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CattleOwner.name"`)}
	}
	if v, ok := coc.mutation.Name(); ok {
		if err := cattleowner.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CattleOwner.name": %w`, err)}
		}
	}
	return nil
}

func (coc *CattleOwnerCreate) sqlSave(ctx context.Context) (*CattleOwner, error) {
	_node, _spec := coc.createSpec()
	if err := sqlgraph.CreateNode(ctx, coc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (coc *CattleOwnerCreate) createSpec() (*CattleOwner, *sqlgraph.CreateSpec) {
	var (
		_node = &CattleOwner{config: coc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cattleowner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cattleowner.FieldID,
			},
		}
	)
	if value, ok := coc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattleowner.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// CattleOwnerCreateBulk is the builder for creating many CattleOwner entities in bulk.
type CattleOwnerCreateBulk struct {
	config
	builders []*CattleOwnerCreate
}

// Save creates the CattleOwner entities in the database.
func (cocb *CattleOwnerCreateBulk) Save(ctx context.Context) ([]*CattleOwner, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cocb.builders))
	nodes := make([]*CattleOwner, len(cocb.builders))
	mutators := make([]Mutator, len(cocb.builders))
	for i := range cocb.builders {
		func(i int, root context.Context) {
			builder := cocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CattleOwnerMutation)
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
					_, err = mutators[i+1].Mutate(root, cocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cocb *CattleOwnerCreateBulk) SaveX(ctx context.Context) []*CattleOwner {
	v, err := cocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cocb *CattleOwnerCreateBulk) Exec(ctx context.Context) error {
	_, err := cocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cocb *CattleOwnerCreateBulk) ExecX(ctx context.Context) {
	if err := cocb.Exec(ctx); err != nil {
		panic(err)
	}
}
