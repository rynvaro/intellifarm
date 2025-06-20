// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/predicate"
	"cattleai/ent/treatmentresult"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TreatmentResultQuery is the builder for querying TreatmentResult entities.
type TreatmentResultQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TreatmentResult
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TreatmentResultQuery builder.
func (trq *TreatmentResultQuery) Where(ps ...predicate.TreatmentResult) *TreatmentResultQuery {
	trq.predicates = append(trq.predicates, ps...)
	return trq
}

// Limit adds a limit step to the query.
func (trq *TreatmentResultQuery) Limit(limit int) *TreatmentResultQuery {
	trq.limit = &limit
	return trq
}

// Offset adds an offset step to the query.
func (trq *TreatmentResultQuery) Offset(offset int) *TreatmentResultQuery {
	trq.offset = &offset
	return trq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (trq *TreatmentResultQuery) Unique(unique bool) *TreatmentResultQuery {
	trq.unique = &unique
	return trq
}

// Order adds an order step to the query.
func (trq *TreatmentResultQuery) Order(o ...OrderFunc) *TreatmentResultQuery {
	trq.order = append(trq.order, o...)
	return trq
}

// First returns the first TreatmentResult entity from the query.
// Returns a *NotFoundError when no TreatmentResult was found.
func (trq *TreatmentResultQuery) First(ctx context.Context) (*TreatmentResult, error) {
	nodes, err := trq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{treatmentresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (trq *TreatmentResultQuery) FirstX(ctx context.Context) *TreatmentResult {
	node, err := trq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TreatmentResult ID from the query.
// Returns a *NotFoundError when no TreatmentResult ID was found.
func (trq *TreatmentResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{treatmentresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (trq *TreatmentResultQuery) FirstIDX(ctx context.Context) int {
	id, err := trq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TreatmentResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TreatmentResult entity is found.
// Returns a *NotFoundError when no TreatmentResult entities are found.
func (trq *TreatmentResultQuery) Only(ctx context.Context) (*TreatmentResult, error) {
	nodes, err := trq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{treatmentresult.Label}
	default:
		return nil, &NotSingularError{treatmentresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (trq *TreatmentResultQuery) OnlyX(ctx context.Context) *TreatmentResult {
	node, err := trq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TreatmentResult ID in the query.
// Returns a *NotSingularError when more than one TreatmentResult ID is found.
// Returns a *NotFoundError when no entities are found.
func (trq *TreatmentResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{treatmentresult.Label}
	default:
		err = &NotSingularError{treatmentresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (trq *TreatmentResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := trq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TreatmentResults.
func (trq *TreatmentResultQuery) All(ctx context.Context) ([]*TreatmentResult, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return trq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (trq *TreatmentResultQuery) AllX(ctx context.Context) []*TreatmentResult {
	nodes, err := trq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TreatmentResult IDs.
func (trq *TreatmentResultQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := trq.Select(treatmentresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (trq *TreatmentResultQuery) IDsX(ctx context.Context) []int {
	ids, err := trq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (trq *TreatmentResultQuery) Count(ctx context.Context) (int, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return trq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (trq *TreatmentResultQuery) CountX(ctx context.Context) int {
	count, err := trq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (trq *TreatmentResultQuery) Exist(ctx context.Context) (bool, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return trq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (trq *TreatmentResultQuery) ExistX(ctx context.Context) bool {
	exist, err := trq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TreatmentResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (trq *TreatmentResultQuery) Clone() *TreatmentResultQuery {
	if trq == nil {
		return nil
	}
	return &TreatmentResultQuery{
		config:     trq.config,
		limit:      trq.limit,
		offset:     trq.offset,
		order:      append([]OrderFunc{}, trq.order...),
		predicates: append([]predicate.TreatmentResult{}, trq.predicates...),
		// clone intermediate query.
		sql:    trq.sql.Clone(),
		path:   trq.path,
		unique: trq.unique,
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
//	client.TreatmentResult.Query().
//		GroupBy(treatmentresult.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (trq *TreatmentResultQuery) GroupBy(field string, fields ...string) *TreatmentResultGroupBy {
	grbuild := &TreatmentResultGroupBy{config: trq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := trq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return trq.sqlQuery(ctx), nil
	}
	grbuild.label = treatmentresult.Label
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
//	client.TreatmentResult.Query().
//		Select(treatmentresult.FieldName).
//		Scan(ctx, &v)
func (trq *TreatmentResultQuery) Select(fields ...string) *TreatmentResultSelect {
	trq.fields = append(trq.fields, fields...)
	selbuild := &TreatmentResultSelect{TreatmentResultQuery: trq}
	selbuild.label = treatmentresult.Label
	selbuild.flds, selbuild.scan = &trq.fields, selbuild.Scan
	return selbuild
}

func (trq *TreatmentResultQuery) prepareQuery(ctx context.Context) error {
	for _, f := range trq.fields {
		if !treatmentresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if trq.path != nil {
		prev, err := trq.path(ctx)
		if err != nil {
			return err
		}
		trq.sql = prev
	}
	return nil
}

func (trq *TreatmentResultQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TreatmentResult, error) {
	var (
		nodes = []*TreatmentResult{}
		_spec = trq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TreatmentResult).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TreatmentResult{config: trq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, trq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (trq *TreatmentResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := trq.querySpec()
	_spec.Node.Columns = trq.fields
	if len(trq.fields) > 0 {
		_spec.Unique = trq.unique != nil && *trq.unique
	}
	return sqlgraph.CountNodes(ctx, trq.driver, _spec)
}

func (trq *TreatmentResultQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := trq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (trq *TreatmentResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   treatmentresult.Table,
			Columns: treatmentresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: treatmentresult.FieldID,
			},
		},
		From:   trq.sql,
		Unique: true,
	}
	if unique := trq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := trq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, treatmentresult.FieldID)
		for i := range fields {
			if fields[i] != treatmentresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := trq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := trq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := trq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := trq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (trq *TreatmentResultQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(trq.driver.Dialect())
	t1 := builder.Table(treatmentresult.Table)
	columns := trq.fields
	if len(columns) == 0 {
		columns = treatmentresult.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if trq.sql != nil {
		selector = trq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if trq.unique != nil && *trq.unique {
		selector.Distinct()
	}
	for _, p := range trq.predicates {
		p(selector)
	}
	for _, p := range trq.order {
		p(selector)
	}
	if offset := trq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := trq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TreatmentResultGroupBy is the group-by builder for TreatmentResult entities.
type TreatmentResultGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (trgb *TreatmentResultGroupBy) Aggregate(fns ...AggregateFunc) *TreatmentResultGroupBy {
	trgb.fns = append(trgb.fns, fns...)
	return trgb
}

// Scan applies the group-by query and scans the result into the given value.
func (trgb *TreatmentResultGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := trgb.path(ctx)
	if err != nil {
		return err
	}
	trgb.sql = query
	return trgb.sqlScan(ctx, v)
}

func (trgb *TreatmentResultGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range trgb.fields {
		if !treatmentresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := trgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (trgb *TreatmentResultGroupBy) sqlQuery() *sql.Selector {
	selector := trgb.sql.Select()
	aggregation := make([]string, 0, len(trgb.fns))
	for _, fn := range trgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(trgb.fields)+len(trgb.fns))
		for _, f := range trgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(trgb.fields...)...)
}

// TreatmentResultSelect is the builder for selecting fields of TreatmentResult entities.
type TreatmentResultSelect struct {
	*TreatmentResultQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (trs *TreatmentResultSelect) Scan(ctx context.Context, v interface{}) error {
	if err := trs.prepareQuery(ctx); err != nil {
		return err
	}
	trs.sql = trs.TreatmentResultQuery.sqlQuery(ctx)
	return trs.sqlScan(ctx, v)
}

func (trs *TreatmentResultSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := trs.sql.Query()
	if err := trs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
