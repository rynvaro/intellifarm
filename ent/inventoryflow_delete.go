// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/inventoryflow"
	"cattleai/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InventoryFlowDelete is the builder for deleting a InventoryFlow entity.
type InventoryFlowDelete struct {
	config
	hooks    []Hook
	mutation *InventoryFlowMutation
}

// Where appends a list predicates to the InventoryFlowDelete builder.
func (ifd *InventoryFlowDelete) Where(ps ...predicate.InventoryFlow) *InventoryFlowDelete {
	ifd.mutation.Where(ps...)
	return ifd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ifd *InventoryFlowDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ifd.hooks) == 0 {
		affected, err = ifd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InventoryFlowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ifd.mutation = mutation
			affected, err = ifd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ifd.hooks) - 1; i >= 0; i-- {
			if ifd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ifd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ifd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifd *InventoryFlowDelete) ExecX(ctx context.Context) int {
	n, err := ifd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ifd *InventoryFlowDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: inventoryflow.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inventoryflow.FieldID,
			},
		},
	}
	if ps := ifd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ifd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// InventoryFlowDeleteOne is the builder for deleting a single InventoryFlow entity.
type InventoryFlowDeleteOne struct {
	ifd *InventoryFlowDelete
}

// Exec executes the deletion query.
func (ifdo *InventoryFlowDeleteOne) Exec(ctx context.Context) error {
	n, err := ifdo.ifd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{inventoryflow.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ifdo *InventoryFlowDeleteOne) ExecX(ctx context.Context) {
	ifdo.ifd.ExecX(ctx)
}
