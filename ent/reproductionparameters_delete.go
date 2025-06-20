// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/predicate"
	"cattleai/ent/reproductionparameters"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReproductionParametersDelete is the builder for deleting a ReproductionParameters entity.
type ReproductionParametersDelete struct {
	config
	hooks    []Hook
	mutation *ReproductionParametersMutation
}

// Where appends a list predicates to the ReproductionParametersDelete builder.
func (rpd *ReproductionParametersDelete) Where(ps ...predicate.ReproductionParameters) *ReproductionParametersDelete {
	rpd.mutation.Where(ps...)
	return rpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rpd *ReproductionParametersDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rpd.hooks) == 0 {
		affected, err = rpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReproductionParametersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rpd.mutation = mutation
			affected, err = rpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rpd.hooks) - 1; i >= 0; i-- {
			if rpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpd *ReproductionParametersDelete) ExecX(ctx context.Context) int {
	n, err := rpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rpd *ReproductionParametersDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: reproductionparameters.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reproductionparameters.FieldID,
			},
		},
	}
	if ps := rpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ReproductionParametersDeleteOne is the builder for deleting a single ReproductionParameters entity.
type ReproductionParametersDeleteOne struct {
	rpd *ReproductionParametersDelete
}

// Exec executes the deletion query.
func (rpdo *ReproductionParametersDeleteOne) Exec(ctx context.Context) error {
	n, err := rpdo.rpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{reproductionparameters.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rpdo *ReproductionParametersDeleteOne) ExecX(ctx context.Context) {
	rpdo.rpd.ExecX(ctx)
}
