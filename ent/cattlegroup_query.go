// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/cattlegroup"
	"cattleai/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CattleGroupQuery is the builder for querying CattleGroup entities.
type CattleGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CattleGroup
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CattleGroupQuery builder.
func (cgq *CattleGroupQuery) Where(ps ...predicate.CattleGroup) *CattleGroupQuery {
	cgq.predicates = append(cgq.predicates, ps...)
	return cgq
}

// Limit adds a limit step to the query.
func (cgq *CattleGroupQuery) Limit(limit int) *CattleGroupQuery {
	cgq.limit = &limit
	return cgq
}

// Offset adds an offset step to the query.
func (cgq *CattleGroupQuery) Offset(offset int) *CattleGroupQuery {
	cgq.offset = &offset
	return cgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cgq *CattleGroupQuery) Unique(unique bool) *CattleGroupQuery {
	cgq.unique = &unique
	return cgq
}

// Order adds an order step to the query.
func (cgq *CattleGroupQuery) Order(o ...OrderFunc) *CattleGroupQuery {
	cgq.order = append(cgq.order, o...)
	return cgq
}

// First returns the first CattleGroup entity from the query.
// Returns a *NotFoundError when no CattleGroup was found.
func (cgq *CattleGroupQuery) First(ctx context.Context) (*CattleGroup, error) {
	nodes, err := cgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cattlegroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cgq *CattleGroupQuery) FirstX(ctx context.Context) *CattleGroup {
	node, err := cgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CattleGroup ID from the query.
// Returns a *NotFoundError when no CattleGroup ID was found.
func (cgq *CattleGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cattlegroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cgq *CattleGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := cgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CattleGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CattleGroup entity is found.
// Returns a *NotFoundError when no CattleGroup entities are found.
func (cgq *CattleGroupQuery) Only(ctx context.Context) (*CattleGroup, error) {
	nodes, err := cgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cattlegroup.Label}
	default:
		return nil, &NotSingularError{cattlegroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cgq *CattleGroupQuery) OnlyX(ctx context.Context) *CattleGroup {
	node, err := cgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CattleGroup ID in the query.
// Returns a *NotSingularError when more than one CattleGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (cgq *CattleGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cattlegroup.Label}
	default:
		err = &NotSingularError{cattlegroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cgq *CattleGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := cgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CattleGroups.
func (cgq *CattleGroupQuery) All(ctx context.Context) ([]*CattleGroup, error) {
	if err := cgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cgq *CattleGroupQuery) AllX(ctx context.Context) []*CattleGroup {
	nodes, err := cgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CattleGroup IDs.
func (cgq *CattleGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cgq.Select(cattlegroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cgq *CattleGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := cgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cgq *CattleGroupQuery) Count(ctx context.Context) (int, error) {
	if err := cgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cgq *CattleGroupQuery) CountX(ctx context.Context) int {
	count, err := cgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cgq *CattleGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := cgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cgq *CattleGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := cgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CattleGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cgq *CattleGroupQuery) Clone() *CattleGroupQuery {
	if cgq == nil {
		return nil
	}
	return &CattleGroupQuery{
		config:     cgq.config,
		limit:      cgq.limit,
		offset:     cgq.offset,
		order:      append([]OrderFunc{}, cgq.order...),
		predicates: append([]predicate.CattleGroup{}, cgq.predicates...),
		// clone intermediate query.
		sql:    cgq.sql.Clone(),
		path:   cgq.path,
		unique: cgq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EarNumber string `json:"earNumber,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CattleGroup.Query().
//		GroupBy(cattlegroup.FieldEarNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cgq *CattleGroupQuery) GroupBy(field string, fields ...string) *CattleGroupGroupBy {
	grbuild := &CattleGroupGroupBy{config: cgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cgq.sqlQuery(ctx), nil
	}
	grbuild.label = cattlegroup.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EarNumber string `json:"earNumber,omitempty"`
//	}
//
//	client.CattleGroup.Query().
//		Select(cattlegroup.FieldEarNumber).
//		Scan(ctx, &v)
func (cgq *CattleGroupQuery) Select(fields ...string) *CattleGroupSelect {
	cgq.fields = append(cgq.fields, fields...)
	selbuild := &CattleGroupSelect{CattleGroupQuery: cgq}
	selbuild.label = cattlegroup.Label
	selbuild.flds, selbuild.scan = &cgq.fields, selbuild.Scan
	return selbuild
}

func (cgq *CattleGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cgq.fields {
		if !cattlegroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cgq.path != nil {
		prev, err := cgq.path(ctx)
		if err != nil {
			return err
		}
		cgq.sql = prev
	}
	return nil
}

func (cgq *CattleGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CattleGroup, error) {
	var (
		nodes = []*CattleGroup{}
		_spec = cgq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CattleGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CattleGroup{config: cgq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cgq *CattleGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cgq.querySpec()
	_spec.Node.Columns = cgq.fields
	if len(cgq.fields) > 0 {
		_spec.Unique = cgq.unique != nil && *cgq.unique
	}
	return sqlgraph.CountNodes(ctx, cgq.driver, _spec)
}

func (cgq *CattleGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cgq *CattleGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cattlegroup.Table,
			Columns: cattlegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cattlegroup.FieldID,
			},
		},
		From:   cgq.sql,
		Unique: true,
	}
	if unique := cgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cattlegroup.FieldID)
		for i := range fields {
			if fields[i] != cattlegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cgq *CattleGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cgq.driver.Dialect())
	t1 := builder.Table(cattlegroup.Table)
	columns := cgq.fields
	if len(columns) == 0 {
		columns = cattlegroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cgq.sql != nil {
		selector = cgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cgq.unique != nil && *cgq.unique {
		selector.Distinct()
	}
	for _, p := range cgq.predicates {
		p(selector)
	}
	for _, p := range cgq.order {
		p(selector)
	}
	if offset := cgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CattleGroupGroupBy is the group-by builder for CattleGroup entities.
type CattleGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cggb *CattleGroupGroupBy) Aggregate(fns ...AggregateFunc) *CattleGroupGroupBy {
	cggb.fns = append(cggb.fns, fns...)
	return cggb
}

// Scan applies the group-by query and scans the result into the given value.
func (cggb *CattleGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cggb.path(ctx)
	if err != nil {
		return err
	}
	cggb.sql = query
	return cggb.sqlScan(ctx, v)
}

func (cggb *CattleGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cggb.fields {
		if !cattlegroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cggb *CattleGroupGroupBy) sqlQuery() *sql.Selector {
	selector := cggb.sql.Select()
	aggregation := make([]string, 0, len(cggb.fns))
	for _, fn := range cggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cggb.fields)+len(cggb.fns))
		for _, f := range cggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cggb.fields...)...)
}

// CattleGroupSelect is the builder for selecting fields of CattleGroup entities.
type CattleGroupSelect struct {
	*CattleGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cgs *CattleGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cgs.prepareQuery(ctx); err != nil {
		return err
	}
	cgs.sql = cgs.CattleGroupQuery.sqlQuery(ctx)
	return cgs.sqlScan(ctx, v)
}

func (cgs *CattleGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cgs.sql.Query()
	if err := cgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
