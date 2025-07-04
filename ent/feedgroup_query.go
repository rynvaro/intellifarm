// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/feedgroup"
	"cattleai/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedGroupQuery is the builder for querying FeedGroup entities.
type FeedGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FeedGroup
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeedGroupQuery builder.
func (fgq *FeedGroupQuery) Where(ps ...predicate.FeedGroup) *FeedGroupQuery {
	fgq.predicates = append(fgq.predicates, ps...)
	return fgq
}

// Limit adds a limit step to the query.
func (fgq *FeedGroupQuery) Limit(limit int) *FeedGroupQuery {
	fgq.limit = &limit
	return fgq
}

// Offset adds an offset step to the query.
func (fgq *FeedGroupQuery) Offset(offset int) *FeedGroupQuery {
	fgq.offset = &offset
	return fgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fgq *FeedGroupQuery) Unique(unique bool) *FeedGroupQuery {
	fgq.unique = &unique
	return fgq
}

// Order adds an order step to the query.
func (fgq *FeedGroupQuery) Order(o ...OrderFunc) *FeedGroupQuery {
	fgq.order = append(fgq.order, o...)
	return fgq
}

// First returns the first FeedGroup entity from the query.
// Returns a *NotFoundError when no FeedGroup was found.
func (fgq *FeedGroupQuery) First(ctx context.Context) (*FeedGroup, error) {
	nodes, err := fgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{feedgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fgq *FeedGroupQuery) FirstX(ctx context.Context) *FeedGroup {
	node, err := fgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FeedGroup ID from the query.
// Returns a *NotFoundError when no FeedGroup ID was found.
func (fgq *FeedGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{feedgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fgq *FeedGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := fgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FeedGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FeedGroup entity is found.
// Returns a *NotFoundError when no FeedGroup entities are found.
func (fgq *FeedGroupQuery) Only(ctx context.Context) (*FeedGroup, error) {
	nodes, err := fgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{feedgroup.Label}
	default:
		return nil, &NotSingularError{feedgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fgq *FeedGroupQuery) OnlyX(ctx context.Context) *FeedGroup {
	node, err := fgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FeedGroup ID in the query.
// Returns a *NotSingularError when more than one FeedGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (fgq *FeedGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{feedgroup.Label}
	default:
		err = &NotSingularError{feedgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fgq *FeedGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := fgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FeedGroups.
func (fgq *FeedGroupQuery) All(ctx context.Context) ([]*FeedGroup, error) {
	if err := fgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fgq *FeedGroupQuery) AllX(ctx context.Context) []*FeedGroup {
	nodes, err := fgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FeedGroup IDs.
func (fgq *FeedGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fgq.Select(feedgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fgq *FeedGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := fgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fgq *FeedGroupQuery) Count(ctx context.Context) (int, error) {
	if err := fgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fgq *FeedGroupQuery) CountX(ctx context.Context) int {
	count, err := fgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fgq *FeedGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := fgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fgq *FeedGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := fgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeedGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fgq *FeedGroupQuery) Clone() *FeedGroupQuery {
	if fgq == nil {
		return nil
	}
	return &FeedGroupQuery{
		config:     fgq.config,
		limit:      fgq.limit,
		offset:     fgq.offset,
		order:      append([]OrderFunc{}, fgq.order...),
		predicates: append([]predicate.FeedGroup{}, fgq.predicates...),
		// clone intermediate query.
		sql:    fgq.sql.Clone(),
		path:   fgq.path,
		unique: fgq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (fgq *FeedGroupQuery) GroupBy(field string, fields ...string) *FeedGroupGroupBy {
	grbuild := &FeedGroupGroupBy{config: fgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fgq.sqlQuery(ctx), nil
	}
	grbuild.label = feedgroup.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (fgq *FeedGroupQuery) Select(fields ...string) *FeedGroupSelect {
	fgq.fields = append(fgq.fields, fields...)
	selbuild := &FeedGroupSelect{FeedGroupQuery: fgq}
	selbuild.label = feedgroup.Label
	selbuild.flds, selbuild.scan = &fgq.fields, selbuild.Scan
	return selbuild
}

func (fgq *FeedGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fgq.fields {
		if !feedgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fgq.path != nil {
		prev, err := fgq.path(ctx)
		if err != nil {
			return err
		}
		fgq.sql = prev
	}
	return nil
}

func (fgq *FeedGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FeedGroup, error) {
	var (
		nodes = []*FeedGroup{}
		_spec = fgq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FeedGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FeedGroup{config: fgq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fgq *FeedGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fgq.querySpec()
	_spec.Node.Columns = fgq.fields
	if len(fgq.fields) > 0 {
		_spec.Unique = fgq.unique != nil && *fgq.unique
	}
	return sqlgraph.CountNodes(ctx, fgq.driver, _spec)
}

func (fgq *FeedGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fgq *FeedGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   feedgroup.Table,
			Columns: feedgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feedgroup.FieldID,
			},
		},
		From:   fgq.sql,
		Unique: true,
	}
	if unique := fgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feedgroup.FieldID)
		for i := range fields {
			if fields[i] != feedgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fgq *FeedGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fgq.driver.Dialect())
	t1 := builder.Table(feedgroup.Table)
	columns := fgq.fields
	if len(columns) == 0 {
		columns = feedgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fgq.sql != nil {
		selector = fgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fgq.unique != nil && *fgq.unique {
		selector.Distinct()
	}
	for _, p := range fgq.predicates {
		p(selector)
	}
	for _, p := range fgq.order {
		p(selector)
	}
	if offset := fgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeedGroupGroupBy is the group-by builder for FeedGroup entities.
type FeedGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fggb *FeedGroupGroupBy) Aggregate(fns ...AggregateFunc) *FeedGroupGroupBy {
	fggb.fns = append(fggb.fns, fns...)
	return fggb
}

// Scan applies the group-by query and scans the result into the given value.
func (fggb *FeedGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fggb.path(ctx)
	if err != nil {
		return err
	}
	fggb.sql = query
	return fggb.sqlScan(ctx, v)
}

func (fggb *FeedGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fggb.fields {
		if !feedgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fggb *FeedGroupGroupBy) sqlQuery() *sql.Selector {
	selector := fggb.sql.Select()
	aggregation := make([]string, 0, len(fggb.fns))
	for _, fn := range fggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fggb.fields)+len(fggb.fns))
		for _, f := range fggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fggb.fields...)...)
}

// FeedGroupSelect is the builder for selecting fields of FeedGroup entities.
type FeedGroupSelect struct {
	*FeedGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fgs *FeedGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fgs.prepareQuery(ctx); err != nil {
		return err
	}
	fgs.sql = fgs.FeedGroupQuery.sqlQuery(ctx)
	return fgs.sqlScan(ctx, v)
}

func (fgs *FeedGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fgs.sql.Query()
	if err := fgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
