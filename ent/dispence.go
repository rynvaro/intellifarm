// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/dispence"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Dispence is the model entity for the Dispence schema.
type Dispence struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dispence) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dispence.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Dispence", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dispence fields.
func (d *Dispence) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dispence.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Dispence.
// Note that you need to call Dispence.Unwrap() before calling this method if this Dispence
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dispence) Update() *DispenceUpdateOne {
	return (&DispenceClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Dispence entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dispence) Unwrap() *Dispence {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dispence is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dispence) String() string {
	var builder strings.Builder
	builder.WriteString("Dispence(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Dispences is a parsable slice of Dispence.
type Dispences []*Dispence

func (d Dispences) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
