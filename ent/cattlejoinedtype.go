// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/cattlejoinedtype"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// CattleJoinedType is the model entity for the CattleJoinedType schema.
type CattleJoinedType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CattleJoinedType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cattlejoinedtype.FieldID:
			values[i] = new(sql.NullInt64)
		case cattlejoinedtype.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CattleJoinedType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CattleJoinedType fields.
func (cjt *CattleJoinedType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cattlejoinedtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cjt.ID = int(value.Int64)
		case cattlejoinedtype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cjt.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CattleJoinedType.
// Note that you need to call CattleJoinedType.Unwrap() before calling this method if this CattleJoinedType
// was returned from a transaction, and the transaction was committed or rolled back.
func (cjt *CattleJoinedType) Update() *CattleJoinedTypeUpdateOne {
	return (&CattleJoinedTypeClient{config: cjt.config}).UpdateOne(cjt)
}

// Unwrap unwraps the CattleJoinedType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cjt *CattleJoinedType) Unwrap() *CattleJoinedType {
	_tx, ok := cjt.config.driver.(*txDriver)
	if !ok {
		panic("ent: CattleJoinedType is not a transactional entity")
	}
	cjt.config.driver = _tx.drv
	return cjt
}

// String implements the fmt.Stringer.
func (cjt *CattleJoinedType) String() string {
	var builder strings.Builder
	builder.WriteString("CattleJoinedType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cjt.ID))
	builder.WriteString("name=")
	builder.WriteString(cjt.Name)
	builder.WriteByte(')')
	return builder.String()
}

// CattleJoinedTypes is a parsable slice of CattleJoinedType.
type CattleJoinedTypes []*CattleJoinedType

func (cjt CattleJoinedTypes) config(cfg config) {
	for _i := range cjt {
		cjt[_i].config = cfg
	}
}
