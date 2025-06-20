// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetLevel sets the "level" field.
func (uc *UserCreate) SetLevel(i int) *UserCreate {
	uc.mutation.SetLevel(i)
	return uc
}

// SetFarmId sets the "farmId" field.
func (uc *UserCreate) SetFarmId(i int64) *UserCreate {
	uc.mutation.SetFarmId(i)
	return uc
}

// SetNillableFarmId sets the "farmId" field if the given value is not nil.
func (uc *UserCreate) SetNillableFarmId(i *int64) *UserCreate {
	if i != nil {
		uc.SetFarmId(*i)
	}
	return uc
}

// SetFarmName sets the "farmName" field.
func (uc *UserCreate) SetFarmName(s string) *UserCreate {
	uc.mutation.SetFarmName(s)
	return uc
}

// SetNillableFarmName sets the "farmName" field if the given value is not nil.
func (uc *UserCreate) SetNillableFarmName(s *string) *UserCreate {
	if s != nil {
		uc.SetFarmName(*s)
	}
	return uc
}

// SetPositionId sets the "positionId" field.
func (uc *UserCreate) SetPositionId(i int64) *UserCreate {
	uc.mutation.SetPositionId(i)
	return uc
}

// SetNillablePositionId sets the "positionId" field if the given value is not nil.
func (uc *UserCreate) SetNillablePositionId(i *int64) *UserCreate {
	if i != nil {
		uc.SetPositionId(*i)
	}
	return uc
}

// SetPositionName sets the "positionName" field.
func (uc *UserCreate) SetPositionName(s string) *UserCreate {
	uc.mutation.SetPositionName(s)
	return uc
}

// SetNillablePositionName sets the "positionName" field if the given value is not nil.
func (uc *UserCreate) SetNillablePositionName(s *string) *UserCreate {
	if s != nil {
		uc.SetPositionName(*s)
	}
	return uc
}

// SetDutyName sets the "dutyName" field.
func (uc *UserCreate) SetDutyName(s string) *UserCreate {
	uc.mutation.SetDutyName(s)
	return uc
}

// SetNillableDutyName sets the "dutyName" field if the given value is not nil.
func (uc *UserCreate) SetNillableDutyName(s *string) *UserCreate {
	if s != nil {
		uc.SetDutyName(*s)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetGender sets the "gender" field.
func (uc *UserCreate) SetGender(i int) *UserCreate {
	uc.mutation.SetGender(i)
	return uc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uc *UserCreate) SetNillableGender(i *int) *UserCreate {
	if i != nil {
		uc.SetGender(*i)
	}
	return uc
}

// SetAge sets the "age" field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.mutation.SetAge(i)
	return uc
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uc *UserCreate) SetNillableAge(i *int) *UserCreate {
	if i != nil {
		uc.SetAge(*i)
	}
	return uc
}

// SetEducation sets the "education" field.
func (uc *UserCreate) SetEducation(s string) *UserCreate {
	uc.mutation.SetEducation(s)
	return uc
}

// SetNillableEducation sets the "education" field if the given value is not nil.
func (uc *UserCreate) SetNillableEducation(s *string) *UserCreate {
	if s != nil {
		uc.SetEducation(*s)
	}
	return uc
}

// SetMajor sets the "major" field.
func (uc *UserCreate) SetMajor(s string) *UserCreate {
	uc.mutation.SetMajor(s)
	return uc
}

// SetNillableMajor sets the "major" field if the given value is not nil.
func (uc *UserCreate) SetNillableMajor(s *string) *UserCreate {
	if s != nil {
		uc.SetMajor(*s)
	}
	return uc
}

// SetJobTitle sets the "jobTitle" field.
func (uc *UserCreate) SetJobTitle(s string) *UserCreate {
	uc.mutation.SetJobTitle(s)
	return uc
}

// SetNillableJobTitle sets the "jobTitle" field if the given value is not nil.
func (uc *UserCreate) SetNillableJobTitle(s *string) *UserCreate {
	if s != nil {
		uc.SetJobTitle(*s)
	}
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetIdCard sets the "idCard" field.
func (uc *UserCreate) SetIdCard(s string) *UserCreate {
	uc.mutation.SetIdCard(s)
	return uc
}

// SetNillableIdCard sets the "idCard" field if the given value is not nil.
func (uc *UserCreate) SetNillableIdCard(s *string) *UserCreate {
	if s != nil {
		uc.SetIdCard(*s)
	}
	return uc
}

// SetAddress sets the "address" field.
func (uc *UserCreate) SetAddress(s string) *UserCreate {
	uc.mutation.SetAddress(s)
	return uc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uc *UserCreate) SetNillableAddress(s *string) *UserCreate {
	if s != nil {
		uc.SetAddress(*s)
	}
	return uc
}

// SetOnJobState sets the "onJobState" field.
func (uc *UserCreate) SetOnJobState(i int) *UserCreate {
	uc.mutation.SetOnJobState(i)
	return uc
}

// SetNillableOnJobState sets the "onJobState" field if the given value is not nil.
func (uc *UserCreate) SetNillableOnJobState(i *int) *UserCreate {
	if i != nil {
		uc.SetOnJobState(*i)
	}
	return uc
}

// SetJoinedAt sets the "joinedAt" field.
func (uc *UserCreate) SetJoinedAt(i int64) *UserCreate {
	uc.mutation.SetJoinedAt(i)
	return uc
}

// SetNillableJoinedAt sets the "joinedAt" field if the given value is not nil.
func (uc *UserCreate) SetNillableJoinedAt(i *int64) *UserCreate {
	if i != nil {
		uc.SetJoinedAt(*i)
	}
	return uc
}

// SetTenantId sets the "tenantId" field.
func (uc *UserCreate) SetTenantId(i int64) *UserCreate {
	uc.mutation.SetTenantId(i)
	return uc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenantId(i *int64) *UserCreate {
	if i != nil {
		uc.SetTenantId(*i)
	}
	return uc
}

// SetTenantName sets the "tenantName" field.
func (uc *UserCreate) SetTenantName(s string) *UserCreate {
	uc.mutation.SetTenantName(s)
	return uc
}

// SetNillableTenantName sets the "tenantName" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenantName(s *string) *UserCreate {
	if s != nil {
		uc.SetTenantName(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetRemarks sets the "remarks" field.
func (uc *UserCreate) SetRemarks(s string) *UserCreate {
	uc.mutation.SetRemarks(s)
	return uc
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (uc *UserCreate) SetNillableRemarks(s *string) *UserCreate {
	if s != nil {
		uc.SetRemarks(*s)
	}
	return uc
}

// SetCreatedAt sets the "createdAt" field.
func (uc *UserCreate) SetCreatedAt(i int64) *UserCreate {
	uc.mutation.SetCreatedAt(i)
	return uc
}

// SetUpdatedAt sets the "updatedAt" field.
func (uc *UserCreate) SetUpdatedAt(i int64) *UserCreate {
	uc.mutation.SetUpdatedAt(i)
	return uc
}

// SetDeleted sets the "deleted" field.
func (uc *UserCreate) SetDeleted(i int) *UserCreate {
	uc.mutation.SetDeleted(i)
	return uc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeleted(i *int) *UserCreate {
	if i != nil {
		uc.SetDeleted(*i)
	}
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.OnJobState(); !ok {
		v := user.DefaultOnJobState
		uc.mutation.SetOnJobState(v)
	}
	if _, ok := uc.mutation.Deleted(); !ok {
		v := user.DefaultDeleted
		uc.mutation.SetDeleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "User.level"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.OnJobState(); !ok {
		return &ValidationError{Name: "onJobState", err: errors.New(`ent: missing required field "User.onJobState"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "User.createdAt"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "User.updatedAt"`)}
	}
	if _, ok := uc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "User.deleted"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldLevel,
		})
		_node.Level = value
	}
	if value, ok := uc.mutation.FarmId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldFarmId,
		})
		_node.FarmId = value
	}
	if value, ok := uc.mutation.FarmName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFarmName,
		})
		_node.FarmName = value
	}
	if value, ok := uc.mutation.PositionId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldPositionId,
		})
		_node.PositionId = value
	}
	if value, ok := uc.mutation.PositionName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPositionName,
		})
		_node.PositionName = value
	}
	if value, ok := uc.mutation.DutyName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDutyName,
		})
		_node.DutyName = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		_node.Name = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := uc.mutation.Education(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEducation,
		})
		_node.Education = value
	}
	if value, ok := uc.mutation.Major(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldMajor,
		})
		_node.Major = value
	}
	if value, ok := uc.mutation.JobTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldJobTitle,
		})
		_node.JobTitle = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := uc.mutation.IdCard(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldIdCard,
		})
		_node.IdCard = value
	}
	if value, ok := uc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := uc.mutation.OnJobState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldOnJobState,
		})
		_node.OnJobState = value
	}
	if value, ok := uc.mutation.JoinedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldJoinedAt,
		})
		_node.JoinedAt = value
	}
	if value, ok := uc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTenantId,
		})
		_node.TenantId = value
	}
	if value, ok := uc.mutation.TenantName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTenantName,
		})
		_node.TenantName = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.Remarks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRemarks,
		})
		_node.Remarks = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldDeleted,
		})
		_node.Deleted = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
