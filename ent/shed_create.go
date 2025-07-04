// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/shed"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShedCreate is the builder for creating a Shed entity.
type ShedCreate struct {
	config
	mutation *ShedMutation
	hooks    []Hook
}

// SetFarmId sets the "farmId" field.
func (sc *ShedCreate) SetFarmId(i int64) *ShedCreate {
	sc.mutation.SetFarmId(i)
	return sc
}

// SetFarmName sets the "farmName" field.
func (sc *ShedCreate) SetFarmName(s string) *ShedCreate {
	sc.mutation.SetFarmName(s)
	return sc
}

// SetName sets the "name" field.
func (sc *ShedCreate) SetName(s string) *ShedCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetShedCateId sets the "shedCateId" field.
func (sc *ShedCreate) SetShedCateId(i int64) *ShedCreate {
	sc.mutation.SetShedCateId(i)
	return sc
}

// SetShedCateName sets the "shedCateName" field.
func (sc *ShedCreate) SetShedCateName(s string) *ShedCreate {
	sc.mutation.SetShedCateName(s)
	return sc
}

// SetShedTypeId sets the "shedTypeId" field.
func (sc *ShedCreate) SetShedTypeId(i int) *ShedCreate {
	sc.mutation.SetShedTypeId(i)
	return sc
}

// SetShedTypeName sets the "shedTypeName" field.
func (sc *ShedCreate) SetShedTypeName(s string) *ShedCreate {
	sc.mutation.SetShedTypeName(s)
	return sc
}

// SetSquare sets the "square" field.
func (sc *ShedCreate) SetSquare(i int64) *ShedCreate {
	sc.mutation.SetSquare(i)
	return sc
}

// SetLength sets the "length" field.
func (sc *ShedCreate) SetLength(i int64) *ShedCreate {
	sc.mutation.SetLength(i)
	return sc
}

// SetWidth sets the "width" field.
func (sc *ShedCreate) SetWidth(i int64) *ShedCreate {
	sc.mutation.SetWidth(i)
	return sc
}

// SetHeight sets the "height" field.
func (sc *ShedCreate) SetHeight(i int64) *ShedCreate {
	sc.mutation.SetHeight(i)
	return sc
}

// SetTenantId sets the "tenantId" field.
func (sc *ShedCreate) SetTenantId(i int64) *ShedCreate {
	sc.mutation.SetTenantId(i)
	return sc
}

// SetTenantName sets the "tenantName" field.
func (sc *ShedCreate) SetTenantName(s string) *ShedCreate {
	sc.mutation.SetTenantName(s)
	return sc
}

// SetRemarks sets the "remarks" field.
func (sc *ShedCreate) SetRemarks(s string) *ShedCreate {
	sc.mutation.SetRemarks(s)
	return sc
}

// SetUserId sets the "userId" field.
func (sc *ShedCreate) SetUserId(i int) *ShedCreate {
	sc.mutation.SetUserId(i)
	return sc
}

// SetUserName sets the "userName" field.
func (sc *ShedCreate) SetUserName(s string) *ShedCreate {
	sc.mutation.SetUserName(s)
	return sc
}

// SetCreatedAt sets the "createdAt" field.
func (sc *ShedCreate) SetCreatedAt(i int64) *ShedCreate {
	sc.mutation.SetCreatedAt(i)
	return sc
}

// SetUpdatedAt sets the "updatedAt" field.
func (sc *ShedCreate) SetUpdatedAt(i int64) *ShedCreate {
	sc.mutation.SetUpdatedAt(i)
	return sc
}

// SetDeleted sets the "deleted" field.
func (sc *ShedCreate) SetDeleted(i int) *ShedCreate {
	sc.mutation.SetDeleted(i)
	return sc
}

// Mutation returns the ShedMutation object of the builder.
func (sc *ShedCreate) Mutation() *ShedMutation {
	return sc.mutation
}

// Save creates the Shed in the database.
func (sc *ShedCreate) Save(ctx context.Context) (*Shed, error) {
	var (
		err  error
		node *Shed
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Shed)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ShedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShedCreate) SaveX(ctx context.Context) *Shed {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ShedCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ShedCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShedCreate) check() error {
	if _, ok := sc.mutation.FarmId(); !ok {
		return &ValidationError{Name: "farmId", err: errors.New(`ent: missing required field "Shed.farmId"`)}
	}
	if _, ok := sc.mutation.FarmName(); !ok {
		return &ValidationError{Name: "farmName", err: errors.New(`ent: missing required field "Shed.farmName"`)}
	}
	if v, ok := sc.mutation.FarmName(); ok {
		if err := shed.FarmNameValidator(v); err != nil {
			return &ValidationError{Name: "farmName", err: fmt.Errorf(`ent: validator failed for field "Shed.farmName": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Shed.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := shed.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Shed.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.ShedCateId(); !ok {
		return &ValidationError{Name: "shedCateId", err: errors.New(`ent: missing required field "Shed.shedCateId"`)}
	}
	if _, ok := sc.mutation.ShedCateName(); !ok {
		return &ValidationError{Name: "shedCateName", err: errors.New(`ent: missing required field "Shed.shedCateName"`)}
	}
	if v, ok := sc.mutation.ShedCateName(); ok {
		if err := shed.ShedCateNameValidator(v); err != nil {
			return &ValidationError{Name: "shedCateName", err: fmt.Errorf(`ent: validator failed for field "Shed.shedCateName": %w`, err)}
		}
	}
	if _, ok := sc.mutation.ShedTypeId(); !ok {
		return &ValidationError{Name: "shedTypeId", err: errors.New(`ent: missing required field "Shed.shedTypeId"`)}
	}
	if _, ok := sc.mutation.ShedTypeName(); !ok {
		return &ValidationError{Name: "shedTypeName", err: errors.New(`ent: missing required field "Shed.shedTypeName"`)}
	}
	if v, ok := sc.mutation.ShedTypeName(); ok {
		if err := shed.ShedTypeNameValidator(v); err != nil {
			return &ValidationError{Name: "shedTypeName", err: fmt.Errorf(`ent: validator failed for field "Shed.shedTypeName": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Square(); !ok {
		return &ValidationError{Name: "square", err: errors.New(`ent: missing required field "Shed.square"`)}
	}
	if _, ok := sc.mutation.Length(); !ok {
		return &ValidationError{Name: "length", err: errors.New(`ent: missing required field "Shed.length"`)}
	}
	if _, ok := sc.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "Shed.width"`)}
	}
	if _, ok := sc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Shed.height"`)}
	}
	if _, ok := sc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "Shed.tenantId"`)}
	}
	if _, ok := sc.mutation.TenantName(); !ok {
		return &ValidationError{Name: "tenantName", err: errors.New(`ent: missing required field "Shed.tenantName"`)}
	}
	if _, ok := sc.mutation.Remarks(); !ok {
		return &ValidationError{Name: "remarks", err: errors.New(`ent: missing required field "Shed.remarks"`)}
	}
	if _, ok := sc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "Shed.userId"`)}
	}
	if _, ok := sc.mutation.UserName(); !ok {
		return &ValidationError{Name: "userName", err: errors.New(`ent: missing required field "Shed.userName"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Shed.createdAt"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Shed.updatedAt"`)}
	}
	if _, ok := sc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Shed.deleted"`)}
	}
	return nil
}

func (sc *ShedCreate) sqlSave(ctx context.Context) (*Shed, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ShedCreate) createSpec() (*Shed, *sqlgraph.CreateSpec) {
	var (
		_node = &Shed{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shed.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shed.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.FarmId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldFarmId,
		})
		_node.FarmId = value
	}
	if value, ok := sc.mutation.FarmName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shed.FieldFarmName,
		})
		_node.FarmName = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shed.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.ShedCateId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldShedCateId,
		})
		_node.ShedCateId = value
	}
	if value, ok := sc.mutation.ShedCateName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shed.FieldShedCateName,
		})
		_node.ShedCateName = value
	}
	if value, ok := sc.mutation.ShedTypeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shed.FieldShedTypeId,
		})
		_node.ShedTypeId = value
	}
	if value, ok := sc.mutation.ShedTypeName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shed.FieldShedTypeName,
		})
		_node.ShedTypeName = value
	}
	if value, ok := sc.mutation.Square(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldSquare,
		})
		_node.Square = value
	}
	if value, ok := sc.mutation.Length(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldLength,
		})
		_node.Length = value
	}
	if value, ok := sc.mutation.Width(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldWidth,
		})
		_node.Width = value
	}
	if value, ok := sc.mutation.Height(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldHeight,
		})
		_node.Height = value
	}
	if value, ok := sc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldTenantId,
		})
		_node.TenantId = value
	}
	if value, ok := sc.mutation.TenantName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shed.FieldTenantName,
		})
		_node.TenantName = value
	}
	if value, ok := sc.mutation.Remarks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shed.FieldRemarks,
		})
		_node.Remarks = value
	}
	if value, ok := sc.mutation.UserId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shed.FieldUserId,
		})
		_node.UserId = value
	}
	if value, ok := sc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shed.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: shed.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shed.FieldDeleted,
		})
		_node.Deleted = value
	}
	return _node, _spec
}

// ShedCreateBulk is the builder for creating many Shed entities in bulk.
type ShedCreateBulk struct {
	config
	builders []*ShedCreate
}

// Save creates the Shed entities in the database.
func (scb *ShedCreateBulk) Save(ctx context.Context) ([]*Shed, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Shed, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShedMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShedCreateBulk) SaveX(ctx context.Context) []*Shed {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ShedCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ShedCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
