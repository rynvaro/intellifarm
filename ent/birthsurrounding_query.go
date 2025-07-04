// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/birthsurrounding"
	"cattleai/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BirthSurroundingQuery is the builder for querying BirthSurrounding entities.
type BirthSurroundingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BirthSurrounding
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BirthSurroundingQuery builder.
func (bsq *BirthSurroundingQuery) Where(ps ...predicate.BirthSurrounding) *BirthSurroundingQuery {
	bsq.predicates = append(bsq.predicates, ps...)
	return bsq
}

// Limit adds a limit step to the query.
func (bsq *BirthSurroundingQuery) Limit(limit int) *BirthSurroundingQuery {
	bsq.limit = &limit
	return bsq
}

// Offset adds an offset step to the query.
func (bsq *BirthSurroundingQuery) Offset(offset int) *BirthSurroundingQuery {
	bsq.offset = &offset
	return bsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bsq *BirthSurroundingQuery) Unique(unique bool) *BirthSurroundingQuery {
	bsq.unique = &unique
	return bsq
}

// Order adds an order step to the query.
func (bsq *BirthSurroundingQuery) Order(o ...OrderFunc) *BirthSurroundingQuery {
	bsq.order = append(bsq.order, o...)
	return bsq
}

// First returns the first BirthSurrounding entity from the query.
// Returns a *NotFoundError when no BirthSurrounding was found.
func (bsq *BirthSurroundingQuery) First(ctx context.Context) (*BirthSurrounding, error) {
	nodes, err := bsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{birthsurrounding.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bsq *BirthSurroundingQuery) FirstX(ctx context.Context) *BirthSurrounding {
	node, err := bsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BirthSurrounding ID from the query.
// Returns a *NotFoundError when no BirthSurrounding ID was found.
func (bsq *BirthSurroundingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{birthsurrounding.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bsq *BirthSurroundingQuery) FirstIDX(ctx context.Context) int {
	id, err := bsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BirthSurrounding entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BirthSurrounding entity is found.
// Returns a *NotFoundError when no BirthSurrounding entities are found.
func (bsq *BirthSurroundingQuery) Only(ctx context.Context) (*BirthSurrounding, error) {
	nodes, err := bsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{birthsurrounding.Label}
	default:
		return nil, &NotSingularError{birthsurrounding.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bsq *BirthSurroundingQuery) OnlyX(ctx context.Context) *BirthSurrounding {
	node, err := bsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BirthSurrounding ID in the query.
// Returns a *NotSingularError when more than one BirthSurrounding ID is found.
// Returns a *NotFoundError when no entities are found.
func (bsq *BirthSurroundingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{birthsurrounding.Label}
	default:
		err = &NotSingularError{birthsurrounding.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bsq *BirthSurroundingQuery) OnlyIDX(ctx context.Context) int {
	id, err := bsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BirthSurroundings.
func (bsq *BirthSurroundingQuery) All(ctx context.Context) ([]*BirthSurrounding, error) {
	if err := bsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bsq *BirthSurroundingQuery) AllX(ctx context.Context) []*BirthSurrounding {
	nodes, err := bsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BirthSurrounding IDs.
func (bsq *BirthSurroundingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bsq.Select(birthsurrounding.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bsq *BirthSurroundingQuery) IDsX(ctx context.Context) []int {
	ids, err := bsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bsq *BirthSurroundingQuery) Count(ctx context.Context) (int, error) {
	if err := bsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bsq *BirthSurroundingQuery) CountX(ctx context.Context) int {
	count, err := bsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bsq *BirthSurroundingQuery) Exist(ctx context.Context) (bool, error) {
	if err := bsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bsq *BirthSurroundingQuery) ExistX(ctx context.Context) bool {
	exist, err := bsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BirthSurroundingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bsq *BirthSurroundingQuery) Clone() *BirthSurroundingQuery {
	if bsq == nil {
		return nil
	}
	return &BirthSurroundingQuery{
		config:     bsq.config,
		limit:      bsq.limit,
		offset:     bsq.offset,
		order:      append([]OrderFunc{}, bsq.order...),
		predicates: append([]predicate.BirthSurrounding{}, bsq.predicates...),
		// clone intermediate query.
		sql:    bsq.sql.Clone(),
		path:   bsq.path,
		unique: bsq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BirthSurrounding.Query().
//		GroupBy(birthsurrounding.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bsq *BirthSurroundingQuery) GroupBy(field string, fields ...string) *BirthSurroundingGroupBy {
	grbuild := &BirthSurroundingGroupBy{config: bsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bsq.sqlQuery(ctx), nil
	}
	grbuild.label = birthsurrounding.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.BirthSurrounding.Query().
//		Select(birthsurrounding.FieldName).
//		Scan(ctx, &v)
func (bsq *BirthSurroundingQuery) Select(fields ...string) *BirthSurroundingSelect {
	bsq.fields = append(bsq.fields, fields...)
	selbuild := &BirthSurroundingSelect{BirthSurroundingQuery: bsq}
	selbuild.label = birthsurrounding.Label
	selbuild.flds, selbuild.scan = &bsq.fields, selbuild.Scan
	return selbuild
}

func (bsq *BirthSurroundingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bsq.fields {
		if !birthsurrounding.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bsq.path != nil {
		prev, err := bsq.path(ctx)
		if err != nil {
			return err
		}
		bsq.sql = prev
	}
	return nil
}

func (bsq *BirthSurroundingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BirthSurrounding, error) {
	var (
		nodes = []*BirthSurrounding{}
		_spec = bsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BirthSurrounding).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BirthSurrounding{config: bsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bsq *BirthSurroundingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bsq.querySpec()
	_spec.Node.Columns = bsq.fields
	if len(bsq.fields) > 0 {
		_spec.Unique = bsq.unique != nil && *bsq.unique
	}
	return sqlgraph.CountNodes(ctx, bsq.driver, _spec)
}

func (bsq *BirthSurroundingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bsq *BirthSurroundingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   birthsurrounding.Table,
			Columns: birthsurrounding.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: birthsurrounding.FieldID,
			},
		},
		From:   bsq.sql,
		Unique: true,
	}
	if unique := bsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, birthsurrounding.FieldID)
		for i := range fields {
			if fields[i] != birthsurrounding.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bsq *BirthSurroundingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bsq.driver.Dialect())
	t1 := builder.Table(birthsurrounding.Table)
	columns := bsq.fields
	if len(columns) == 0 {
		columns = birthsurrounding.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bsq.sql != nil {
		selector = bsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bsq.unique != nil && *bsq.unique {
		selector.Distinct()
	}
	for _, p := range bsq.predicates {
		p(selector)
	}
	for _, p := range bsq.order {
		p(selector)
	}
	if offset := bsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BirthSurroundingGroupBy is the group-by builder for BirthSurrounding entities.
type BirthSurroundingGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bsgb *BirthSurroundingGroupBy) Aggregate(fns ...AggregateFunc) *BirthSurroundingGroupBy {
	bsgb.fns = append(bsgb.fns, fns...)
	return bsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bsgb *BirthSurroundingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bsgb.path(ctx)
	if err != nil {
		return err
	}
	bsgb.sql = query
	return bsgb.sqlScan(ctx, v)
}

func (bsgb *BirthSurroundingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bsgb.fields {
		if !birthsurrounding.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bsgb *BirthSurroundingGroupBy) sqlQuery() *sql.Selector {
	selector := bsgb.sql.Select()
	aggregation := make([]string, 0, len(bsgb.fns))
	for _, fn := range bsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bsgb.fields)+len(bsgb.fns))
		for _, f := range bsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bsgb.fields...)...)
}

// BirthSurroundingSelect is the builder for selecting fields of BirthSurrounding entities.
type BirthSurroundingSelect struct {
	*BirthSurroundingQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bss *BirthSurroundingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bss.prepareQuery(ctx); err != nil {
		return err
	}
	bss.sql = bss.BirthSurroundingQuery.sqlQuery(ctx)
	return bss.sqlScan(ctx, v)
}

func (bss *BirthSurroundingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bss.sql.Query()
	if err := bss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
