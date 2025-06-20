// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/event"
	"cattleai/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetEarNumber sets the "earNumber" field.
func (eu *EventUpdate) SetEarNumber(s string) *EventUpdate {
	eu.mutation.SetEarNumber(s)
	return eu
}

// SetEventTypeId sets the "eventTypeId" field.
func (eu *EventUpdate) SetEventTypeId(i int) *EventUpdate {
	eu.mutation.ResetEventTypeId()
	eu.mutation.SetEventTypeId(i)
	return eu
}

// SetNillableEventTypeId sets the "eventTypeId" field if the given value is not nil.
func (eu *EventUpdate) SetNillableEventTypeId(i *int) *EventUpdate {
	if i != nil {
		eu.SetEventTypeId(*i)
	}
	return eu
}

// AddEventTypeId adds i to the "eventTypeId" field.
func (eu *EventUpdate) AddEventTypeId(i int) *EventUpdate {
	eu.mutation.AddEventTypeId(i)
	return eu
}

// ClearEventTypeId clears the value of the "eventTypeId" field.
func (eu *EventUpdate) ClearEventTypeId() *EventUpdate {
	eu.mutation.ClearEventTypeId()
	return eu
}

// SetEventTypeName sets the "eventTypeName" field.
func (eu *EventUpdate) SetEventTypeName(s string) *EventUpdate {
	eu.mutation.SetEventTypeName(s)
	return eu
}

// SetEventSubTypeId sets the "eventSubTypeId" field.
func (eu *EventUpdate) SetEventSubTypeId(i int) *EventUpdate {
	eu.mutation.ResetEventSubTypeId()
	eu.mutation.SetEventSubTypeId(i)
	return eu
}

// SetNillableEventSubTypeId sets the "eventSubTypeId" field if the given value is not nil.
func (eu *EventUpdate) SetNillableEventSubTypeId(i *int) *EventUpdate {
	if i != nil {
		eu.SetEventSubTypeId(*i)
	}
	return eu
}

// AddEventSubTypeId adds i to the "eventSubTypeId" field.
func (eu *EventUpdate) AddEventSubTypeId(i int) *EventUpdate {
	eu.mutation.AddEventSubTypeId(i)
	return eu
}

// ClearEventSubTypeId clears the value of the "eventSubTypeId" field.
func (eu *EventUpdate) ClearEventSubTypeId() *EventUpdate {
	eu.mutation.ClearEventSubTypeId()
	return eu
}

// SetEventSubTypeName sets the "eventSubTypeName" field.
func (eu *EventUpdate) SetEventSubTypeName(s string) *EventUpdate {
	eu.mutation.SetEventSubTypeName(s)
	return eu
}

// SetTenantId sets the "tenantId" field.
func (eu *EventUpdate) SetTenantId(i int64) *EventUpdate {
	eu.mutation.ResetTenantId()
	eu.mutation.SetTenantId(i)
	return eu
}

// AddTenantId adds i to the "tenantId" field.
func (eu *EventUpdate) AddTenantId(i int64) *EventUpdate {
	eu.mutation.AddTenantId(i)
	return eu
}

// SetTenantName sets the "tenantName" field.
func (eu *EventUpdate) SetTenantName(s string) *EventUpdate {
	eu.mutation.SetTenantName(s)
	return eu
}

// SetFarmId sets the "farmId" field.
func (eu *EventUpdate) SetFarmId(i int64) *EventUpdate {
	eu.mutation.ResetFarmId()
	eu.mutation.SetFarmId(i)
	return eu
}

// AddFarmId adds i to the "farmId" field.
func (eu *EventUpdate) AddFarmId(i int64) *EventUpdate {
	eu.mutation.AddFarmId(i)
	return eu
}

// SetFarmName sets the "farmName" field.
func (eu *EventUpdate) SetFarmName(s string) *EventUpdate {
	eu.mutation.SetFarmName(s)
	return eu
}

// SetShedId sets the "shedId" field.
func (eu *EventUpdate) SetShedId(i int64) *EventUpdate {
	eu.mutation.ResetShedId()
	eu.mutation.SetShedId(i)
	return eu
}

// SetNillableShedId sets the "shedId" field if the given value is not nil.
func (eu *EventUpdate) SetNillableShedId(i *int64) *EventUpdate {
	if i != nil {
		eu.SetShedId(*i)
	}
	return eu
}

// AddShedId adds i to the "shedId" field.
func (eu *EventUpdate) AddShedId(i int64) *EventUpdate {
	eu.mutation.AddShedId(i)
	return eu
}

// ClearShedId clears the value of the "shedId" field.
func (eu *EventUpdate) ClearShedId() *EventUpdate {
	eu.mutation.ClearShedId()
	return eu
}

// SetShedName sets the "shedName" field.
func (eu *EventUpdate) SetShedName(s string) *EventUpdate {
	eu.mutation.SetShedName(s)
	return eu
}

// SetNillableShedName sets the "shedName" field if the given value is not nil.
func (eu *EventUpdate) SetNillableShedName(s *string) *EventUpdate {
	if s != nil {
		eu.SetShedName(*s)
	}
	return eu
}

// ClearShedName clears the value of the "shedName" field.
func (eu *EventUpdate) ClearShedName() *EventUpdate {
	eu.mutation.ClearShedName()
	return eu
}

// SetTimes sets the "times" field.
func (eu *EventUpdate) SetTimes(i int) *EventUpdate {
	eu.mutation.ResetTimes()
	eu.mutation.SetTimes(i)
	return eu
}

// SetNillableTimes sets the "times" field if the given value is not nil.
func (eu *EventUpdate) SetNillableTimes(i *int) *EventUpdate {
	if i != nil {
		eu.SetTimes(*i)
	}
	return eu
}

// AddTimes adds i to the "times" field.
func (eu *EventUpdate) AddTimes(i int) *EventUpdate {
	eu.mutation.AddTimes(i)
	return eu
}

// ClearTimes clears the value of the "times" field.
func (eu *EventUpdate) ClearTimes() *EventUpdate {
	eu.mutation.ClearTimes()
	return eu
}

// SetCreatedAt sets the "createdAt" field.
func (eu *EventUpdate) SetCreatedAt(i int64) *EventUpdate {
	eu.mutation.ResetCreatedAt()
	eu.mutation.SetCreatedAt(i)
	return eu
}

// AddCreatedAt adds i to the "createdAt" field.
func (eu *EventUpdate) AddCreatedAt(i int64) *EventUpdate {
	eu.mutation.AddCreatedAt(i)
	return eu
}

// SetDeleted sets the "deleted" field.
func (eu *EventUpdate) SetDeleted(i int) *EventUpdate {
	eu.mutation.ResetDeleted()
	eu.mutation.SetDeleted(i)
	return eu
}

// AddDeleted adds i to the "deleted" field.
func (eu *EventUpdate) AddDeleted(i int) *EventUpdate {
	eu.mutation.AddDeleted(i)
	return eu
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.EarNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEarNumber,
		})
	}
	if value, ok := eu.mutation.EventTypeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldEventTypeId,
		})
	}
	if value, ok := eu.mutation.AddedEventTypeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldEventTypeId,
		})
	}
	if eu.mutation.EventTypeIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldEventTypeId,
		})
	}
	if value, ok := eu.mutation.EventTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventTypeName,
		})
	}
	if value, ok := eu.mutation.EventSubTypeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldEventSubTypeId,
		})
	}
	if value, ok := eu.mutation.AddedEventSubTypeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldEventSubTypeId,
		})
	}
	if eu.mutation.EventSubTypeIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldEventSubTypeId,
		})
	}
	if value, ok := eu.mutation.EventSubTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventSubTypeName,
		})
	}
	if value, ok := eu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldTenantId,
		})
	}
	if value, ok := eu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldTenantId,
		})
	}
	if value, ok := eu.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldTenantName,
		})
	}
	if value, ok := eu.mutation.FarmId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldFarmId,
		})
	}
	if value, ok := eu.mutation.AddedFarmId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldFarmId,
		})
	}
	if value, ok := eu.mutation.FarmName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldFarmName,
		})
	}
	if value, ok := eu.mutation.ShedId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldShedId,
		})
	}
	if value, ok := eu.mutation.AddedShedId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldShedId,
		})
	}
	if eu.mutation.ShedIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: event.FieldShedId,
		})
	}
	if value, ok := eu.mutation.ShedName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldShedName,
		})
	}
	if eu.mutation.ShedNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: event.FieldShedName,
		})
	}
	if value, ok := eu.mutation.Times(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldTimes,
		})
	}
	if value, ok := eu.mutation.AddedTimes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldTimes,
		})
	}
	if eu.mutation.TimesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldTimes,
		})
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := eu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := eu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldDeleted,
		})
	}
	if value, ok := eu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldDeleted,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetEarNumber sets the "earNumber" field.
func (euo *EventUpdateOne) SetEarNumber(s string) *EventUpdateOne {
	euo.mutation.SetEarNumber(s)
	return euo
}

// SetEventTypeId sets the "eventTypeId" field.
func (euo *EventUpdateOne) SetEventTypeId(i int) *EventUpdateOne {
	euo.mutation.ResetEventTypeId()
	euo.mutation.SetEventTypeId(i)
	return euo
}

// SetNillableEventTypeId sets the "eventTypeId" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableEventTypeId(i *int) *EventUpdateOne {
	if i != nil {
		euo.SetEventTypeId(*i)
	}
	return euo
}

// AddEventTypeId adds i to the "eventTypeId" field.
func (euo *EventUpdateOne) AddEventTypeId(i int) *EventUpdateOne {
	euo.mutation.AddEventTypeId(i)
	return euo
}

// ClearEventTypeId clears the value of the "eventTypeId" field.
func (euo *EventUpdateOne) ClearEventTypeId() *EventUpdateOne {
	euo.mutation.ClearEventTypeId()
	return euo
}

// SetEventTypeName sets the "eventTypeName" field.
func (euo *EventUpdateOne) SetEventTypeName(s string) *EventUpdateOne {
	euo.mutation.SetEventTypeName(s)
	return euo
}

// SetEventSubTypeId sets the "eventSubTypeId" field.
func (euo *EventUpdateOne) SetEventSubTypeId(i int) *EventUpdateOne {
	euo.mutation.ResetEventSubTypeId()
	euo.mutation.SetEventSubTypeId(i)
	return euo
}

// SetNillableEventSubTypeId sets the "eventSubTypeId" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableEventSubTypeId(i *int) *EventUpdateOne {
	if i != nil {
		euo.SetEventSubTypeId(*i)
	}
	return euo
}

// AddEventSubTypeId adds i to the "eventSubTypeId" field.
func (euo *EventUpdateOne) AddEventSubTypeId(i int) *EventUpdateOne {
	euo.mutation.AddEventSubTypeId(i)
	return euo
}

// ClearEventSubTypeId clears the value of the "eventSubTypeId" field.
func (euo *EventUpdateOne) ClearEventSubTypeId() *EventUpdateOne {
	euo.mutation.ClearEventSubTypeId()
	return euo
}

// SetEventSubTypeName sets the "eventSubTypeName" field.
func (euo *EventUpdateOne) SetEventSubTypeName(s string) *EventUpdateOne {
	euo.mutation.SetEventSubTypeName(s)
	return euo
}

// SetTenantId sets the "tenantId" field.
func (euo *EventUpdateOne) SetTenantId(i int64) *EventUpdateOne {
	euo.mutation.ResetTenantId()
	euo.mutation.SetTenantId(i)
	return euo
}

// AddTenantId adds i to the "tenantId" field.
func (euo *EventUpdateOne) AddTenantId(i int64) *EventUpdateOne {
	euo.mutation.AddTenantId(i)
	return euo
}

// SetTenantName sets the "tenantName" field.
func (euo *EventUpdateOne) SetTenantName(s string) *EventUpdateOne {
	euo.mutation.SetTenantName(s)
	return euo
}

// SetFarmId sets the "farmId" field.
func (euo *EventUpdateOne) SetFarmId(i int64) *EventUpdateOne {
	euo.mutation.ResetFarmId()
	euo.mutation.SetFarmId(i)
	return euo
}

// AddFarmId adds i to the "farmId" field.
func (euo *EventUpdateOne) AddFarmId(i int64) *EventUpdateOne {
	euo.mutation.AddFarmId(i)
	return euo
}

// SetFarmName sets the "farmName" field.
func (euo *EventUpdateOne) SetFarmName(s string) *EventUpdateOne {
	euo.mutation.SetFarmName(s)
	return euo
}

// SetShedId sets the "shedId" field.
func (euo *EventUpdateOne) SetShedId(i int64) *EventUpdateOne {
	euo.mutation.ResetShedId()
	euo.mutation.SetShedId(i)
	return euo
}

// SetNillableShedId sets the "shedId" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableShedId(i *int64) *EventUpdateOne {
	if i != nil {
		euo.SetShedId(*i)
	}
	return euo
}

// AddShedId adds i to the "shedId" field.
func (euo *EventUpdateOne) AddShedId(i int64) *EventUpdateOne {
	euo.mutation.AddShedId(i)
	return euo
}

// ClearShedId clears the value of the "shedId" field.
func (euo *EventUpdateOne) ClearShedId() *EventUpdateOne {
	euo.mutation.ClearShedId()
	return euo
}

// SetShedName sets the "shedName" field.
func (euo *EventUpdateOne) SetShedName(s string) *EventUpdateOne {
	euo.mutation.SetShedName(s)
	return euo
}

// SetNillableShedName sets the "shedName" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableShedName(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetShedName(*s)
	}
	return euo
}

// ClearShedName clears the value of the "shedName" field.
func (euo *EventUpdateOne) ClearShedName() *EventUpdateOne {
	euo.mutation.ClearShedName()
	return euo
}

// SetTimes sets the "times" field.
func (euo *EventUpdateOne) SetTimes(i int) *EventUpdateOne {
	euo.mutation.ResetTimes()
	euo.mutation.SetTimes(i)
	return euo
}

// SetNillableTimes sets the "times" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableTimes(i *int) *EventUpdateOne {
	if i != nil {
		euo.SetTimes(*i)
	}
	return euo
}

// AddTimes adds i to the "times" field.
func (euo *EventUpdateOne) AddTimes(i int) *EventUpdateOne {
	euo.mutation.AddTimes(i)
	return euo
}

// ClearTimes clears the value of the "times" field.
func (euo *EventUpdateOne) ClearTimes() *EventUpdateOne {
	euo.mutation.ClearTimes()
	return euo
}

// SetCreatedAt sets the "createdAt" field.
func (euo *EventUpdateOne) SetCreatedAt(i int64) *EventUpdateOne {
	euo.mutation.ResetCreatedAt()
	euo.mutation.SetCreatedAt(i)
	return euo
}

// AddCreatedAt adds i to the "createdAt" field.
func (euo *EventUpdateOne) AddCreatedAt(i int64) *EventUpdateOne {
	euo.mutation.AddCreatedAt(i)
	return euo
}

// SetDeleted sets the "deleted" field.
func (euo *EventUpdateOne) SetDeleted(i int) *EventUpdateOne {
	euo.mutation.ResetDeleted()
	euo.mutation.SetDeleted(i)
	return euo
}

// AddDeleted adds i to the "deleted" field.
func (euo *EventUpdateOne) AddDeleted(i int) *EventUpdateOne {
	euo.mutation.AddDeleted(i)
	return euo
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Event)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.EarNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEarNumber,
		})
	}
	if value, ok := euo.mutation.EventTypeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldEventTypeId,
		})
	}
	if value, ok := euo.mutation.AddedEventTypeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldEventTypeId,
		})
	}
	if euo.mutation.EventTypeIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldEventTypeId,
		})
	}
	if value, ok := euo.mutation.EventTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventTypeName,
		})
	}
	if value, ok := euo.mutation.EventSubTypeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldEventSubTypeId,
		})
	}
	if value, ok := euo.mutation.AddedEventSubTypeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldEventSubTypeId,
		})
	}
	if euo.mutation.EventSubTypeIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldEventSubTypeId,
		})
	}
	if value, ok := euo.mutation.EventSubTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventSubTypeName,
		})
	}
	if value, ok := euo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldTenantId,
		})
	}
	if value, ok := euo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldTenantId,
		})
	}
	if value, ok := euo.mutation.TenantName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldTenantName,
		})
	}
	if value, ok := euo.mutation.FarmId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldFarmId,
		})
	}
	if value, ok := euo.mutation.AddedFarmId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldFarmId,
		})
	}
	if value, ok := euo.mutation.FarmName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldFarmName,
		})
	}
	if value, ok := euo.mutation.ShedId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldShedId,
		})
	}
	if value, ok := euo.mutation.AddedShedId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldShedId,
		})
	}
	if euo.mutation.ShedIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: event.FieldShedId,
		})
	}
	if value, ok := euo.mutation.ShedName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldShedName,
		})
	}
	if euo.mutation.ShedNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: event.FieldShedName,
		})
	}
	if value, ok := euo.mutation.Times(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldTimes,
		})
	}
	if value, ok := euo.mutation.AddedTimes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldTimes,
		})
	}
	if euo.mutation.TimesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldTimes,
		})
	}
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := euo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := euo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldDeleted,
		})
	}
	if value, ok := euo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldDeleted,
		})
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
