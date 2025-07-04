// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/estrus"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EstrusCreate is the builder for creating a Estrus entity.
type EstrusCreate struct {
	config
	mutation *EstrusMutation
	hooks    []Hook
}

// SetCattleId sets the "cattleId" field.
func (ec *EstrusCreate) SetCattleId(i int64) *EstrusCreate {
	ec.mutation.SetCattleId(i)
	return ec
}

// SetTenantId sets the "tenantId" field.
func (ec *EstrusCreate) SetTenantId(i int64) *EstrusCreate {
	ec.mutation.SetTenantId(i)
	return ec
}

// SetTenantName sets the "tenantName" field.
func (ec *EstrusCreate) SetTenantName(s string) *EstrusCreate {
	ec.mutation.SetTenantName(s)
	return ec
}

// SetFarmId sets the "farmId" field.
func (ec *EstrusCreate) SetFarmId(i int64) *EstrusCreate {
	ec.mutation.SetFarmId(i)
	return ec
}

// SetFarmName sets the "farmName" field.
func (ec *EstrusCreate) SetFarmName(s string) *EstrusCreate {
	ec.mutation.SetFarmName(s)
	return ec
}

// SetShedId sets the "shedId" field.
func (ec *EstrusCreate) SetShedId(i int64) *EstrusCreate {
	ec.mutation.SetShedId(i)
	return ec
}

// SetShedName sets the "shedName" field.
func (ec *EstrusCreate) SetShedName(s string) *EstrusCreate {
	ec.mutation.SetShedName(s)
	return ec
}

// SetName sets the "name" field.
func (ec *EstrusCreate) SetName(s string) *EstrusCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ec *EstrusCreate) SetNillableName(s *string) *EstrusCreate {
	if s != nil {
		ec.SetName(*s)
	}
	return ec
}

// SetEarNumber sets the "earNumber" field.
func (ec *EstrusCreate) SetEarNumber(s string) *EstrusCreate {
	ec.mutation.SetEarNumber(s)
	return ec
}

// SetTimes sets the "times" field.
func (ec *EstrusCreate) SetTimes(i int) *EstrusCreate {
	ec.mutation.SetTimes(i)
	return ec
}

// SetEstrusAt sets the "estrusAt" field.
func (ec *EstrusCreate) SetEstrusAt(i int64) *EstrusCreate {
	ec.mutation.SetEstrusAt(i)
	return ec
}

// SetEstrusTypeId sets the "estrusTypeId" field.
func (ec *EstrusCreate) SetEstrusTypeId(i int) *EstrusCreate {
	ec.mutation.SetEstrusTypeId(i)
	return ec
}

// SetEstrusTypeName sets the "estrusTypeName" field.
func (ec *EstrusCreate) SetEstrusTypeName(s string) *EstrusCreate {
	ec.mutation.SetEstrusTypeName(s)
	return ec
}

// SetUserName sets the "userName" field.
func (ec *EstrusCreate) SetUserName(s string) *EstrusCreate {
	ec.mutation.SetUserName(s)
	return ec
}

// SetRemarks sets the "remarks" field.
func (ec *EstrusCreate) SetRemarks(s string) *EstrusCreate {
	ec.mutation.SetRemarks(s)
	return ec
}

// SetCreatedAt sets the "createdAt" field.
func (ec *EstrusCreate) SetCreatedAt(i int64) *EstrusCreate {
	ec.mutation.SetCreatedAt(i)
	return ec
}

// SetUpdatedAt sets the "updatedAt" field.
func (ec *EstrusCreate) SetUpdatedAt(i int64) *EstrusCreate {
	ec.mutation.SetUpdatedAt(i)
	return ec
}

// SetDeleted sets the "deleted" field.
func (ec *EstrusCreate) SetDeleted(i int) *EstrusCreate {
	ec.mutation.SetDeleted(i)
	return ec
}

// Mutation returns the EstrusMutation object of the builder.
func (ec *EstrusCreate) Mutation() *EstrusMutation {
	return ec.mutation
}

// Save creates the Estrus in the database.
func (ec *EstrusCreate) Save(ctx context.Context) (*Estrus, error) {
	var (
		err  error
		node *Estrus
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EstrusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Estrus)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EstrusMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EstrusCreate) SaveX(ctx context.Context) *Estrus {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EstrusCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EstrusCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EstrusCreate) check() error {
	if _, ok := ec.mutation.CattleId(); !ok {
		return &ValidationError{Name: "cattleId", err: errors.New(`ent: missing required field "Estrus.cattleId"`)}
	}
	if _, ok := ec.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "Estrus.tenantId"`)}
	}
	if _, ok := ec.mutation.TenantName(); !ok {
		return &ValidationError{Name: "tenantName", err: errors.New(`ent: missing required field "Estrus.tenantName"`)}
	}
	if _, ok := ec.mutation.FarmId(); !ok {
		return &ValidationError{Name: "farmId", err: errors.New(`ent: missing required field "Estrus.farmId"`)}
	}
	if _, ok := ec.mutation.FarmName(); !ok {
		return &ValidationError{Name: "farmName", err: errors.New(`ent: missing required field "Estrus.farmName"`)}
	}
	if _, ok := ec.mutation.ShedId(); !ok {
		return &ValidationError{Name: "shedId", err: errors.New(`ent: missing required field "Estrus.shedId"`)}
	}
	if _, ok := ec.mutation.ShedName(); !ok {
		return &ValidationError{Name: "shedName", err: errors.New(`ent: missing required field "Estrus.shedName"`)}
	}
	if _, ok := ec.mutation.EarNumber(); !ok {
		return &ValidationError{Name: "earNumber", err: errors.New(`ent: missing required field "Estrus.earNumber"`)}
	}
	if _, ok := ec.mutation.Times(); !ok {
		return &ValidationError{Name: "times", err: errors.New(`ent: missing required field "Estrus.times"`)}
	}
	if _, ok := ec.mutation.EstrusAt(); !ok {
		return &ValidationError{Name: "estrusAt", err: errors.New(`ent: missing required field "Estrus.estrusAt"`)}
	}
	if _, ok := ec.mutation.EstrusTypeId(); !ok {
		return &ValidationError{Name: "estrusTypeId", err: errors.New(`ent: missing required field "Estrus.estrusTypeId"`)}
	}
	if _, ok := ec.mutation.EstrusTypeName(); !ok {
		return &ValidationError{Name: "estrusTypeName", err: errors.New(`ent: missing required field "Estrus.estrusTypeName"`)}
	}
	if _, ok := ec.mutation.UserName(); !ok {
		return &ValidationError{Name: "userName", err: errors.New(`ent: missing required field "Estrus.userName"`)}
	}
	if _, ok := ec.mutation.Remarks(); !ok {
		return &ValidationError{Name: "remarks", err: errors.New(`ent: missing required field "Estrus.remarks"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Estrus.createdAt"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Estrus.updatedAt"`)}
	}
	if _, ok := ec.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Estrus.deleted"`)}
	}
	return nil
}

func (ec *EstrusCreate) sqlSave(ctx context.Context) (*Estrus, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *EstrusCreate) createSpec() (*Estrus, *sqlgraph.CreateSpec) {
	var (
		_node = &Estrus{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: estrus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: estrus.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.CattleId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: estrus.FieldCattleId,
		})
		_node.CattleId = value
	}
	if value, ok := ec.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: estrus.FieldTenantId,
		})
		_node.TenantId = value
	}
	if value, ok := ec.mutation.TenantName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: estrus.FieldTenantName,
		})
		_node.TenantName = value
	}
	if value, ok := ec.mutation.FarmId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: estrus.FieldFarmId,
		})
		_node.FarmId = value
	}
	if value, ok := ec.mutation.FarmName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: estrus.FieldFarmName,
		})
		_node.FarmName = value
	}
	if value, ok := ec.mutation.ShedId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: estrus.FieldShedId,
		})
		_node.ShedId = value
	}
	if value, ok := ec.mutation.ShedName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: estrus.FieldShedName,
		})
		_node.ShedName = value
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: estrus.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ec.mutation.EarNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: estrus.FieldEarNumber,
		})
		_node.EarNumber = value
	}
	if value, ok := ec.mutation.Times(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: estrus.FieldTimes,
		})
		_node.Times = value
	}
	if value, ok := ec.mutation.EstrusAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: estrus.FieldEstrusAt,
		})
		_node.EstrusAt = value
	}
	if value, ok := ec.mutation.EstrusTypeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: estrus.FieldEstrusTypeId,
		})
		_node.EstrusTypeId = value
	}
	if value, ok := ec.mutation.EstrusTypeName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: estrus.FieldEstrusTypeName,
		})
		_node.EstrusTypeName = value
	}
	if value, ok := ec.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: estrus.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := ec.mutation.Remarks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: estrus.FieldRemarks,
		})
		_node.Remarks = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: estrus.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: estrus.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: estrus.FieldDeleted,
		})
		_node.Deleted = value
	}
	return _node, _spec
}

// EstrusCreateBulk is the builder for creating many Estrus entities in bulk.
type EstrusCreateBulk struct {
	config
	builders []*EstrusCreate
}

// Save creates the Estrus entities in the database.
func (ecb *EstrusCreateBulk) Save(ctx context.Context) ([]*Estrus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Estrus, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EstrusMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EstrusCreateBulk) SaveX(ctx context.Context) []*Estrus {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EstrusCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EstrusCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
