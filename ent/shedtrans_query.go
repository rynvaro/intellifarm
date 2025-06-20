// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/predicate"
	"cattleai/ent/shedtrans"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShedTransQuery is the builder for querying ShedTrans entities.
type ShedTransQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ShedTrans
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShedTransQuery builder.
func (stq *ShedTransQuery) Where(ps ...predicate.ShedTrans) *ShedTransQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit adds a limit step to the query.
func (stq *ShedTransQuery) Limit(limit int) *ShedTransQuery {
	stq.limit = &limit
	return stq
}

// Offset adds an offset step to the query.
func (stq *ShedTransQuery) Offset(offset int) *ShedTransQuery {
	stq.offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *ShedTransQuery) Unique(unique bool) *ShedTransQuery {
	stq.unique = &unique
	return stq
}

// Order adds an order step to the query.
func (stq *ShedTransQuery) Order(o ...OrderFunc) *ShedTransQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// First returns the first ShedTrans entity from the query.
// Returns a *NotFoundError when no ShedTrans was found.
func (stq *ShedTransQuery) First(ctx context.Context) (*ShedTrans, error) {
	nodes, err := stq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shedtrans.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *ShedTransQuery) FirstX(ctx context.Context) *ShedTrans {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShedTrans ID from the query.
// Returns a *NotFoundError when no ShedTrans ID was found.
func (stq *ShedTransQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shedtrans.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *ShedTransQuery) FirstIDX(ctx context.Context) int {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShedTrans entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShedTrans entity is found.
// Returns a *NotFoundError when no ShedTrans entities are found.
func (stq *ShedTransQuery) Only(ctx context.Context) (*ShedTrans, error) {
	nodes, err := stq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shedtrans.Label}
	default:
		return nil, &NotSingularError{shedtrans.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *ShedTransQuery) OnlyX(ctx context.Context) *ShedTrans {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShedTrans ID in the query.
// Returns a *NotSingularError when more than one ShedTrans ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *ShedTransQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shedtrans.Label}
	default:
		err = &NotSingularError{shedtrans.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *ShedTransQuery) OnlyIDX(ctx context.Context) int {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShedTransSlice.
func (stq *ShedTransQuery) All(ctx context.Context) ([]*ShedTrans, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return stq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (stq *ShedTransQuery) AllX(ctx context.Context) []*ShedTrans {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShedTrans IDs.
func (stq *ShedTransQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := stq.Select(shedtrans.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *ShedTransQuery) IDsX(ctx context.Context) []int {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *ShedTransQuery) Count(ctx context.Context) (int, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return stq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (stq *ShedTransQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *ShedTransQuery) Exist(ctx context.Context) (bool, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return stq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *ShedTransQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShedTransQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *ShedTransQuery) Clone() *ShedTransQuery {
	if stq == nil {
		return nil
	}
	return &ShedTransQuery{
		config:     stq.config,
		limit:      stq.limit,
		offset:     stq.offset,
		order:      append([]OrderFunc{}, stq.order...),
		predicates: append([]predicate.ShedTrans{}, stq.predicates...),
		// clone intermediate query.
		sql:    stq.sql.Clone(),
		path:   stq.path,
		unique: stq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (stq *ShedTransQuery) GroupBy(field string, fields ...string) *ShedTransGroupBy {
	grbuild := &ShedTransGroupBy{config: stq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return stq.sqlQuery(ctx), nil
	}
	grbuild.label = shedtrans.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (stq *ShedTransQuery) Select(fields ...string) *ShedTransSelect {
	stq.fields = append(stq.fields, fields...)
	selbuild := &ShedTransSelect{ShedTransQuery: stq}
	selbuild.label = shedtrans.Label
	selbuild.flds, selbuild.scan = &stq.fields, selbuild.Scan
	return selbuild
}

func (stq *ShedTransQuery) prepareQuery(ctx context.Context) error {
	for _, f := range stq.fields {
		if !shedtrans.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *ShedTransQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShedTrans, error) {
	var (
		nodes = []*ShedTrans{}
		_spec = stq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ShedTrans).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ShedTrans{config: stq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (stq *ShedTransQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	_spec.Node.Columns = stq.fields
	if len(stq.fields) > 0 {
		_spec.Unique = stq.unique != nil && *stq.unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *ShedTransQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := stq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (stq *ShedTransQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shedtrans.Table,
			Columns: shedtrans.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shedtrans.FieldID,
			},
		},
		From:   stq.sql,
		Unique: true,
	}
	if unique := stq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := stq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shedtrans.FieldID)
		for i := range fields {
			if fields[i] != shedtrans.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *ShedTransQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(shedtrans.Table)
	columns := stq.fields
	if len(columns) == 0 {
		columns = shedtrans.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.unique != nil && *stq.unique {
		selector.Distinct()
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShedTransGroupBy is the group-by builder for ShedTrans entities.
type ShedTransGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *ShedTransGroupBy) Aggregate(fns ...AggregateFunc) *ShedTransGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the group-by query and scans the result into the given value.
func (stgb *ShedTransGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := stgb.path(ctx)
	if err != nil {
		return err
	}
	stgb.sql = query
	return stgb.sqlScan(ctx, v)
}

func (stgb *ShedTransGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range stgb.fields {
		if !shedtrans.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := stgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (stgb *ShedTransGroupBy) sqlQuery() *sql.Selector {
	selector := stgb.sql.Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(stgb.fields)+len(stgb.fns))
		for _, f := range stgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(stgb.fields...)...)
}

// ShedTransSelect is the builder for selecting fields of ShedTrans entities.
type ShedTransSelect struct {
	*ShedTransQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sts *ShedTransSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	sts.sql = sts.ShedTransQuery.sqlQuery(ctx)
	return sts.sqlScan(ctx, v)
}

func (sts *ShedTransSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sts.sql.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
