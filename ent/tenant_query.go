// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/predicate"
	"cattleai/ent/tenant"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TenantQuery is the builder for querying Tenant entities.
type TenantQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Tenant
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TenantQuery builder.
func (tq *TenantQuery) Where(ps ...predicate.Tenant) *TenantQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TenantQuery) Limit(limit int) *TenantQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TenantQuery) Offset(offset int) *TenantQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TenantQuery) Unique(unique bool) *TenantQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TenantQuery) Order(o ...OrderFunc) *TenantQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// First returns the first Tenant entity from the query.
// Returns a *NotFoundError when no Tenant was found.
func (tq *TenantQuery) First(ctx context.Context) (*Tenant, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tenant.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TenantQuery) FirstX(ctx context.Context) *Tenant {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tenant ID from the query.
// Returns a *NotFoundError when no Tenant ID was found.
func (tq *TenantQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tenant.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TenantQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tenant entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tenant entity is found.
// Returns a *NotFoundError when no Tenant entities are found.
func (tq *TenantQuery) Only(ctx context.Context) (*Tenant, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tenant.Label}
	default:
		return nil, &NotSingularError{tenant.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TenantQuery) OnlyX(ctx context.Context) *Tenant {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tenant ID in the query.
// Returns a *NotSingularError when more than one Tenant ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TenantQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tenant.Label}
	default:
		err = &NotSingularError{tenant.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TenantQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tenants.
func (tq *TenantQuery) All(ctx context.Context) ([]*Tenant, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TenantQuery) AllX(ctx context.Context) []*Tenant {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tenant IDs.
func (tq *TenantQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(tenant.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TenantQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TenantQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TenantQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TenantQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TenantQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TenantQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TenantQuery) Clone() *TenantQuery {
	if tq == nil {
		return nil
	}
	return &TenantQuery{
		config:     tq.config,
		limit:      tq.limit,
		offset:     tq.offset,
		order:      append([]OrderFunc{}, tq.order...),
		predicates: append([]predicate.Tenant{}, tq.predicates...),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Company string `json:"company,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Tenant.Query().
//		GroupBy(tenant.FieldCompany).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TenantQuery) GroupBy(field string, fields ...string) *TenantGroupBy {
	grbuild := &TenantGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = tenant.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Company string `json:"company,omitempty"`
//	}
//
//	client.Tenant.Query().
//		Select(tenant.FieldCompany).
//		Scan(ctx, &v)
func (tq *TenantQuery) Select(fields ...string) *TenantSelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &TenantSelect{TenantQuery: tq}
	selbuild.label = tenant.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
}

func (tq *TenantQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !tenant.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TenantQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tenant, error) {
	var (
		nodes = []*Tenant{}
		_spec = tq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Tenant).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Tenant{config: tq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tq *TenantQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TenantQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *TenantQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tenant.Table,
			Columns: tenant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tenant.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tenant.FieldID)
		for i := range fields {
			if fields[i] != tenant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TenantQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tenant.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = tenant.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TenantGroupBy is the group-by builder for Tenant entities.
type TenantGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TenantGroupBy) Aggregate(fns ...AggregateFunc) *TenantGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TenantGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

func (tgb *TenantGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !tenant.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TenantGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TenantSelect is the builder for selecting fields of Tenant entities.
type TenantSelect struct {
	*TenantQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TenantSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TenantQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

func (ts *TenantSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
