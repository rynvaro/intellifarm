// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/diseaseinfo"
	"cattleai/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DiseaseInfoUpdate is the builder for updating DiseaseInfo entities.
type DiseaseInfoUpdate struct {
	config
	hooks    []Hook
	mutation *DiseaseInfoMutation
}

// Where appends a list predicates to the DiseaseInfoUpdate builder.
func (diu *DiseaseInfoUpdate) Where(ps ...predicate.DiseaseInfo) *DiseaseInfoUpdate {
	diu.mutation.Where(ps...)
	return diu
}

// SetName sets the "name" field.
func (diu *DiseaseInfoUpdate) SetName(s string) *DiseaseInfoUpdate {
	diu.mutation.SetName(s)
	return diu
}

// SetCode sets the "code" field.
func (diu *DiseaseInfoUpdate) SetCode(s string) *DiseaseInfoUpdate {
	diu.mutation.SetCode(s)
	return diu
}

// SetType sets the "type" field.
func (diu *DiseaseInfoUpdate) SetType(s string) *DiseaseInfoUpdate {
	diu.mutation.SetType(s)
	return diu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (diu *DiseaseInfoUpdate) SetNillableType(s *string) *DiseaseInfoUpdate {
	if s != nil {
		diu.SetType(*s)
	}
	return diu
}

// ClearType clears the value of the "type" field.
func (diu *DiseaseInfoUpdate) ClearType() *DiseaseInfoUpdate {
	diu.mutation.ClearType()
	return diu
}

// SetDescription sets the "description" field.
func (diu *DiseaseInfoUpdate) SetDescription(s string) *DiseaseInfoUpdate {
	diu.mutation.SetDescription(s)
	return diu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (diu *DiseaseInfoUpdate) SetNillableDescription(s *string) *DiseaseInfoUpdate {
	if s != nil {
		diu.SetDescription(*s)
	}
	return diu
}

// ClearDescription clears the value of the "description" field.
func (diu *DiseaseInfoUpdate) ClearDescription() *DiseaseInfoUpdate {
	diu.mutation.ClearDescription()
	return diu
}

// SetTenantId sets the "tenantId" field.
func (diu *DiseaseInfoUpdate) SetTenantId(i int64) *DiseaseInfoUpdate {
	diu.mutation.ResetTenantId()
	diu.mutation.SetTenantId(i)
	return diu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (diu *DiseaseInfoUpdate) SetNillableTenantId(i *int64) *DiseaseInfoUpdate {
	if i != nil {
		diu.SetTenantId(*i)
	}
	return diu
}

// AddTenantId adds i to the "tenantId" field.
func (diu *DiseaseInfoUpdate) AddTenantId(i int64) *DiseaseInfoUpdate {
	diu.mutation.AddTenantId(i)
	return diu
}

// ClearTenantId clears the value of the "tenantId" field.
func (diu *DiseaseInfoUpdate) ClearTenantId() *DiseaseInfoUpdate {
	diu.mutation.ClearTenantId()
	return diu
}

// SetTenantName sets the "tenantName" field.
func (diu *DiseaseInfoUpdate) SetTenantName(s string) *DiseaseInfoUpdate {
	diu.mutation.SetTenantName(s)
	return diu
}

// SetNillableTenantName sets the "tenantName" field if the given value is not nil.
func (diu *DiseaseInfoUpdate) SetNillableTenantName(s *string) *DiseaseInfoUpdate {
	if s != nil {
		diu.SetTenantName(*s)
	}
	return diu
}

// ClearTenantName clears the value of the "tenantName" field.
func (diu *DiseaseInfoUpdate) ClearTenantName() *DiseaseInfoUpdate {
	diu.mutation.ClearTenantName()
	return diu
}

// SetRemarks sets the "remarks" field.
func (diu *DiseaseInfoUpdate) SetRemarks(s string) *DiseaseInfoUpdate {
	diu.mutation.SetRemarks(s)
	return diu
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (diu *DiseaseInfoUpdate) SetNillableRemarks(s *string) *DiseaseInfoUpdate {
	if s != nil {
		diu.SetRemarks(*s)
	}
	return diu
}

// ClearRemarks clears the value of the "remarks" field.
func (diu *DiseaseInfoUpdate) ClearRemarks() *DiseaseInfoUpdate {
	diu.mutation.ClearRemarks()
	return diu
}

// SetCreatedAt sets the "createdAt" field.
func (diu *DiseaseInfoUpdate) SetCreatedAt(i int64) *DiseaseInfoUpdate {
	diu.mutation.ResetCreatedAt()
	diu.mutation.SetCreatedAt(i)
	return diu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (diu *DiseaseInfoUpdate) SetNillableCreatedAt(i *int64) *DiseaseInfoUpdate {
	if i != nil {
		diu.SetCreatedAt(*i)
	}
	return diu
}

// AddCreatedAt adds i to the "createdAt" field.
func (diu *DiseaseInfoUpdate) AddCreatedAt(i int64) *DiseaseInfoUpdate {
	diu.mutation.AddCreatedAt(i)
	return diu
}

// ClearCreatedAt clears the value of the "createdAt" field.
func (diu *DiseaseInfoUpdate) ClearCreatedAt() *DiseaseInfoUpdate {
	diu.mutation.ClearCreatedAt()
	return diu
}

// SetUpdatedAt sets the "updatedAt" field.
func (diu *DiseaseInfoUpdate) SetUpdatedAt(i int64) *DiseaseInfoUpdate {
	diu.mutation.ResetUpdatedAt()
	diu.mutation.SetUpdatedAt(i)
	return diu
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (diu *DiseaseInfoUpdate) SetNillableUpdatedAt(i *int64) *DiseaseInfoUpdate {
	if i != nil {
		diu.SetUpdatedAt(*i)
	}
	return diu
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (diu *DiseaseInfoUpdate) AddUpdatedAt(i int64) *DiseaseInfoUpdate {
	diu.mutation.AddUpdatedAt(i)
	return diu
}

// ClearUpdatedAt clears the value of the "updatedAt" field.
func (diu *DiseaseInfoUpdate) ClearUpdatedAt() *DiseaseInfoUpdate {
	diu.mutation.ClearUpdatedAt()
	return diu
}

// SetDeleted sets the "deleted" field.
func (diu *DiseaseInfoUpdate) SetDeleted(i int) *DiseaseInfoUpdate {
	diu.mutation.ResetDeleted()
	diu.mutation.SetDeleted(i)
	return diu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (diu *DiseaseInfoUpdate) SetNillableDeleted(i *int) *DiseaseInfoUpdate {
	if i != nil {
		diu.SetDeleted(*i)
	}
	return diu
}

// AddDeleted adds i to the "deleted" field.
func (diu *DiseaseInfoUpdate) AddDeleted(i int) *DiseaseInfoUpdate {
	diu.mutation.AddDeleted(i)
	return diu
}

// ClearDeleted clears the value of the "deleted" field.
func (diu *DiseaseInfoUpdate) ClearDeleted() *DiseaseInfoUpdate {
	diu.mutation.ClearDeleted()
	return diu
}

// Mutation returns the DiseaseInfoMutation object of the builder.
func (diu *DiseaseInfoUpdate) Mutation() *DiseaseInfoMutation {
	return diu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (diu *DiseaseInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(diu.hooks) == 0 {
		affected, err = diu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			diu.mutation = mutation
			affected, err = diu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(diu.hooks) - 1; i >= 0; i-- {
			if diu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = diu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, diu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (diu *DiseaseInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := diu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (diu *DiseaseInfoUpdate) Exec(ctx context.Context) error {
	_, err := diu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diu *DiseaseInfoUpdate) ExecX(ctx context.Context) {
	if err := diu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (diu *DiseaseInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   diseaseinfo.Table,
			Columns: diseaseinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diseaseinfo.FieldID,
			},
		},
	}
	if ps := diu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldName,
		})
	}
	if value, ok := diu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldCode,
		})
	}
	if value, ok := diu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldType,
		})
	}
	if diu.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: diseaseinfo.FieldType,
		})
	}
	if value, ok := diu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldDescription,
		})
	}
	if diu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: diseaseinfo.FieldDescription,
		})
	}
	if value, ok := diu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldTenantId,
		})
	}
	if value, ok := diu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldTenantId,
		})
	}
	if diu.mutation.TenantIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: diseaseinfo.FieldTenantId,
		})
	}
	if value, ok := diu.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldTenantName,
		})
	}
	if diu.mutation.TenantNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: diseaseinfo.FieldTenantName,
		})
	}
	if value, ok := diu.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldRemarks,
		})
	}
	if diu.mutation.RemarksCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: diseaseinfo.FieldRemarks,
		})
	}
	if value, ok := diu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldCreatedAt,
		})
	}
	if value, ok := diu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldCreatedAt,
		})
	}
	if diu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: diseaseinfo.FieldCreatedAt,
		})
	}
	if value, ok := diu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldUpdatedAt,
		})
	}
	if value, ok := diu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldUpdatedAt,
		})
	}
	if diu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: diseaseinfo.FieldUpdatedAt,
		})
	}
	if value, ok := diu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diseaseinfo.FieldDeleted,
		})
	}
	if value, ok := diu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diseaseinfo.FieldDeleted,
		})
	}
	if diu.mutation.DeletedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: diseaseinfo.FieldDeleted,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, diu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{diseaseinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DiseaseInfoUpdateOne is the builder for updating a single DiseaseInfo entity.
type DiseaseInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DiseaseInfoMutation
}

// SetName sets the "name" field.
func (diuo *DiseaseInfoUpdateOne) SetName(s string) *DiseaseInfoUpdateOne {
	diuo.mutation.SetName(s)
	return diuo
}

// SetCode sets the "code" field.
func (diuo *DiseaseInfoUpdateOne) SetCode(s string) *DiseaseInfoUpdateOne {
	diuo.mutation.SetCode(s)
	return diuo
}

// SetType sets the "type" field.
func (diuo *DiseaseInfoUpdateOne) SetType(s string) *DiseaseInfoUpdateOne {
	diuo.mutation.SetType(s)
	return diuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (diuo *DiseaseInfoUpdateOne) SetNillableType(s *string) *DiseaseInfoUpdateOne {
	if s != nil {
		diuo.SetType(*s)
	}
	return diuo
}

// ClearType clears the value of the "type" field.
func (diuo *DiseaseInfoUpdateOne) ClearType() *DiseaseInfoUpdateOne {
	diuo.mutation.ClearType()
	return diuo
}

// SetDescription sets the "description" field.
func (diuo *DiseaseInfoUpdateOne) SetDescription(s string) *DiseaseInfoUpdateOne {
	diuo.mutation.SetDescription(s)
	return diuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (diuo *DiseaseInfoUpdateOne) SetNillableDescription(s *string) *DiseaseInfoUpdateOne {
	if s != nil {
		diuo.SetDescription(*s)
	}
	return diuo
}

// ClearDescription clears the value of the "description" field.
func (diuo *DiseaseInfoUpdateOne) ClearDescription() *DiseaseInfoUpdateOne {
	diuo.mutation.ClearDescription()
	return diuo
}

// SetTenantId sets the "tenantId" field.
func (diuo *DiseaseInfoUpdateOne) SetTenantId(i int64) *DiseaseInfoUpdateOne {
	diuo.mutation.ResetTenantId()
	diuo.mutation.SetTenantId(i)
	return diuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (diuo *DiseaseInfoUpdateOne) SetNillableTenantId(i *int64) *DiseaseInfoUpdateOne {
	if i != nil {
		diuo.SetTenantId(*i)
	}
	return diuo
}

// AddTenantId adds i to the "tenantId" field.
func (diuo *DiseaseInfoUpdateOne) AddTenantId(i int64) *DiseaseInfoUpdateOne {
	diuo.mutation.AddTenantId(i)
	return diuo
}

// ClearTenantId clears the value of the "tenantId" field.
func (diuo *DiseaseInfoUpdateOne) ClearTenantId() *DiseaseInfoUpdateOne {
	diuo.mutation.ClearTenantId()
	return diuo
}

// SetTenantName sets the "tenantName" field.
func (diuo *DiseaseInfoUpdateOne) SetTenantName(s string) *DiseaseInfoUpdateOne {
	diuo.mutation.SetTenantName(s)
	return diuo
}

// SetNillableTenantName sets the "tenantName" field if the given value is not nil.
func (diuo *DiseaseInfoUpdateOne) SetNillableTenantName(s *string) *DiseaseInfoUpdateOne {
	if s != nil {
		diuo.SetTenantName(*s)
	}
	return diuo
}

// ClearTenantName clears the value of the "tenantName" field.
func (diuo *DiseaseInfoUpdateOne) ClearTenantName() *DiseaseInfoUpdateOne {
	diuo.mutation.ClearTenantName()
	return diuo
}

// SetRemarks sets the "remarks" field.
func (diuo *DiseaseInfoUpdateOne) SetRemarks(s string) *DiseaseInfoUpdateOne {
	diuo.mutation.SetRemarks(s)
	return diuo
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (diuo *DiseaseInfoUpdateOne) SetNillableRemarks(s *string) *DiseaseInfoUpdateOne {
	if s != nil {
		diuo.SetRemarks(*s)
	}
	return diuo
}

// ClearRemarks clears the value of the "remarks" field.
func (diuo *DiseaseInfoUpdateOne) ClearRemarks() *DiseaseInfoUpdateOne {
	diuo.mutation.ClearRemarks()
	return diuo
}

// SetCreatedAt sets the "createdAt" field.
func (diuo *DiseaseInfoUpdateOne) SetCreatedAt(i int64) *DiseaseInfoUpdateOne {
	diuo.mutation.ResetCreatedAt()
	diuo.mutation.SetCreatedAt(i)
	return diuo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (diuo *DiseaseInfoUpdateOne) SetNillableCreatedAt(i *int64) *DiseaseInfoUpdateOne {
	if i != nil {
		diuo.SetCreatedAt(*i)
	}
	return diuo
}

// AddCreatedAt adds i to the "createdAt" field.
func (diuo *DiseaseInfoUpdateOne) AddCreatedAt(i int64) *DiseaseInfoUpdateOne {
	diuo.mutation.AddCreatedAt(i)
	return diuo
}

// ClearCreatedAt clears the value of the "createdAt" field.
func (diuo *DiseaseInfoUpdateOne) ClearCreatedAt() *DiseaseInfoUpdateOne {
	diuo.mutation.ClearCreatedAt()
	return diuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (diuo *DiseaseInfoUpdateOne) SetUpdatedAt(i int64) *DiseaseInfoUpdateOne {
	diuo.mutation.ResetUpdatedAt()
	diuo.mutation.SetUpdatedAt(i)
	return diuo
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (diuo *DiseaseInfoUpdateOne) SetNillableUpdatedAt(i *int64) *DiseaseInfoUpdateOne {
	if i != nil {
		diuo.SetUpdatedAt(*i)
	}
	return diuo
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (diuo *DiseaseInfoUpdateOne) AddUpdatedAt(i int64) *DiseaseInfoUpdateOne {
	diuo.mutation.AddUpdatedAt(i)
	return diuo
}

// ClearUpdatedAt clears the value of the "updatedAt" field.
func (diuo *DiseaseInfoUpdateOne) ClearUpdatedAt() *DiseaseInfoUpdateOne {
	diuo.mutation.ClearUpdatedAt()
	return diuo
}

// SetDeleted sets the "deleted" field.
func (diuo *DiseaseInfoUpdateOne) SetDeleted(i int) *DiseaseInfoUpdateOne {
	diuo.mutation.ResetDeleted()
	diuo.mutation.SetDeleted(i)
	return diuo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (diuo *DiseaseInfoUpdateOne) SetNillableDeleted(i *int) *DiseaseInfoUpdateOne {
	if i != nil {
		diuo.SetDeleted(*i)
	}
	return diuo
}

// AddDeleted adds i to the "deleted" field.
func (diuo *DiseaseInfoUpdateOne) AddDeleted(i int) *DiseaseInfoUpdateOne {
	diuo.mutation.AddDeleted(i)
	return diuo
}

// ClearDeleted clears the value of the "deleted" field.
func (diuo *DiseaseInfoUpdateOne) ClearDeleted() *DiseaseInfoUpdateOne {
	diuo.mutation.ClearDeleted()
	return diuo
}

// Mutation returns the DiseaseInfoMutation object of the builder.
func (diuo *DiseaseInfoUpdateOne) Mutation() *DiseaseInfoMutation {
	return diuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (diuo *DiseaseInfoUpdateOne) Select(field string, fields ...string) *DiseaseInfoUpdateOne {
	diuo.fields = append([]string{field}, fields...)
	return diuo
}

// Save executes the query and returns the updated DiseaseInfo entity.
func (diuo *DiseaseInfoUpdateOne) Save(ctx context.Context) (*DiseaseInfo, error) {
	var (
		err  error
		node *DiseaseInfo
	)
	if len(diuo.hooks) == 0 {
		node, err = diuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			diuo.mutation = mutation
			node, err = diuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(diuo.hooks) - 1; i >= 0; i-- {
			if diuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = diuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, diuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DiseaseInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DiseaseInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (diuo *DiseaseInfoUpdateOne) SaveX(ctx context.Context) *DiseaseInfo {
	node, err := diuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (diuo *DiseaseInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := diuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diuo *DiseaseInfoUpdateOne) ExecX(ctx context.Context) {
	if err := diuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (diuo *DiseaseInfoUpdateOne) sqlSave(ctx context.Context) (_node *DiseaseInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   diseaseinfo.Table,
			Columns: diseaseinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diseaseinfo.FieldID,
			},
		},
	}
	id, ok := diuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DiseaseInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := diuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, diseaseinfo.FieldID)
		for _, f := range fields {
			if !diseaseinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != diseaseinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := diuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldName,
		})
	}
	if value, ok := diuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldCode,
		})
	}
	if value, ok := diuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldType,
		})
	}
	if diuo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: diseaseinfo.FieldType,
		})
	}
	if value, ok := diuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldDescription,
		})
	}
	if diuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: diseaseinfo.FieldDescription,
		})
	}
	if value, ok := diuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldTenantId,
		})
	}
	if value, ok := diuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldTenantId,
		})
	}
	if diuo.mutation.TenantIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: diseaseinfo.FieldTenantId,
		})
	}
	if value, ok := diuo.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldTenantName,
		})
	}
	if diuo.mutation.TenantNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: diseaseinfo.FieldTenantName,
		})
	}
	if value, ok := diuo.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseaseinfo.FieldRemarks,
		})
	}
	if diuo.mutation.RemarksCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: diseaseinfo.FieldRemarks,
		})
	}
	if value, ok := diuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldCreatedAt,
		})
	}
	if value, ok := diuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldCreatedAt,
		})
	}
	if diuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: diseaseinfo.FieldCreatedAt,
		})
	}
	if value, ok := diuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldUpdatedAt,
		})
	}
	if value, ok := diuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: diseaseinfo.FieldUpdatedAt,
		})
	}
	if diuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: diseaseinfo.FieldUpdatedAt,
		})
	}
	if value, ok := diuo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diseaseinfo.FieldDeleted,
		})
	}
	if value, ok := diuo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diseaseinfo.FieldDeleted,
		})
	}
	if diuo.mutation.DeletedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: diseaseinfo.FieldDeleted,
		})
	}
	_node = &DiseaseInfo{config: diuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, diuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{diseaseinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
