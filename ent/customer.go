// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/customer"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 类型
	Type string `json:"type,omitempty"`
	// 性别
	Gender string `json:"gender,omitempty"`
	// 省
	Province string `json:"province,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty"`
	// 联系地址
	Address string `json:"address,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 开户行
	Bank string `json:"bank,omitempty"`
	// 银行账号
	Account string `json:"account,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenantId,omitempty"`
	// 租户组织名称
	TenantName string `json:"tenantName,omitempty"`
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
func (*Customer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldID, customer.FieldTenantId, customer.FieldCreatedAt, customer.FieldUpdatedAt, customer.FieldDeleted:
			values[i] = new(sql.NullInt64)
		case customer.FieldName, customer.FieldType, customer.FieldGender, customer.FieldProvince, customer.FieldPhone, customer.FieldAddress, customer.FieldEmail, customer.FieldBank, customer.FieldAccount, customer.FieldTenantName, customer.FieldRemarks:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Customer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case customer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case customer.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case customer.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				c.Gender = value.String
			}
		case customer.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				c.Province = value.String
			}
		case customer.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case customer.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = value.String
			}
		case customer.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case customer.FieldBank:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bank", values[i])
			} else if value.Valid {
				c.Bank = value.String
			}
		case customer.FieldAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account", values[i])
			} else if value.Valid {
				c.Account = value.String
			}
		case customer.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				c.TenantId = value.Int64
			}
		case customer.FieldTenantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenantName", values[i])
			} else if value.Valid {
				c.TenantName = value.String
			}
		case customer.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				c.Remarks = value.String
			}
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Int64
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Int64
			}
		case customer.FieldDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				c.Deleted = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(c.Type)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(c.Gender)
	builder.WriteString(", ")
	builder.WriteString("province=")
	builder.WriteString(c.Province)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(c.Address)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("bank=")
	builder.WriteString(c.Bank)
	builder.WriteString(", ")
	builder.WriteString("account=")
	builder.WriteString(c.Account)
	builder.WriteString(", ")
	builder.WriteString("tenantId=")
	builder.WriteString(fmt.Sprintf("%v", c.TenantId))
	builder.WriteString(", ")
	builder.WriteString("tenantName=")
	builder.WriteString(c.TenantName)
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

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
