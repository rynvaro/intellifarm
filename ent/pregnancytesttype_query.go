// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/predicate"
	"cattleai/ent/pregnancytesttype"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PregnancyTestTypeQuery is the builder for querying PregnancyTestType entities.
type PregnancyTestTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PregnancyTestType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PregnancyTestTypeQuery builder.
func (pttq *PregnancyTestTypeQuery) Where(ps ...predicate.PregnancyTestType) *PregnancyTestTypeQuery {
	pttq.predicates = append(pttq.predicates, ps...)
	return pttq
}

// Limit adds a limit step to the query.
func (pttq *PregnancyTestTypeQuery) Limit(limit int) *PregnancyTestTypeQuery {
	pttq.limit = &limit
	return pttq
}

// Offset adds an offset step to the query.
func (pttq *PregnancyTestTypeQuery) Offset(offset int) *PregnancyTestTypeQuery {
	pttq.offset = &offset
	return pttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pttq *PregnancyTestTypeQuery) Unique(unique bool) *PregnancyTestTypeQuery {
	pttq.unique = &unique
	return pttq
}

// Order adds an order step to the query.
func (pttq *PregnancyTestTypeQuery) Order(o ...OrderFunc) *PregnancyTestTypeQuery {
	pttq.order = append(pttq.order, o...)
	return pttq
}

// First returns the first PregnancyTestType entity from the query.
// Returns a *NotFoundError when no PregnancyTestType was found.
func (pttq *PregnancyTestTypeQuery) First(ctx context.Context) (*PregnancyTestType, error) {
	nodes, err := pttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pregnancytesttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pttq *PregnancyTestTypeQuery) FirstX(ctx context.Context) *PregnancyTestType {
	node, err := pttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PregnancyTestType ID from the query.
// Returns a *NotFoundError when no PregnancyTestType ID was found.
func (pttq *PregnancyTestTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pregnancytesttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pttq *PregnancyTestTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := pttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PregnancyTestType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PregnancyTestType entity is found.
// Returns a *NotFoundError when no PregnancyTestType entities are found.
func (pttq *PregnancyTestTypeQuery) Only(ctx context.Context) (*PregnancyTestType, error) {
	nodes, err := pttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pregnancytesttype.Label}
	default:
		return nil, &NotSingularError{pregnancytesttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pttq *PregnancyTestTypeQuery) OnlyX(ctx context.Context) *PregnancyTestType {
	node, err := pttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PregnancyTestType ID in the query.
// Returns a *NotSingularError when more than one PregnancyTestType ID is found.
// Returns a *NotFoundError when no entities are found.
func (pttq *PregnancyTestTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pregnancytesttype.Label}
	default:
		err = &NotSingularError{pregnancytesttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pttq *PregnancyTestTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PregnancyTestTypes.
func (pttq *PregnancyTestTypeQuery) All(ctx context.Context) ([]*PregnancyTestType, error) {
	if err := pttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pttq *PregnancyTestTypeQuery) AllX(ctx context.Context) []*PregnancyTestType {
	nodes, err := pttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PregnancyTestType IDs.
func (pttq *PregnancyTestTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pttq.Select(pregnancytesttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pttq *PregnancyTestTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := pttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pttq *PregnancyTestTypeQuery) Count(ctx context.Context) (int, error) {
	if err := pttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pttq *PregnancyTestTypeQuery) CountX(ctx context.Context) int {
	count, err := pttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pttq *PregnancyTestTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := pttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pttq *PregnancyTestTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := pttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PregnancyTestTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pttq *PregnancyTestTypeQuery) Clone() *PregnancyTestTypeQuery {
	if pttq == nil {
		return nil
	}
	return &PregnancyTestTypeQuery{
		config:     pttq.config,
		limit:      pttq.limit,
		offset:     pttq.offset,
		order:      append([]OrderFunc{}, pttq.order...),
		predicates: append([]predicate.PregnancyTestType{}, pttq.predicates...),
		// clone intermediate query.
		sql:    pttq.sql.Clone(),
		path:   pttq.path,
		unique: pttq.unique,
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
//	client.PregnancyTestType.Query().
//		GroupBy(pregnancytesttype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pttq *PregnancyTestTypeQuery) GroupBy(field string, fields ...string) *PregnancyTestTypeGroupBy {
	grbuild := &PregnancyTestTypeGroupBy{config: pttq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pttq.sqlQuery(ctx), nil
	}
	grbuild.label = pregnancytesttype.Label
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
//	client.PregnancyTestType.Query().
//		Select(pregnancytesttype.FieldName).
//		Scan(ctx, &v)
func (pttq *PregnancyTestTypeQuery) Select(fields ...string) *PregnancyTestTypeSelect {
	pttq.fields = append(pttq.fields, fields...)
	selbuild := &PregnancyTestTypeSelect{PregnancyTestTypeQuery: pttq}
	selbuild.label = pregnancytesttype.Label
	selbuild.flds, selbuild.scan = &pttq.fields, selbuild.Scan
	return selbuild
}

func (pttq *PregnancyTestTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pttq.fields {
		if !pregnancytesttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pttq.path != nil {
		prev, err := pttq.path(ctx)
		if err != nil {
			return err
		}
		pttq.sql = prev
	}
	return nil
}

func (pttq *PregnancyTestTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PregnancyTestType, error) {
	var (
		nodes = []*PregnancyTestType{}
		_spec = pttq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PregnancyTestType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PregnancyTestType{config: pttq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pttq *PregnancyTestTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pttq.querySpec()
	_spec.Node.Columns = pttq.fields
	if len(pttq.fields) > 0 {
		_spec.Unique = pttq.unique != nil && *pttq.unique
	}
	return sqlgraph.CountNodes(ctx, pttq.driver, _spec)
}

func (pttq *PregnancyTestTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pttq *PregnancyTestTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pregnancytesttype.Table,
			Columns: pregnancytesttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pregnancytesttype.FieldID,
			},
		},
		From:   pttq.sql,
		Unique: true,
	}
	if unique := pttq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pregnancytesttype.FieldID)
		for i := range fields {
			if fields[i] != pregnancytesttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pttq *PregnancyTestTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pttq.driver.Dialect())
	t1 := builder.Table(pregnancytesttype.Table)
	columns := pttq.fields
	if len(columns) == 0 {
		columns = pregnancytesttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pttq.sql != nil {
		selector = pttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pttq.unique != nil && *pttq.unique {
		selector.Distinct()
	}
	for _, p := range pttq.predicates {
		p(selector)
	}
	for _, p := range pttq.order {
		p(selector)
	}
	if offset := pttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PregnancyTestTypeGroupBy is the group-by builder for PregnancyTestType entities.
type PregnancyTestTypeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pttgb *PregnancyTestTypeGroupBy) Aggregate(fns ...AggregateFunc) *PregnancyTestTypeGroupBy {
	pttgb.fns = append(pttgb.fns, fns...)
	return pttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pttgb *PregnancyTestTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pttgb.path(ctx)
	if err != nil {
		return err
	}
	pttgb.sql = query
	return pttgb.sqlScan(ctx, v)
}

func (pttgb *PregnancyTestTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pttgb.fields {
		if !pregnancytesttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pttgb *PregnancyTestTypeGroupBy) sqlQuery() *sql.Selector {
	selector := pttgb.sql.Select()
	aggregation := make([]string, 0, len(pttgb.fns))
	for _, fn := range pttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pttgb.fields)+len(pttgb.fns))
		for _, f := range pttgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pttgb.fields...)...)
}

// PregnancyTestTypeSelect is the builder for selecting fields of PregnancyTestType entities.
type PregnancyTestTypeSelect struct {
	*PregnancyTestTypeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ptts *PregnancyTestTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ptts.prepareQuery(ctx); err != nil {
		return err
	}
	ptts.sql = ptts.PregnancyTestTypeQuery.sqlQuery(ctx)
	return ptts.sqlScan(ctx, v)
}

func (ptts *PregnancyTestTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ptts.sql.Query()
	if err := ptts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
