// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/inventoryflow"
	"cattleai/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InventoryFlowUpdate is the builder for updating InventoryFlow entities.
type InventoryFlowUpdate struct {
	config
	hooks    []Hook
	mutation *InventoryFlowMutation
}

// Where appends a list predicates to the InventoryFlowUpdate builder.
func (ifu *InventoryFlowUpdate) Where(ps ...predicate.InventoryFlow) *InventoryFlowUpdate {
	ifu.mutation.Where(ps...)
	return ifu
}

// SetSysMaterialId sets the "sysMaterialId" field.
func (ifu *InventoryFlowUpdate) SetSysMaterialId(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetSysMaterialId()
	ifu.mutation.SetSysMaterialId(i)
	return ifu
}

// AddSysMaterialId adds i to the "sysMaterialId" field.
func (ifu *InventoryFlowUpdate) AddSysMaterialId(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddSysMaterialId(i)
	return ifu
}

// SetMaterialId sets the "materialId" field.
func (ifu *InventoryFlowUpdate) SetMaterialId(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetMaterialId()
	ifu.mutation.SetMaterialId(i)
	return ifu
}

// AddMaterialId adds i to the "materialId" field.
func (ifu *InventoryFlowUpdate) AddMaterialId(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddMaterialId(i)
	return ifu
}

// SetMaterialName sets the "materialName" field.
func (ifu *InventoryFlowUpdate) SetMaterialName(s string) *InventoryFlowUpdate {
	ifu.mutation.SetMaterialName(s)
	return ifu
}

// SetMaterialCode sets the "materialCode" field.
func (ifu *InventoryFlowUpdate) SetMaterialCode(s string) *InventoryFlowUpdate {
	ifu.mutation.SetMaterialCode(s)
	return ifu
}

// SetSeqNumber sets the "seqNumber" field.
func (ifu *InventoryFlowUpdate) SetSeqNumber(s string) *InventoryFlowUpdate {
	ifu.mutation.SetSeqNumber(s)
	return ifu
}

// SetDate sets the "date" field.
func (ifu *InventoryFlowUpdate) SetDate(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetDate()
	ifu.mutation.SetDate(i)
	return ifu
}

// AddDate adds i to the "date" field.
func (ifu *InventoryFlowUpdate) AddDate(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddDate(i)
	return ifu
}

// SetType sets the "type" field.
func (ifu *InventoryFlowUpdate) SetType(i int) *InventoryFlowUpdate {
	ifu.mutation.ResetType()
	ifu.mutation.SetType(i)
	return ifu
}

// AddType adds i to the "type" field.
func (ifu *InventoryFlowUpdate) AddType(i int) *InventoryFlowUpdate {
	ifu.mutation.AddType(i)
	return ifu
}

// SetStatus sets the "status" field.
func (ifu *InventoryFlowUpdate) SetStatus(s string) *InventoryFlowUpdate {
	ifu.mutation.SetStatus(s)
	return ifu
}

// SetCount sets the "count" field.
func (ifu *InventoryFlowUpdate) SetCount(i int) *InventoryFlowUpdate {
	ifu.mutation.ResetCount()
	ifu.mutation.SetCount(i)
	return ifu
}

// AddCount adds i to the "count" field.
func (ifu *InventoryFlowUpdate) AddCount(i int) *InventoryFlowUpdate {
	ifu.mutation.AddCount(i)
	return ifu
}

// SetUnit sets the "unit" field.
func (ifu *InventoryFlowUpdate) SetUnit(s string) *InventoryFlowUpdate {
	ifu.mutation.SetUnit(s)
	return ifu
}

// SetBefore sets the "before" field.
func (ifu *InventoryFlowUpdate) SetBefore(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetBefore()
	ifu.mutation.SetBefore(i)
	return ifu
}

// AddBefore adds i to the "before" field.
func (ifu *InventoryFlowUpdate) AddBefore(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddBefore(i)
	return ifu
}

// SetAfter sets the "after" field.
func (ifu *InventoryFlowUpdate) SetAfter(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetAfter()
	ifu.mutation.SetAfter(i)
	return ifu
}

// AddAfter adds i to the "after" field.
func (ifu *InventoryFlowUpdate) AddAfter(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddAfter(i)
	return ifu
}

// SetUserName sets the "userName" field.
func (ifu *InventoryFlowUpdate) SetUserName(s string) *InventoryFlowUpdate {
	ifu.mutation.SetUserName(s)
	return ifu
}

// SetTenantId sets the "tenantId" field.
func (ifu *InventoryFlowUpdate) SetTenantId(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetTenantId()
	ifu.mutation.SetTenantId(i)
	return ifu
}

// AddTenantId adds i to the "tenantId" field.
func (ifu *InventoryFlowUpdate) AddTenantId(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddTenantId(i)
	return ifu
}

// SetTenantName sets the "tenantName" field.
func (ifu *InventoryFlowUpdate) SetTenantName(s string) *InventoryFlowUpdate {
	ifu.mutation.SetTenantName(s)
	return ifu
}

// SetFarmId sets the "farmId" field.
func (ifu *InventoryFlowUpdate) SetFarmId(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetFarmId()
	ifu.mutation.SetFarmId(i)
	return ifu
}

// AddFarmId adds i to the "farmId" field.
func (ifu *InventoryFlowUpdate) AddFarmId(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddFarmId(i)
	return ifu
}

// SetFarmName sets the "farmName" field.
func (ifu *InventoryFlowUpdate) SetFarmName(s string) *InventoryFlowUpdate {
	ifu.mutation.SetFarmName(s)
	return ifu
}

// SetRemarks sets the "remarks" field.
func (ifu *InventoryFlowUpdate) SetRemarks(s string) *InventoryFlowUpdate {
	ifu.mutation.SetRemarks(s)
	return ifu
}

// SetIsChecked sets the "isChecked" field.
func (ifu *InventoryFlowUpdate) SetIsChecked(b bool) *InventoryFlowUpdate {
	ifu.mutation.SetIsChecked(b)
	return ifu
}

// SetReportFileAddress sets the "reportFileAddress" field.
func (ifu *InventoryFlowUpdate) SetReportFileAddress(s string) *InventoryFlowUpdate {
	ifu.mutation.SetReportFileAddress(s)
	return ifu
}

// SetCreatedAt sets the "createdAt" field.
func (ifu *InventoryFlowUpdate) SetCreatedAt(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetCreatedAt()
	ifu.mutation.SetCreatedAt(i)
	return ifu
}

// AddCreatedAt adds i to the "createdAt" field.
func (ifu *InventoryFlowUpdate) AddCreatedAt(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddCreatedAt(i)
	return ifu
}

// SetUpdatedAt sets the "updatedAt" field.
func (ifu *InventoryFlowUpdate) SetUpdatedAt(i int64) *InventoryFlowUpdate {
	ifu.mutation.ResetUpdatedAt()
	ifu.mutation.SetUpdatedAt(i)
	return ifu
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (ifu *InventoryFlowUpdate) AddUpdatedAt(i int64) *InventoryFlowUpdate {
	ifu.mutation.AddUpdatedAt(i)
	return ifu
}

// SetDeleted sets the "deleted" field.
func (ifu *InventoryFlowUpdate) SetDeleted(i int) *InventoryFlowUpdate {
	ifu.mutation.ResetDeleted()
	ifu.mutation.SetDeleted(i)
	return ifu
}

// AddDeleted adds i to the "deleted" field.
func (ifu *InventoryFlowUpdate) AddDeleted(i int) *InventoryFlowUpdate {
	ifu.mutation.AddDeleted(i)
	return ifu
}

// Mutation returns the InventoryFlowMutation object of the builder.
func (ifu *InventoryFlowUpdate) Mutation() *InventoryFlowMutation {
	return ifu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ifu *InventoryFlowUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ifu.hooks) == 0 {
		affected, err = ifu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InventoryFlowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ifu.mutation = mutation
			affected, err = ifu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ifu.hooks) - 1; i >= 0; i-- {
			if ifu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ifu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ifu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ifu *InventoryFlowUpdate) SaveX(ctx context.Context) int {
	affected, err := ifu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ifu *InventoryFlowUpdate) Exec(ctx context.Context) error {
	_, err := ifu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifu *InventoryFlowUpdate) ExecX(ctx context.Context) {
	if err := ifu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ifu *InventoryFlowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   inventoryflow.Table,
			Columns: inventoryflow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inventoryflow.FieldID,
			},
		},
	}
	if ps := ifu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ifu.mutation.SysMaterialId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldSysMaterialId,
		})
	}
	if value, ok := ifu.mutation.AddedSysMaterialId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldSysMaterialId,
		})
	}
	if value, ok := ifu.mutation.MaterialId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldMaterialId,
		})
	}
	if value, ok := ifu.mutation.AddedMaterialId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldMaterialId,
		})
	}
	if value, ok := ifu.mutation.MaterialName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldMaterialName,
		})
	}
	if value, ok := ifu.mutation.MaterialCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldMaterialCode,
		})
	}
	if value, ok := ifu.mutation.SeqNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldSeqNumber,
		})
	}
	if value, ok := ifu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldDate,
		})
	}
	if value, ok := ifu.mutation.AddedDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldDate,
		})
	}
	if value, ok := ifu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldType,
		})
	}
	if value, ok := ifu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldType,
		})
	}
	if value, ok := ifu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldStatus,
		})
	}
	if value, ok := ifu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldCount,
		})
	}
	if value, ok := ifu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldCount,
		})
	}
	if value, ok := ifu.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldUnit,
		})
	}
	if value, ok := ifu.mutation.Before(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldBefore,
		})
	}
	if value, ok := ifu.mutation.AddedBefore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldBefore,
		})
	}
	if value, ok := ifu.mutation.After(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldAfter,
		})
	}
	if value, ok := ifu.mutation.AddedAfter(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldAfter,
		})
	}
	if value, ok := ifu.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldUserName,
		})
	}
	if value, ok := ifu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldTenantId,
		})
	}
	if value, ok := ifu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldTenantId,
		})
	}
	if value, ok := ifu.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldTenantName,
		})
	}
	if value, ok := ifu.mutation.FarmId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldFarmId,
		})
	}
	if value, ok := ifu.mutation.AddedFarmId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldFarmId,
		})
	}
	if value, ok := ifu.mutation.FarmName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldFarmName,
		})
	}
	if value, ok := ifu.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldRemarks,
		})
	}
	if value, ok := ifu.mutation.IsChecked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: inventoryflow.FieldIsChecked,
		})
	}
	if value, ok := ifu.mutation.ReportFileAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldReportFileAddress,
		})
	}
	if value, ok := ifu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldCreatedAt,
		})
	}
	if value, ok := ifu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldCreatedAt,
		})
	}
	if value, ok := ifu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldUpdatedAt,
		})
	}
	if value, ok := ifu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldUpdatedAt,
		})
	}
	if value, ok := ifu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldDeleted,
		})
	}
	if value, ok := ifu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldDeleted,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ifu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inventoryflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// InventoryFlowUpdateOne is the builder for updating a single InventoryFlow entity.
type InventoryFlowUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InventoryFlowMutation
}

// SetSysMaterialId sets the "sysMaterialId" field.
func (ifuo *InventoryFlowUpdateOne) SetSysMaterialId(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetSysMaterialId()
	ifuo.mutation.SetSysMaterialId(i)
	return ifuo
}

// AddSysMaterialId adds i to the "sysMaterialId" field.
func (ifuo *InventoryFlowUpdateOne) AddSysMaterialId(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddSysMaterialId(i)
	return ifuo
}

// SetMaterialId sets the "materialId" field.
func (ifuo *InventoryFlowUpdateOne) SetMaterialId(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetMaterialId()
	ifuo.mutation.SetMaterialId(i)
	return ifuo
}

// AddMaterialId adds i to the "materialId" field.
func (ifuo *InventoryFlowUpdateOne) AddMaterialId(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddMaterialId(i)
	return ifuo
}

// SetMaterialName sets the "materialName" field.
func (ifuo *InventoryFlowUpdateOne) SetMaterialName(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetMaterialName(s)
	return ifuo
}

// SetMaterialCode sets the "materialCode" field.
func (ifuo *InventoryFlowUpdateOne) SetMaterialCode(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetMaterialCode(s)
	return ifuo
}

// SetSeqNumber sets the "seqNumber" field.
func (ifuo *InventoryFlowUpdateOne) SetSeqNumber(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetSeqNumber(s)
	return ifuo
}

// SetDate sets the "date" field.
func (ifuo *InventoryFlowUpdateOne) SetDate(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetDate()
	ifuo.mutation.SetDate(i)
	return ifuo
}

// AddDate adds i to the "date" field.
func (ifuo *InventoryFlowUpdateOne) AddDate(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddDate(i)
	return ifuo
}

// SetType sets the "type" field.
func (ifuo *InventoryFlowUpdateOne) SetType(i int) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetType()
	ifuo.mutation.SetType(i)
	return ifuo
}

// AddType adds i to the "type" field.
func (ifuo *InventoryFlowUpdateOne) AddType(i int) *InventoryFlowUpdateOne {
	ifuo.mutation.AddType(i)
	return ifuo
}

// SetStatus sets the "status" field.
func (ifuo *InventoryFlowUpdateOne) SetStatus(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetStatus(s)
	return ifuo
}

// SetCount sets the "count" field.
func (ifuo *InventoryFlowUpdateOne) SetCount(i int) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetCount()
	ifuo.mutation.SetCount(i)
	return ifuo
}

// AddCount adds i to the "count" field.
func (ifuo *InventoryFlowUpdateOne) AddCount(i int) *InventoryFlowUpdateOne {
	ifuo.mutation.AddCount(i)
	return ifuo
}

// SetUnit sets the "unit" field.
func (ifuo *InventoryFlowUpdateOne) SetUnit(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetUnit(s)
	return ifuo
}

// SetBefore sets the "before" field.
func (ifuo *InventoryFlowUpdateOne) SetBefore(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetBefore()
	ifuo.mutation.SetBefore(i)
	return ifuo
}

// AddBefore adds i to the "before" field.
func (ifuo *InventoryFlowUpdateOne) AddBefore(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddBefore(i)
	return ifuo
}

// SetAfter sets the "after" field.
func (ifuo *InventoryFlowUpdateOne) SetAfter(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetAfter()
	ifuo.mutation.SetAfter(i)
	return ifuo
}

// AddAfter adds i to the "after" field.
func (ifuo *InventoryFlowUpdateOne) AddAfter(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddAfter(i)
	return ifuo
}

// SetUserName sets the "userName" field.
func (ifuo *InventoryFlowUpdateOne) SetUserName(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetUserName(s)
	return ifuo
}

// SetTenantId sets the "tenantId" field.
func (ifuo *InventoryFlowUpdateOne) SetTenantId(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetTenantId()
	ifuo.mutation.SetTenantId(i)
	return ifuo
}

// AddTenantId adds i to the "tenantId" field.
func (ifuo *InventoryFlowUpdateOne) AddTenantId(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddTenantId(i)
	return ifuo
}

// SetTenantName sets the "tenantName" field.
func (ifuo *InventoryFlowUpdateOne) SetTenantName(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetTenantName(s)
	return ifuo
}

// SetFarmId sets the "farmId" field.
func (ifuo *InventoryFlowUpdateOne) SetFarmId(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetFarmId()
	ifuo.mutation.SetFarmId(i)
	return ifuo
}

// AddFarmId adds i to the "farmId" field.
func (ifuo *InventoryFlowUpdateOne) AddFarmId(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddFarmId(i)
	return ifuo
}

// SetFarmName sets the "farmName" field.
func (ifuo *InventoryFlowUpdateOne) SetFarmName(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetFarmName(s)
	return ifuo
}

// SetRemarks sets the "remarks" field.
func (ifuo *InventoryFlowUpdateOne) SetRemarks(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetRemarks(s)
	return ifuo
}

// SetIsChecked sets the "isChecked" field.
func (ifuo *InventoryFlowUpdateOne) SetIsChecked(b bool) *InventoryFlowUpdateOne {
	ifuo.mutation.SetIsChecked(b)
	return ifuo
}

// SetReportFileAddress sets the "reportFileAddress" field.
func (ifuo *InventoryFlowUpdateOne) SetReportFileAddress(s string) *InventoryFlowUpdateOne {
	ifuo.mutation.SetReportFileAddress(s)
	return ifuo
}

// SetCreatedAt sets the "createdAt" field.
func (ifuo *InventoryFlowUpdateOne) SetCreatedAt(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetCreatedAt()
	ifuo.mutation.SetCreatedAt(i)
	return ifuo
}

// AddCreatedAt adds i to the "createdAt" field.
func (ifuo *InventoryFlowUpdateOne) AddCreatedAt(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddCreatedAt(i)
	return ifuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (ifuo *InventoryFlowUpdateOne) SetUpdatedAt(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetUpdatedAt()
	ifuo.mutation.SetUpdatedAt(i)
	return ifuo
}

// AddUpdatedAt adds i to the "updatedAt" field.
func (ifuo *InventoryFlowUpdateOne) AddUpdatedAt(i int64) *InventoryFlowUpdateOne {
	ifuo.mutation.AddUpdatedAt(i)
	return ifuo
}

// SetDeleted sets the "deleted" field.
func (ifuo *InventoryFlowUpdateOne) SetDeleted(i int) *InventoryFlowUpdateOne {
	ifuo.mutation.ResetDeleted()
	ifuo.mutation.SetDeleted(i)
	return ifuo
}

// AddDeleted adds i to the "deleted" field.
func (ifuo *InventoryFlowUpdateOne) AddDeleted(i int) *InventoryFlowUpdateOne {
	ifuo.mutation.AddDeleted(i)
	return ifuo
}

// Mutation returns the InventoryFlowMutation object of the builder.
func (ifuo *InventoryFlowUpdateOne) Mutation() *InventoryFlowMutation {
	return ifuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ifuo *InventoryFlowUpdateOne) Select(field string, fields ...string) *InventoryFlowUpdateOne {
	ifuo.fields = append([]string{field}, fields...)
	return ifuo
}

// Save executes the query and returns the updated InventoryFlow entity.
func (ifuo *InventoryFlowUpdateOne) Save(ctx context.Context) (*InventoryFlow, error) {
	var (
		err  error
		node *InventoryFlow
	)
	if len(ifuo.hooks) == 0 {
		node, err = ifuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InventoryFlowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ifuo.mutation = mutation
			node, err = ifuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ifuo.hooks) - 1; i >= 0; i-- {
			if ifuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ifuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ifuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*InventoryFlow)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from InventoryFlowMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ifuo *InventoryFlowUpdateOne) SaveX(ctx context.Context) *InventoryFlow {
	node, err := ifuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ifuo *InventoryFlowUpdateOne) Exec(ctx context.Context) error {
	_, err := ifuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifuo *InventoryFlowUpdateOne) ExecX(ctx context.Context) {
	if err := ifuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ifuo *InventoryFlowUpdateOne) sqlSave(ctx context.Context) (_node *InventoryFlow, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   inventoryflow.Table,
			Columns: inventoryflow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inventoryflow.FieldID,
			},
		},
	}
	id, ok := ifuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InventoryFlow.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ifuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inventoryflow.FieldID)
		for _, f := range fields {
			if !inventoryflow.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != inventoryflow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ifuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ifuo.mutation.SysMaterialId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldSysMaterialId,
		})
	}
	if value, ok := ifuo.mutation.AddedSysMaterialId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldSysMaterialId,
		})
	}
	if value, ok := ifuo.mutation.MaterialId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldMaterialId,
		})
	}
	if value, ok := ifuo.mutation.AddedMaterialId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldMaterialId,
		})
	}
	if value, ok := ifuo.mutation.MaterialName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldMaterialName,
		})
	}
	if value, ok := ifuo.mutation.MaterialCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldMaterialCode,
		})
	}
	if value, ok := ifuo.mutation.SeqNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldSeqNumber,
		})
	}
	if value, ok := ifuo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldDate,
		})
	}
	if value, ok := ifuo.mutation.AddedDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldDate,
		})
	}
	if value, ok := ifuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldType,
		})
	}
	if value, ok := ifuo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldType,
		})
	}
	if value, ok := ifuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldStatus,
		})
	}
	if value, ok := ifuo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldCount,
		})
	}
	if value, ok := ifuo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldCount,
		})
	}
	if value, ok := ifuo.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldUnit,
		})
	}
	if value, ok := ifuo.mutation.Before(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldBefore,
		})
	}
	if value, ok := ifuo.mutation.AddedBefore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldBefore,
		})
	}
	if value, ok := ifuo.mutation.After(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldAfter,
		})
	}
	if value, ok := ifuo.mutation.AddedAfter(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldAfter,
		})
	}
	if value, ok := ifuo.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldUserName,
		})
	}
	if value, ok := ifuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldTenantId,
		})
	}
	if value, ok := ifuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldTenantId,
		})
	}
	if value, ok := ifuo.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldTenantName,
		})
	}
	if value, ok := ifuo.mutation.FarmId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldFarmId,
		})
	}
	if value, ok := ifuo.mutation.AddedFarmId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldFarmId,
		})
	}
	if value, ok := ifuo.mutation.FarmName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldFarmName,
		})
	}
	if value, ok := ifuo.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldRemarks,
		})
	}
	if value, ok := ifuo.mutation.IsChecked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: inventoryflow.FieldIsChecked,
		})
	}
	if value, ok := ifuo.mutation.ReportFileAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryflow.FieldReportFileAddress,
		})
	}
	if value, ok := ifuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldCreatedAt,
		})
	}
	if value, ok := ifuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldCreatedAt,
		})
	}
	if value, ok := ifuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldUpdatedAt,
		})
	}
	if value, ok := ifuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: inventoryflow.FieldUpdatedAt,
		})
	}
	if value, ok := ifuo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldDeleted,
		})
	}
	if value, ok := ifuo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryflow.FieldDeleted,
		})
	}
	_node = &InventoryFlow{config: ifuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ifuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inventoryflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
