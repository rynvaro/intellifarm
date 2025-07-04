// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/abortionreason"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AbortionReasonCreate is the builder for creating a AbortionReason entity.
type AbortionReasonCreate struct {
	config
	mutation *AbortionReasonMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (arc *AbortionReasonCreate) SetName(s string) *AbortionReasonCreate {
	arc.mutation.SetName(s)
	return arc
}

// SetCode sets the "code" field.
func (arc *AbortionReasonCreate) SetCode(s string) *AbortionReasonCreate {
	arc.mutation.SetCode(s)
	return arc
}

// SetTenantId sets the "tenantId" field.
func (arc *AbortionReasonCreate) SetTenantId(i int64) *AbortionReasonCreate {
	arc.mutation.SetTenantId(i)
	return arc
}

// SetTenantName sets the "tenantName" field.
func (arc *AbortionReasonCreate) SetTenantName(s string) *AbortionReasonCreate {
	arc.mutation.SetTenantName(s)
	return arc
}

// SetOrder sets the "order" field.
func (arc *AbortionReasonCreate) SetOrder(i int) *AbortionReasonCreate {
	arc.mutation.SetOrder(i)
	return arc
}

// SetRemarks sets the "remarks" field.
func (arc *AbortionReasonCreate) SetRemarks(s string) *AbortionReasonCreate {
	arc.mutation.SetRemarks(s)
	return arc
}

// SetCreatedAt sets the "createdAt" field.
func (arc *AbortionReasonCreate) SetCreatedAt(i int64) *AbortionReasonCreate {
	arc.mutation.SetCreatedAt(i)
	return arc
}

// SetUpdatedAt sets the "updatedAt" field.
func (arc *AbortionReasonCreate) SetUpdatedAt(i int64) *AbortionReasonCreate {
	arc.mutation.SetUpdatedAt(i)
	return arc
}

// SetDeleted sets the "deleted" field.
func (arc *AbortionReasonCreate) SetDeleted(i int) *AbortionReasonCreate {
	arc.mutation.SetDeleted(i)
	return arc
}

// Mutation returns the AbortionReasonMutation object of the builder.
func (arc *AbortionReasonCreate) Mutation() *AbortionReasonMutation {
	return arc.mutation
}

// Save creates the AbortionReason in the database.
func (arc *AbortionReasonCreate) Save(ctx context.Context) (*AbortionReason, error) {
	var (
		err  error
		node *AbortionReason
	)
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AbortionReasonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			if node, err = arc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			if arc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = arc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, arc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AbortionReason)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AbortionReasonMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AbortionReasonCreate) SaveX(ctx context.Context) *AbortionReason {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AbortionReasonCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AbortionReasonCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AbortionReasonCreate) check() error {
	if _, ok := arc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AbortionReason.name"`)}
	}
	if _, ok := arc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "AbortionReason.code"`)}
	}
	if _, ok := arc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "AbortionReason.tenantId"`)}
	}
	if _, ok := arc.mutation.TenantName(); !ok {
		return &ValidationError{Name: "tenantName", err: errors.New(`ent: missing required field "AbortionReason.tenantName"`)}
	}
	if _, ok := arc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "AbortionReason.order"`)}
	}
	if _, ok := arc.mutation.Remarks(); !ok {
		return &ValidationError{Name: "remarks", err: errors.New(`ent: missing required field "AbortionReason.remarks"`)}
	}
	if _, ok := arc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "AbortionReason.createdAt"`)}
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "AbortionReason.updatedAt"`)}
	}
	if _, ok := arc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "AbortionReason.deleted"`)}
	}
	return nil
}

func (arc *AbortionReasonCreate) sqlSave(ctx context.Context) (*AbortionReason, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (arc *AbortionReasonCreate) createSpec() (*AbortionReason, *sqlgraph.CreateSpec) {
	var (
		_node = &AbortionReason{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: abortionreason.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: abortionreason.FieldID,
			},
		}
	)
	if value, ok := arc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: abortionreason.FieldName,
		})
		_node.Name = value
	}
	if value, ok := arc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: abortionreason.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := arc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: abortionreason.FieldTenantId,
		})
		_node.TenantId = value
	}
	if value, ok := arc.mutation.TenantName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: abortionreason.FieldTenantName,
		})
		_node.TenantName = value
	}
	if value, ok := arc.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abortionreason.FieldOrder,
		})
		_node.Order = value
	}
	if value, ok := arc.mutation.Remarks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: abortionreason.FieldRemarks,
		})
		_node.Remarks = value
	}
	if value, ok := arc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: abortionreason.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := arc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: abortionreason.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := arc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abortionreason.FieldDeleted,
		})
		_node.Deleted = value
	}
	return _node, _spec
}

// AbortionReasonCreateBulk is the builder for creating many AbortionReason entities in bulk.
type AbortionReasonCreateBulk struct {
	config
	builders []*AbortionReasonCreate
}

// Save creates the AbortionReason entities in the database.
func (arcb *AbortionReasonCreateBulk) Save(ctx context.Context) ([]*AbortionReason, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AbortionReason, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AbortionReasonMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AbortionReasonCreateBulk) SaveX(ctx context.Context) []*AbortionReason {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AbortionReasonCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AbortionReasonCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}
