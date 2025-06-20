// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/materialtest"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// MaterialTest is the model entity for the MaterialTest schema.
type MaterialTest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// 检测单号
	SeqNumber string `json:"seqNumber,omitempty"`
	// 入库单号
	AddSeqNumber string `json:"addSeqNumber,omitempty"`
	// 检测日期
	Date int64 `json:"date,omitempty"`
	// Type holds the value of the "type" field.
	Type int `json:"type,omitempty"`
	// Category holds the value of the "category" field.
	Category int `json:"category,omitempty"`
	// MaterialCategory holds the value of the "materialCategory" field.
	MaterialCategory int `json:"materialCategory,omitempty"`
	// UserName holds the value of the "userName" field.
	UserName string `json:"userName,omitempty"`
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
func (*MaterialTest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case materialtest.FieldID, materialtest.FieldDate, materialtest.FieldType, materialtest.FieldCategory, materialtest.FieldMaterialCategory, materialtest.FieldTenantId, materialtest.FieldCreatedAt, materialtest.FieldUpdatedAt, materialtest.FieldDeleted:
			values[i] = new(sql.NullInt64)
		case materialtest.FieldName, materialtest.FieldCode, materialtest.FieldSeqNumber, materialtest.FieldAddSeqNumber, materialtest.FieldUserName, materialtest.FieldTenantName, materialtest.FieldRemarks:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MaterialTest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MaterialTest fields.
func (mt *MaterialTest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case materialtest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mt.ID = int(value.Int64)
		case materialtest.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mt.Name = value.String
			}
		case materialtest.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				mt.Code = value.String
			}
		case materialtest.FieldSeqNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seqNumber", values[i])
			} else if value.Valid {
				mt.SeqNumber = value.String
			}
		case materialtest.FieldAddSeqNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field addSeqNumber", values[i])
			} else if value.Valid {
				mt.AddSeqNumber = value.String
			}
		case materialtest.FieldDate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				mt.Date = value.Int64
			}
		case materialtest.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mt.Type = int(value.Int64)
			}
		case materialtest.FieldCategory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				mt.Category = int(value.Int64)
			}
		case materialtest.FieldMaterialCategory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field materialCategory", values[i])
			} else if value.Valid {
				mt.MaterialCategory = int(value.Int64)
			}
		case materialtest.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userName", values[i])
			} else if value.Valid {
				mt.UserName = value.String
			}
		case materialtest.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				mt.TenantId = value.Int64
			}
		case materialtest.FieldTenantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenantName", values[i])
			} else if value.Valid {
				mt.TenantName = value.String
			}
		case materialtest.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				mt.Remarks = value.String
			}
		case materialtest.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				mt.CreatedAt = value.Int64
			}
		case materialtest.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				mt.UpdatedAt = value.Int64
			}
		case materialtest.FieldDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				mt.Deleted = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MaterialTest.
// Note that you need to call MaterialTest.Unwrap() before calling this method if this MaterialTest
// was returned from a transaction, and the transaction was committed or rolled back.
func (mt *MaterialTest) Update() *MaterialTestUpdateOne {
	return (&MaterialTestClient{config: mt.config}).UpdateOne(mt)
}

// Unwrap unwraps the MaterialTest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mt *MaterialTest) Unwrap() *MaterialTest {
	_tx, ok := mt.config.driver.(*txDriver)
	if !ok {
		panic("ent: MaterialTest is not a transactional entity")
	}
	mt.config.driver = _tx.drv
	return mt
}

// String implements the fmt.Stringer.
func (mt *MaterialTest) String() string {
	var builder strings.Builder
	builder.WriteString("MaterialTest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mt.ID))
	builder.WriteString("name=")
	builder.WriteString(mt.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(mt.Code)
	builder.WriteString(", ")
	builder.WriteString("seqNumber=")
	builder.WriteString(mt.SeqNumber)
	builder.WriteString(", ")
	builder.WriteString("addSeqNumber=")
	builder.WriteString(mt.AddSeqNumber)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(fmt.Sprintf("%v", mt.Date))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", mt.Type))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", mt.Category))
	builder.WriteString(", ")
	builder.WriteString("materialCategory=")
	builder.WriteString(fmt.Sprintf("%v", mt.MaterialCategory))
	builder.WriteString(", ")
	builder.WriteString("userName=")
	builder.WriteString(mt.UserName)
	builder.WriteString(", ")
	builder.WriteString("tenantId=")
	builder.WriteString(fmt.Sprintf("%v", mt.TenantId))
	builder.WriteString(", ")
	builder.WriteString("tenantName=")
	builder.WriteString(mt.TenantName)
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(mt.Remarks)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(fmt.Sprintf("%v", mt.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(fmt.Sprintf("%v", mt.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", mt.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// MaterialTests is a parsable slice of MaterialTest.
type MaterialTests []*MaterialTest

func (mt MaterialTests) config(cfg config) {
	for _i := range mt {
		mt[_i].config = cfg
	}
}
