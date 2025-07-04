// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/materialtest"
	"cattleai/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MaterialTestQuery is the builder for querying MaterialTest entities.
type MaterialTestQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MaterialTest
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MaterialTestQuery builder.
func (mtq *MaterialTestQuery) Where(ps ...predicate.MaterialTest) *MaterialTestQuery {
	mtq.predicates = append(mtq.predicates, ps...)
	return mtq
}

// Limit adds a limit step to the query.
func (mtq *MaterialTestQuery) Limit(limit int) *MaterialTestQuery {
	mtq.limit = &limit
	return mtq
}

// Offset adds an offset step to the query.
func (mtq *MaterialTestQuery) Offset(offset int) *MaterialTestQuery {
	mtq.offset = &offset
	return mtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mtq *MaterialTestQuery) Unique(unique bool) *MaterialTestQuery {
	mtq.unique = &unique
	return mtq
}

// Order adds an order step to the query.
func (mtq *MaterialTestQuery) Order(o ...OrderFunc) *MaterialTestQuery {
	mtq.order = append(mtq.order, o...)
	return mtq
}

// First returns the first MaterialTest entity from the query.
// Returns a *NotFoundError when no MaterialTest was found.
func (mtq *MaterialTestQuery) First(ctx context.Context) (*MaterialTest, error) {
	nodes, err := mtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{materialtest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mtq *MaterialTestQuery) FirstX(ctx context.Context) *MaterialTest {
	node, err := mtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MaterialTest ID from the query.
// Returns a *NotFoundError when no MaterialTest ID was found.
func (mtq *MaterialTestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{materialtest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mtq *MaterialTestQuery) FirstIDX(ctx context.Context) int {
	id, err := mtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MaterialTest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MaterialTest entity is found.
// Returns a *NotFoundError when no MaterialTest entities are found.
func (mtq *MaterialTestQuery) Only(ctx context.Context) (*MaterialTest, error) {
	nodes, err := mtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{materialtest.Label}
	default:
		return nil, &NotSingularError{materialtest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mtq *MaterialTestQuery) OnlyX(ctx context.Context) *MaterialTest {
	node, err := mtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MaterialTest ID in the query.
// Returns a *NotSingularError when more than one MaterialTest ID is found.
// Returns a *NotFoundError when no entities are found.
func (mtq *MaterialTestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{materialtest.Label}
	default:
		err = &NotSingularError{materialtest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mtq *MaterialTestQuery) OnlyIDX(ctx context.Context) int {
	id, err := mtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MaterialTests.
func (mtq *MaterialTestQuery) All(ctx context.Context) ([]*MaterialTest, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mtq *MaterialTestQuery) AllX(ctx context.Context) []*MaterialTest {
	nodes, err := mtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MaterialTest IDs.
func (mtq *MaterialTestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mtq.Select(materialtest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mtq *MaterialTestQuery) IDsX(ctx context.Context) []int {
	ids, err := mtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mtq *MaterialTestQuery) Count(ctx context.Context) (int, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mtq *MaterialTestQuery) CountX(ctx context.Context) int {
	count, err := mtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mtq *MaterialTestQuery) Exist(ctx context.Context) (bool, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mtq *MaterialTestQuery) ExistX(ctx context.Context) bool {
	exist, err := mtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MaterialTestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mtq *MaterialTestQuery) Clone() *MaterialTestQuery {
	if mtq == nil {
		return nil
	}
	return &MaterialTestQuery{
		config:     mtq.config,
		limit:      mtq.limit,
		offset:     mtq.offset,
		order:      append([]OrderFunc{}, mtq.order...),
		predicates: append([]predicate.MaterialTest{}, mtq.predicates...),
		// clone intermediate query.
		sql:    mtq.sql.Clone(),
		path:   mtq.path,
		unique: mtq.unique,
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
//	client.MaterialTest.Query().
//		GroupBy(materialtest.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mtq *MaterialTestQuery) GroupBy(field string, fields ...string) *MaterialTestGroupBy {
	grbuild := &MaterialTestGroupBy{config: mtq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mtq.sqlQuery(ctx), nil
	}
	grbuild.label = materialtest.Label
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
//	client.MaterialTest.Query().
//		Select(materialtest.FieldName).
//		Scan(ctx, &v)
func (mtq *MaterialTestQuery) Select(fields ...string) *MaterialTestSelect {
	mtq.fields = append(mtq.fields, fields...)
	selbuild := &MaterialTestSelect{MaterialTestQuery: mtq}
	selbuild.label = materialtest.Label
	selbuild.flds, selbuild.scan = &mtq.fields, selbuild.Scan
	return selbuild
}

func (mtq *MaterialTestQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mtq.fields {
		if !materialtest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mtq.path != nil {
		prev, err := mtq.path(ctx)
		if err != nil {
			return err
		}
		mtq.sql = prev
	}
	return nil
}

func (mtq *MaterialTestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MaterialTest, error) {
	var (
		nodes = []*MaterialTest{}
		_spec = mtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*MaterialTest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &MaterialTest{config: mtq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mtq *MaterialTestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mtq.querySpec()
	_spec.Node.Columns = mtq.fields
	if len(mtq.fields) > 0 {
		_spec.Unique = mtq.unique != nil && *mtq.unique
	}
	return sqlgraph.CountNodes(ctx, mtq.driver, _spec)
}

func (mtq *MaterialTestQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mtq *MaterialTestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   materialtest.Table,
			Columns: materialtest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: materialtest.FieldID,
			},
		},
		From:   mtq.sql,
		Unique: true,
	}
	if unique := mtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, materialtest.FieldID)
		for i := range fields {
			if fields[i] != materialtest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mtq *MaterialTestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mtq.driver.Dialect())
	t1 := builder.Table(materialtest.Table)
	columns := mtq.fields
	if len(columns) == 0 {
		columns = materialtest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mtq.sql != nil {
		selector = mtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mtq.unique != nil && *mtq.unique {
		selector.Distinct()
	}
	for _, p := range mtq.predicates {
		p(selector)
	}
	for _, p := range mtq.order {
		p(selector)
	}
	if offset := mtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MaterialTestGroupBy is the group-by builder for MaterialTest entities.
type MaterialTestGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mtgb *MaterialTestGroupBy) Aggregate(fns ...AggregateFunc) *MaterialTestGroupBy {
	mtgb.fns = append(mtgb.fns, fns...)
	return mtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mtgb *MaterialTestGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mtgb.path(ctx)
	if err != nil {
		return err
	}
	mtgb.sql = query
	return mtgb.sqlScan(ctx, v)
}

func (mtgb *MaterialTestGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mtgb.fields {
		if !materialtest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mtgb *MaterialTestGroupBy) sqlQuery() *sql.Selector {
	selector := mtgb.sql.Select()
	aggregation := make([]string, 0, len(mtgb.fns))
	for _, fn := range mtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mtgb.fields)+len(mtgb.fns))
		for _, f := range mtgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mtgb.fields...)...)
}

// MaterialTestSelect is the builder for selecting fields of MaterialTest entities.
type MaterialTestSelect struct {
	*MaterialTestQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mts *MaterialTestSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mts.prepareQuery(ctx); err != nil {
		return err
	}
	mts.sql = mts.MaterialTestQuery.sqlQuery(ctx)
	return mts.sqlScan(ctx, v)
}

func (mts *MaterialTestSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mts.sql.Query()
	if err := mts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
