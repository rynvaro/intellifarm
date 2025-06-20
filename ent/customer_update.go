// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/customer"
	"cattleai/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CustomerUpdate) SetName(s string) *CustomerUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetType sets the "type" field.
func (cu *CustomerUpdate) SetType(s string) *CustomerUpdate {
	cu.mutation.SetType(s)
	return cu
}

// SetGender sets the "gender" field.
func (cu *CustomerUpdate) SetGender(s string) *CustomerUpdate {
	cu.mutation.SetGender(s)
	return cu
}

// SetProvince sets the "province" field.
func (cu *CustomerUpdate) SetProvince(s string) *CustomerUpdate {
	cu.mutation.SetProvince(s)
	return cu
}

// SetPhone sets the "phone" field.
func (cu *CustomerUpdate) SetPhone(s string) *CustomerUpdate {
	cu.mutation.SetPhone(s)
	return cu
}

// SetAddress sets the "address" field.
func (cu *CustomerUpdate) SetAddress(s string) *CustomerUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetEmail sets the "email" field.
func (cu *CustomerUpdate) SetEmail(s string) *CustomerUpdate {
	cu.mutation.SetEmail(s)
	return cu
}

// SetBank sets the "bank" field.
func (cu *CustomerUpdate) SetBank(s string) *CustomerUpdate {
	cu.mutation.SetBank(s)
	return cu
}

// SetAccount sets the "account" field.
func (cu *CustomerUpdate) SetAccount(s string) *CustomerUpdate {
	cu.mutation.SetAccount(s)
	return cu
}

// SetTenantId sets the "tenantId" field.
func (cu *CustomerUpdate) SetTenantId(i int64) *CustomerUpdate {
	cu.mutation.ResetTenantId()
	cu.mutation.SetTenantId(i)
	return cu
}

// AddTenantId adds i to the "tenantId" field.
func (cu *CustomerUpdate) AddTenantId(i int64) *CustomerUpdate {
	cu.mutation.AddTenantId(i)
	return cu
}

// SetTenantName sets the "tenantName" field.
func (cu *CustomerUpdate) SetTenantName(s string) *CustomerUpdate {
	cu.mutation.SetTenantName(s)
	return cu
}

// SetRemarks sets the "remarks" field.
func (cu *CustomerUpdate) SetRemarks(s string) *CustomerUpdate {
	cu.mutation.SetRemarks(s)
	return cu
}

// SetCreatedAt sets the "createdAt" field.
func (cu *CustomerUpdate) SetCreatedAt(i int64) *CustomerUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(i)
	return cu
}

// AddCreatedAt adds i to the "createdAt" field.
func (cu *CustomerUpdate) AddCreatedAt(i int64) *CustomerUpdate {
	cu.mutation.AddCreatedAt(i)
	return cu
}

// SetUpdatedAt sets the "updatedAt" field.
func (cu *CustomerUpdate) SetUpdatedAt(i int64) *CustomerUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(i)
	return cu
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (cu *CustomerUpdate) AddUpdatedAt(i int64) *CustomerUpdate {
	cu.mutation.AddUpdatedAt(i)
	return cu
}

// SetDeleted sets the "deleted" field.
func (cu *CustomerUpdate) SetDeleted(i int) *CustomerUpdate {
	cu.mutation.ResetDeleted()
	cu.mutation.SetDeleted(i)
	return cu
}

// AddDeleted adds i to the "deleted" field.
func (cu *CustomerUpdate) AddDeleted(i int) *CustomerUpdate {
	cu.mutation.AddDeleted(i)
	return cu
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldType,
		})
	}
	if value, ok := cu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldGender,
		})
	}
	if value, ok := cu.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldProvince,
		})
	}
	if value, ok := cu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
	}
	if value, ok := cu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldEmail,
		})
	}
	if value, ok := cu.mutation.Bank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldBank,
		})
	}
	if value, ok := cu.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAccount,
		})
	}
	if value, ok := cu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldTenantId,
		})
	}
	if value, ok := cu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldTenantId,
		})
	}
	if value, ok := cu.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldTenantName,
		})
	}
	if value, ok := cu.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldRemarks,
		})
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customer.FieldDeleted,
		})
	}
	if value, ok := cu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customer.FieldDeleted,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomerMutation
}

// SetName sets the "name" field.
func (cuo *CustomerUpdateOne) SetName(s string) *CustomerUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetType sets the "type" field.
func (cuo *CustomerUpdateOne) SetType(s string) *CustomerUpdateOne {
	cuo.mutation.SetType(s)
	return cuo
}

// SetGender sets the "gender" field.
func (cuo *CustomerUpdateOne) SetGender(s string) *CustomerUpdateOne {
	cuo.mutation.SetGender(s)
	return cuo
}

// SetProvince sets the "province" field.
func (cuo *CustomerUpdateOne) SetProvince(s string) *CustomerUpdateOne {
	cuo.mutation.SetProvince(s)
	return cuo
}

// SetPhone sets the "phone" field.
func (cuo *CustomerUpdateOne) SetPhone(s string) *CustomerUpdateOne {
	cuo.mutation.SetPhone(s)
	return cuo
}

// SetAddress sets the "address" field.
func (cuo *CustomerUpdateOne) SetAddress(s string) *CustomerUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetEmail sets the "email" field.
func (cuo *CustomerUpdateOne) SetEmail(s string) *CustomerUpdateOne {
	cuo.mutation.SetEmail(s)
	return cuo
}

// SetBank sets the "bank" field.
func (cuo *CustomerUpdateOne) SetBank(s string) *CustomerUpdateOne {
	cuo.mutation.SetBank(s)
	return cuo
}

// SetAccount sets the "account" field.
func (cuo *CustomerUpdateOne) SetAccount(s string) *CustomerUpdateOne {
	cuo.mutation.SetAccount(s)
	return cuo
}

// SetTenantId sets the "tenantId" field.
func (cuo *CustomerUpdateOne) SetTenantId(i int64) *CustomerUpdateOne {
	cuo.mutation.ResetTenantId()
	cuo.mutation.SetTenantId(i)
	return cuo
}

// AddTenantId adds i to the "tenantId" field.
func (cuo *CustomerUpdateOne) AddTenantId(i int64) *CustomerUpdateOne {
	cuo.mutation.AddTenantId(i)
	return cuo
}

// SetTenantName sets the "tenantName" field.
func (cuo *CustomerUpdateOne) SetTenantName(s string) *CustomerUpdateOne {
	cuo.mutation.SetTenantName(s)
	return cuo
}

// SetRemarks sets the "remarks" field.
func (cuo *CustomerUpdateOne) SetRemarks(s string) *CustomerUpdateOne {
	cuo.mutation.SetRemarks(s)
	return cuo
}

// SetCreatedAt sets the "createdAt" field.
func (cuo *CustomerUpdateOne) SetCreatedAt(i int64) *CustomerUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(i)
	return cuo
}

// AddCreatedAt adds i to the "createdAt" field.
func (cuo *CustomerUpdateOne) AddCreatedAt(i int64) *CustomerUpdateOne {
	cuo.mutation.AddCreatedAt(i)
	return cuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (cuo *CustomerUpdateOne) SetUpdatedAt(i int64) *CustomerUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(i)
	return cuo
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (cuo *CustomerUpdateOne) AddUpdatedAt(i int64) *CustomerUpdateOne {
	cuo.mutation.AddUpdatedAt(i)
	return cuo
}

// SetDeleted sets the "deleted" field.
func (cuo *CustomerUpdateOne) SetDeleted(i int) *CustomerUpdateOne {
	cuo.mutation.ResetDeleted()
	cuo.mutation.SetDeleted(i)
	return cuo
}

// AddDeleted adds i to the "deleted" field.
func (cuo *CustomerUpdateOne) AddDeleted(i int) *CustomerUpdateOne {
	cuo.mutation.AddDeleted(i)
	return cuo
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CustomerUpdateOne) Select(field string, fields ...string) *CustomerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Customer entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Customer)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CustomerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Customer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for _, f := range fields {
			if !customer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldType,
		})
	}
	if value, ok := cuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldGender,
		})
	}
	if value, ok := cuo.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldProvince,
		})
	}
	if value, ok := cuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
	}
	if value, ok := cuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldEmail,
		})
	}
	if value, ok := cuo.mutation.Bank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldBank,
		})
	}
	if value, ok := cuo.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAccount,
		})
	}
	if value, ok := cuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldTenantId,
		})
	}
	if value, ok := cuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldTenantId,
		})
	}
	if value, ok := cuo.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldTenantName,
		})
	}
	if value, ok := cuo.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldRemarks,
		})
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customer.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customer.FieldDeleted,
		})
	}
	if value, ok := cuo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customer.FieldDeleted,
		})
	}
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
