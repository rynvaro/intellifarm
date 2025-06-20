// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/shedcategory"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ShedCategory is the model entity for the ShedCategory schema.
type ShedCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 类别名称
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShedCategory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case shedcategory.FieldID:
			values[i] = new(sql.NullInt64)
		case shedcategory.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ShedCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShedCategory fields.
func (sc *ShedCategory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shedcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case shedcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sc.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ShedCategory.
// Note that you need to call ShedCategory.Unwrap() before calling this method if this ShedCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *ShedCategory) Update() *ShedCategoryUpdateOne {
	return (&ShedCategoryClient{config: sc.config}).UpdateOne(sc)
}

// Unwrap unwraps the ShedCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *ShedCategory) Unwrap() *ShedCategory {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShedCategory is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *ShedCategory) String() string {
	var builder strings.Builder
	builder.WriteString("ShedCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("name=")
	builder.WriteString(sc.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ShedCategories is a parsable slice of ShedCategory.
type ShedCategories []*ShedCategory

func (sc ShedCategories) config(cfg config) {
	for _i := range sc {
		sc[_i].config = cfg
	}
}
