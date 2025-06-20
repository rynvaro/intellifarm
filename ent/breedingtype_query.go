// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/breedingtype"
	"cattleai/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BreedingTypeQuery is the builder for querying BreedingType entities.
type BreedingTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BreedingType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BreedingTypeQuery builder.
func (btq *BreedingTypeQuery) Where(ps ...predicate.BreedingType) *BreedingTypeQuery {
	btq.predicates = append(btq.predicates, ps...)
	return btq
}

// Limit adds a limit step to the query.
func (btq *BreedingTypeQuery) Limit(limit int) *BreedingTypeQuery {
	btq.limit = &limit
	return btq
}

// Offset adds an offset step to the query.
func (btq *BreedingTypeQuery) Offset(offset int) *BreedingTypeQuery {
	btq.offset = &offset
	return btq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (btq *BreedingTypeQuery) Unique(unique bool) *BreedingTypeQuery {
	btq.unique = &unique
	return btq
}

// Order adds an order step to the query.
func (btq *BreedingTypeQuery) Order(o ...OrderFunc) *BreedingTypeQuery {
	btq.order = append(btq.order, o...)
	return btq
}

// First returns the first BreedingType entity from the query.
// Returns a *NotFoundError when no BreedingType was found.
func (btq *BreedingTypeQuery) First(ctx context.Context) (*BreedingType, error) {
	nodes, err := btq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{breedingtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (btq *BreedingTypeQuery) FirstX(ctx context.Context) *BreedingType {
	node, err := btq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BreedingType ID from the query.
// Returns a *NotFoundError when no BreedingType ID was found.
func (btq *BreedingTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = btq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{breedingtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (btq *BreedingTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := btq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BreedingType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BreedingType entity is found.
// Returns a *NotFoundError when no BreedingType entities are found.
func (btq *BreedingTypeQuery) Only(ctx context.Context) (*BreedingType, error) {
	nodes, err := btq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{breedingtype.Label}
	default:
		return nil, &NotSingularError{breedingtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (btq *BreedingTypeQuery) OnlyX(ctx context.Context) *BreedingType {
	node, err := btq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BreedingType ID in the query.
// Returns a *NotSingularError when more than one BreedingType ID is found.
// Returns a *NotFoundError when no entities are found.
func (btq *BreedingTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = btq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{breedingtype.Label}
	default:
		err = &NotSingularError{breedingtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (btq *BreedingTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := btq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BreedingTypes.
func (btq *BreedingTypeQuery) All(ctx context.Context) ([]*BreedingType, error) {
	if err := btq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return btq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (btq *BreedingTypeQuery) AllX(ctx context.Context) []*BreedingType {
	nodes, err := btq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BreedingType IDs.
func (btq *BreedingTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := btq.Select(breedingtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (btq *BreedingTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := btq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (btq *BreedingTypeQuery) Count(ctx context.Context) (int, error) {
	if err := btq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return btq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (btq *BreedingTypeQuery) CountX(ctx context.Context) int {
	count, err := btq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (btq *BreedingTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := btq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return btq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (btq *BreedingTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := btq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BreedingTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (btq *BreedingTypeQuery) Clone() *BreedingTypeQuery {
	if btq == nil {
		return nil
	}
	return &BreedingTypeQuery{
		config:     btq.config,
		limit:      btq.limit,
		offset:     btq.offset,
		order:      append([]OrderFunc{}, btq.order...),
		predicates: append([]predicate.BreedingType{}, btq.predicates...),
		// clone intermediate query.
		sql:    btq.sql.Clone(),
		path:   btq.path,
		unique: btq.unique,
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
//	client.BreedingType.Query().
//		GroupBy(breedingtype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (btq *BreedingTypeQuery) GroupBy(field string, fields ...string) *BreedingTypeGroupBy {
	grbuild := &BreedingTypeGroupBy{config: btq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := btq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return btq.sqlQuery(ctx), nil
	}
	grbuild.label = breedingtype.Label
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
//	client.BreedingType.Query().
//		Select(breedingtype.FieldName).
//		Scan(ctx, &v)
func (btq *BreedingTypeQuery) Select(fields ...string) *BreedingTypeSelect {
	btq.fields = append(btq.fields, fields...)
	selbuild := &BreedingTypeSelect{BreedingTypeQuery: btq}
	selbuild.label = breedingtype.Label
	selbuild.flds, selbuild.scan = &btq.fields, selbuild.Scan
	return selbuild
}

func (btq *BreedingTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range btq.fields {
		if !breedingtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if btq.path != nil {
		prev, err := btq.path(ctx)
		if err != nil {
			return err
		}
		btq.sql = prev
	}
	return nil
}

func (btq *BreedingTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BreedingType, error) {
	var (
		nodes = []*BreedingType{}
		_spec = btq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BreedingType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BreedingType{config: btq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, btq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (btq *BreedingTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := btq.querySpec()
	_spec.Node.Columns = btq.fields
	if len(btq.fields) > 0 {
		_spec.Unique = btq.unique != nil && *btq.unique
	}
	return sqlgraph.CountNodes(ctx, btq.driver, _spec)
}

func (btq *BreedingTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := btq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (btq *BreedingTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   breedingtype.Table,
			Columns: breedingtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: breedingtype.FieldID,
			},
		},
		From:   btq.sql,
		Unique: true,
	}
	if unique := btq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := btq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, breedingtype.FieldID)
		for i := range fields {
			if fields[i] != breedingtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := btq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := btq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := btq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := btq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (btq *BreedingTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(btq.driver.Dialect())
	t1 := builder.Table(breedingtype.Table)
	columns := btq.fields
	if len(columns) == 0 {
		columns = breedingtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if btq.sql != nil {
		selector = btq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if btq.unique != nil && *btq.unique {
		selector.Distinct()
	}
	for _, p := range btq.predicates {
		p(selector)
	}
	for _, p := range btq.order {
		p(selector)
	}
	if offset := btq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := btq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BreedingTypeGroupBy is the group-by builder for BreedingType entities.
type BreedingTypeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (btgb *BreedingTypeGroupBy) Aggregate(fns ...AggregateFunc) *BreedingTypeGroupBy {
	btgb.fns = append(btgb.fns, fns...)
	return btgb
}

// Scan applies the group-by query and scans the result into the given value.
func (btgb *BreedingTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := btgb.path(ctx)
	if err != nil {
		return err
	}
	btgb.sql = query
	return btgb.sqlScan(ctx, v)
}

func (btgb *BreedingTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range btgb.fields {
		if !breedingtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := btgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := btgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (btgb *BreedingTypeGroupBy) sqlQuery() *sql.Selector {
	selector := btgb.sql.Select()
	aggregation := make([]string, 0, len(btgb.fns))
	for _, fn := range btgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(btgb.fields)+len(btgb.fns))
		for _, f := range btgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(btgb.fields...)...)
}

// BreedingTypeSelect is the builder for selecting fields of BreedingType entities.
type BreedingTypeSelect struct {
	*BreedingTypeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bts *BreedingTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bts.prepareQuery(ctx); err != nil {
		return err
	}
	bts.sql = bts.BreedingTypeQuery.sqlQuery(ctx)
	return bts.sqlScan(ctx, v)
}

func (bts *BreedingTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bts.sql.Query()
	if err := bts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
