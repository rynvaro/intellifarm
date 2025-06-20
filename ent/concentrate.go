// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/concentrate"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Concentrate is the model entity for the Concentrate schema.
type Concentrate struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 精料名称
	Name string `json:"name,omitempty"`
	// 配方Id
	FormulaId int64 `json:"formulaId,omitempty"`
	// 配方名称
	FormulaName string `json:"formulaName,omitempty"`
	// 配方编号
	FormulaCode string `json:"formulaCode,omitempty"`
	// 库存
	Inventory int64 `json:"inventory,omitempty"`
	// UserName holds the value of the "userName" field.
	UserName string `json:"userName,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenantId,omitempty"`
	// 租户组织名称
	TenantName string `json:"tenantName,omitempty"`
	// 牛场ID
	FarmId int64 `json:"farmId,omitempty"`
	// 牛场名称
	FarmName string `json:"farmName,omitempty"`
	// 备注
	Remarks string `json:"remarks,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"createdAt,omitempty"`
	// 更新时间
	UpdatedAt int64 `json:"updatedAt,omitempty"`
	// 是否已删除
	Deleted int `json:"deleted,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Concentrate) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case concentrate.FieldID, concentrate.FieldFormulaId, concentrate.FieldInventory, concentrate.FieldTenantId, concentrate.FieldFarmId, concentrate.FieldCreatedAt, concentrate.FieldUpdatedAt, concentrate.FieldDeleted:
			values[i] = new(sql.NullInt64)
		case concentrate.FieldName, concentrate.FieldFormulaName, concentrate.FieldFormulaCode, concentrate.FieldUserName, concentrate.FieldTenantName, concentrate.FieldFarmName, concentrate.FieldRemarks:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Concentrate", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Concentrate fields.
func (c *Concentrate) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case concentrate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case concentrate.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case concentrate.FieldFormulaId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field formulaId", values[i])
			} else if value.Valid {
				c.FormulaId = value.Int64
			}
		case concentrate.FieldFormulaName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field formulaName", values[i])
			} else if value.Valid {
				c.FormulaName = value.String
			}
		case concentrate.FieldFormulaCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field formulaCode", values[i])
			} else if value.Valid {
				c.FormulaCode = value.String
			}
		case concentrate.FieldInventory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field inventory", values[i])
			} else if value.Valid {
				c.Inventory = value.Int64
			}
		case concentrate.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userName", values[i])
			} else if value.Valid {
				c.UserName = value.String
			}
		case concentrate.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				c.TenantId = value.Int64
			}
		case concentrate.FieldTenantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenantName", values[i])
			} else if value.Valid {
				c.TenantName = value.String
			}
		case concentrate.FieldFarmId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field farmId", values[i])
			} else if value.Valid {
				c.FarmId = value.Int64
			}
		case concentrate.FieldFarmName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field farmName", values[i])
			} else if value.Valid {
				c.FarmName = value.String
			}
		case concentrate.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				c.Remarks = value.String
			}
		case concentrate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Int64
			}
		case concentrate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Int64
			}
		case concentrate.FieldDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				c.Deleted = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Concentrate.
// Note that you need to call Concentrate.Unwrap() before calling this method if this Concentrate
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Concentrate) Update() *ConcentrateUpdateOne {
	return (&ConcentrateClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Concentrate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Concentrate) Unwrap() *Concentrate {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Concentrate is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Concentrate) String() string {
	var builder strings.Builder
	builder.WriteString("Concentrate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("formulaId=")
	builder.WriteString(fmt.Sprintf("%v", c.FormulaId))
	builder.WriteString(", ")
	builder.WriteString("formulaName=")
	builder.WriteString(c.FormulaName)
	builder.WriteString(", ")
	builder.WriteString("formulaCode=")
	builder.WriteString(c.FormulaCode)
	builder.WriteString(", ")
	builder.WriteString("inventory=")
	builder.WriteString(fmt.Sprintf("%v", c.Inventory))
	builder.WriteString(", ")
	builder.WriteString("userName=")
	builder.WriteString(c.UserName)
	builder.WriteString(", ")
	builder.WriteString("tenantId=")
	builder.WriteString(fmt.Sprintf("%v", c.TenantId))
	builder.WriteString(", ")
	builder.WriteString("tenantName=")
	builder.WriteString(c.TenantName)
	builder.WriteString(", ")
	builder.WriteString("farmId=")
	builder.WriteString(fmt.Sprintf("%v", c.FarmId))
	builder.WriteString(", ")
	builder.WriteString("farmName=")
	builder.WriteString(c.FarmName)
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(c.Remarks)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", c.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// Concentrates is a parsable slice of Concentrate.
type Concentrates []*Concentrate

func (c Concentrates) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
