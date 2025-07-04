// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/cattlemove"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CattleMoveCreate is the builder for creating a CattleMove entity.
type CattleMoveCreate struct {
	config
	mutation *CattleMoveMutation
	hooks    []Hook
}

// SetCattleId sets the "cattleId" field.
func (cmc *CattleMoveCreate) SetCattleId(i int64) *CattleMoveCreate {
	cmc.mutation.SetCattleId(i)
	return cmc
}

// SetTenantId sets the "tenantId" field.
func (cmc *CattleMoveCreate) SetTenantId(i int64) *CattleMoveCreate {
	cmc.mutation.SetTenantId(i)
	return cmc
}

// SetTenantName sets the "tenantName" field.
func (cmc *CattleMoveCreate) SetTenantName(s string) *CattleMoveCreate {
	cmc.mutation.SetTenantName(s)
	return cmc
}

// SetFarmId sets the "farmId" field.
func (cmc *CattleMoveCreate) SetFarmId(i int64) *CattleMoveCreate {
	cmc.mutation.SetFarmId(i)
	return cmc
}

// SetFarmName sets the "farmName" field.
func (cmc *CattleMoveCreate) SetFarmName(s string) *CattleMoveCreate {
	cmc.mutation.SetFarmName(s)
	return cmc
}

// SetShedId sets the "shedId" field.
func (cmc *CattleMoveCreate) SetShedId(i int64) *CattleMoveCreate {
	cmc.mutation.SetShedId(i)
	return cmc
}

// SetShedName sets the "shedName" field.
func (cmc *CattleMoveCreate) SetShedName(s string) *CattleMoveCreate {
	cmc.mutation.SetShedName(s)
	return cmc
}

// SetEarNumber sets the "earNumber" field.
func (cmc *CattleMoveCreate) SetEarNumber(s string) *CattleMoveCreate {
	cmc.mutation.SetEarNumber(s)
	return cmc
}

// SetDate sets the "date" field.
func (cmc *CattleMoveCreate) SetDate(i int64) *CattleMoveCreate {
	cmc.mutation.SetDate(i)
	return cmc
}

// SetFromShedId sets the "fromShedId" field.
func (cmc *CattleMoveCreate) SetFromShedId(i int64) *CattleMoveCreate {
	cmc.mutation.SetFromShedId(i)
	return cmc
}

// SetFromShed sets the "fromShed" field.
func (cmc *CattleMoveCreate) SetFromShed(s string) *CattleMoveCreate {
	cmc.mutation.SetFromShed(s)
	return cmc
}

// SetToShedId sets the "toShedId" field.
func (cmc *CattleMoveCreate) SetToShedId(i int64) *CattleMoveCreate {
	cmc.mutation.SetToShedId(i)
	return cmc
}

// SetToShed sets the "toShed" field.
func (cmc *CattleMoveCreate) SetToShed(s string) *CattleMoveCreate {
	cmc.mutation.SetToShed(s)
	return cmc
}

// SetUserName sets the "userName" field.
func (cmc *CattleMoveCreate) SetUserName(s string) *CattleMoveCreate {
	cmc.mutation.SetUserName(s)
	return cmc
}

// SetReasonId sets the "reasonId" field.
func (cmc *CattleMoveCreate) SetReasonId(i int64) *CattleMoveCreate {
	cmc.mutation.SetReasonId(i)
	return cmc
}

// SetReasonName sets the "reasonName" field.
func (cmc *CattleMoveCreate) SetReasonName(s string) *CattleMoveCreate {
	cmc.mutation.SetReasonName(s)
	return cmc
}

// SetRemarks sets the "remarks" field.
func (cmc *CattleMoveCreate) SetRemarks(s string) *CattleMoveCreate {
	cmc.mutation.SetRemarks(s)
	return cmc
}

// SetCreatedAt sets the "createdAt" field.
func (cmc *CattleMoveCreate) SetCreatedAt(i int64) *CattleMoveCreate {
	cmc.mutation.SetCreatedAt(i)
	return cmc
}

// SetUpdatedAt sets the "updatedAt" field.
func (cmc *CattleMoveCreate) SetUpdatedAt(i int64) *CattleMoveCreate {
	cmc.mutation.SetUpdatedAt(i)
	return cmc
}

// SetDeleted sets the "deleted" field.
func (cmc *CattleMoveCreate) SetDeleted(i int) *CattleMoveCreate {
	cmc.mutation.SetDeleted(i)
	return cmc
}

// Mutation returns the CattleMoveMutation object of the builder.
func (cmc *CattleMoveCreate) Mutation() *CattleMoveMutation {
	return cmc.mutation
}

// Save creates the CattleMove in the database.
func (cmc *CattleMoveCreate) Save(ctx context.Context) (*CattleMove, error) {
	var (
		err  error
		node *CattleMove
	)
	if len(cmc.hooks) == 0 {
		if err = cmc.check(); err != nil {
			return nil, err
		}
		node, err = cmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CattleMoveMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cmc.check(); err != nil {
				return nil, err
			}
			cmc.mutation = mutation
			if node, err = cmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cmc.hooks) - 1; i >= 0; i-- {
			if cmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cmc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CattleMove)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CattleMoveMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cmc *CattleMoveCreate) SaveX(ctx context.Context) *CattleMove {
	v, err := cmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cmc *CattleMoveCreate) Exec(ctx context.Context) error {
	_, err := cmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmc *CattleMoveCreate) ExecX(ctx context.Context) {
	if err := cmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmc *CattleMoveCreate) check() error {
	if _, ok := cmc.mutation.CattleId(); !ok {
		return &ValidationError{Name: "cattleId", err: errors.New(`ent: missing required field "CattleMove.cattleId"`)}
	}
	if _, ok := cmc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "CattleMove.tenantId"`)}
	}
	if _, ok := cmc.mutation.TenantName(); !ok {
		return &ValidationError{Name: "tenantName", err: errors.New(`ent: missing required field "CattleMove.tenantName"`)}
	}
	if _, ok := cmc.mutation.FarmId(); !ok {
		return &ValidationError{Name: "farmId", err: errors.New(`ent: missing required field "CattleMove.farmId"`)}
	}
	if _, ok := cmc.mutation.FarmName(); !ok {
		return &ValidationError{Name: "farmName", err: errors.New(`ent: missing required field "CattleMove.farmName"`)}
	}
	if _, ok := cmc.mutation.ShedId(); !ok {
		return &ValidationError{Name: "shedId", err: errors.New(`ent: missing required field "CattleMove.shedId"`)}
	}
	if _, ok := cmc.mutation.ShedName(); !ok {
		return &ValidationError{Name: "shedName", err: errors.New(`ent: missing required field "CattleMove.shedName"`)}
	}
	if _, ok := cmc.mutation.EarNumber(); !ok {
		return &ValidationError{Name: "earNumber", err: errors.New(`ent: missing required field "CattleMove.earNumber"`)}
	}
	if _, ok := cmc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "CattleMove.date"`)}
	}
	if _, ok := cmc.mutation.FromShedId(); !ok {
		return &ValidationError{Name: "fromShedId", err: errors.New(`ent: missing required field "CattleMove.fromShedId"`)}
	}
	if _, ok := cmc.mutation.FromShed(); !ok {
		return &ValidationError{Name: "fromShed", err: errors.New(`ent: missing required field "CattleMove.fromShed"`)}
	}
	if _, ok := cmc.mutation.ToShedId(); !ok {
		return &ValidationError{Name: "toShedId", err: errors.New(`ent: missing required field "CattleMove.toShedId"`)}
	}
	if _, ok := cmc.mutation.ToShed(); !ok {
		return &ValidationError{Name: "toShed", err: errors.New(`ent: missing required field "CattleMove.toShed"`)}
	}
	if _, ok := cmc.mutation.UserName(); !ok {
		return &ValidationError{Name: "userName", err: errors.New(`ent: missing required field "CattleMove.userName"`)}
	}
	if _, ok := cmc.mutation.ReasonId(); !ok {
		return &ValidationError{Name: "reasonId", err: errors.New(`ent: missing required field "CattleMove.reasonId"`)}
	}
	if _, ok := cmc.mutation.ReasonName(); !ok {
		return &ValidationError{Name: "reasonName", err: errors.New(`ent: missing required field "CattleMove.reasonName"`)}
	}
	if _, ok := cmc.mutation.Remarks(); !ok {
		return &ValidationError{Name: "remarks", err: errors.New(`ent: missing required field "CattleMove.remarks"`)}
	}
	if _, ok := cmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "CattleMove.createdAt"`)}
	}
	if _, ok := cmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "CattleMove.updatedAt"`)}
	}
	if _, ok := cmc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "CattleMove.deleted"`)}
	}
	return nil
}

func (cmc *CattleMoveCreate) sqlSave(ctx context.Context) (*CattleMove, error) {
	_node, _spec := cmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cmc *CattleMoveCreate) createSpec() (*CattleMove, *sqlgraph.CreateSpec) {
	var (
		_node = &CattleMove{config: cmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cattlemove.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cattlemove.FieldID,
			},
		}
	)
	if value, ok := cmc.mutation.CattleId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldCattleId,
		})
		_node.CattleId = value
	}
	if value, ok := cmc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldTenantId,
		})
		_node.TenantId = value
	}
	if value, ok := cmc.mutation.TenantName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldTenantName,
		})
		_node.TenantName = value
	}
	if value, ok := cmc.mutation.FarmId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldFarmId,
		})
		_node.FarmId = value
	}
	if value, ok := cmc.mutation.FarmName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldFarmName,
		})
		_node.FarmName = value
	}
	if value, ok := cmc.mutation.ShedId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldShedId,
		})
		_node.ShedId = value
	}
	if value, ok := cmc.mutation.ShedName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldShedName,
		})
		_node.ShedName = value
	}
	if value, ok := cmc.mutation.EarNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldEarNumber,
		})
		_node.EarNumber = value
	}
	if value, ok := cmc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := cmc.mutation.FromShedId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldFromShedId,
		})
		_node.FromShedId = value
	}
	if value, ok := cmc.mutation.FromShed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldFromShed,
		})
		_node.FromShed = value
	}
	if value, ok := cmc.mutation.ToShedId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldToShedId,
		})
		_node.ToShedId = value
	}
	if value, ok := cmc.mutation.ToShed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldToShed,
		})
		_node.ToShed = value
	}
	if value, ok := cmc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := cmc.mutation.ReasonId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldReasonId,
		})
		_node.ReasonId = value
	}
	if value, ok := cmc.mutation.ReasonName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldReasonName,
		})
		_node.ReasonName = value
	}
	if value, ok := cmc.mutation.Remarks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cattlemove.FieldRemarks,
		})
		_node.Remarks = value
	}
	if value, ok := cmc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cmc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cattlemove.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cmc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cattlemove.FieldDeleted,
		})
		_node.Deleted = value
	}
	return _node, _spec
}

// CattleMoveCreateBulk is the builder for creating many CattleMove entities in bulk.
type CattleMoveCreateBulk struct {
	config
	builders []*CattleMoveCreate
}

// Save creates the CattleMove entities in the database.
func (cmcb *CattleMoveCreateBulk) Save(ctx context.Context) ([]*CattleMove, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cmcb.builders))
	nodes := make([]*CattleMove, len(cmcb.builders))
	mutators := make([]Mutator, len(cmcb.builders))
	for i := range cmcb.builders {
		func(i int, root context.Context) {
			builder := cmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CattleMoveMutation)
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
					_, err = mutators[i+1].Mutate(root, cmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cmcb *CattleMoveCreateBulk) SaveX(ctx context.Context) []*CattleMove {
	v, err := cmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cmcb *CattleMoveCreateBulk) Exec(ctx context.Context) error {
	_, err := cmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmcb *CattleMoveCreateBulk) ExecX(ctx context.Context) {
	if err := cmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
