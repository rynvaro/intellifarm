// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/feedrecord"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// FeedRecord is the model entity for the FeedRecord schema.
type FeedRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ShedName holds the value of the "shedName" field.
	ShedName string `json:"shedName,omitempty"`
	// Date holds the value of the "date" field.
	Date int64 `json:"date,omitempty"`
	// RationCode holds the value of the "rationCode" field.
	RationCode string `json:"rationCode,omitempty"`
	// RationName holds the value of the "rationName" field.
	RationName string `json:"rationName,omitempty"`
	// 饲喂量
	RationAmount int64 `json:"rationAmount,omitempty"`
	// 饲喂头数
	Count int64 `json:"count,omitempty"`
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
func (*FeedRecord) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case feedrecord.FieldID, feedrecord.FieldDate, feedrecord.FieldRationAmount, feedrecord.FieldCount, feedrecord.FieldTenantId, feedrecord.FieldCreatedAt, feedrecord.FieldUpdatedAt, feedrecord.FieldDeleted:
			values[i] = new(sql.NullInt64)
		case feedrecord.FieldName, feedrecord.FieldShedName, feedrecord.FieldRationCode, feedrecord.FieldRationName, feedrecord.FieldUserName, feedrecord.FieldTenantName, feedrecord.FieldRemarks:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FeedRecord", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeedRecord fields.
func (fr *FeedRecord) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feedrecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fr.ID = int(value.Int64)
		case feedrecord.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fr.Name = value.String
			}
		case feedrecord.FieldShedName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shedName", values[i])
			} else if value.Valid {
				fr.ShedName = value.String
			}
		case feedrecord.FieldDate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				fr.Date = value.Int64
			}
		case feedrecord.FieldRationCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rationCode", values[i])
			} else if value.Valid {
				fr.RationCode = value.String
			}
		case feedrecord.FieldRationName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rationName", values[i])
			} else if value.Valid {
				fr.RationName = value.String
			}
		case feedrecord.FieldRationAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rationAmount", values[i])
			} else if value.Valid {
				fr.RationAmount = value.Int64
			}
		case feedrecord.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				fr.Count = value.Int64
			}
		case feedrecord.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userName", values[i])
			} else if value.Valid {
				fr.UserName = value.String
			}
		case feedrecord.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				fr.TenantId = value.Int64
			}
		case feedrecord.FieldTenantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenantName", values[i])
			} else if value.Valid {
				fr.TenantName = value.String
			}
		case feedrecord.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				fr.Remarks = value.String
			}
		case feedrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				fr.CreatedAt = value.Int64
			}
		case feedrecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				fr.UpdatedAt = value.Int64
			}
		case feedrecord.FieldDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				fr.Deleted = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this FeedRecord.
// Note that you need to call FeedRecord.Unwrap() before calling this method if this FeedRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *FeedRecord) Update() *FeedRecordUpdateOne {
	return (&FeedRecordClient{config: fr.config}).UpdateOne(fr)
}

// Unwrap unwraps the FeedRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fr *FeedRecord) Unwrap() *FeedRecord {
	_tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeedRecord is not a transactional entity")
	}
	fr.config.driver = _tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *FeedRecord) String() string {
	var builder strings.Builder
	builder.WriteString("FeedRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fr.ID))
	builder.WriteString("name=")
	builder.WriteString(fr.Name)
	builder.WriteString(", ")
	builder.WriteString("shedName=")
	builder.WriteString(fr.ShedName)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(fmt.Sprintf("%v", fr.Date))
	builder.WriteString(", ")
	builder.WriteString("rationCode=")
	builder.WriteString(fr.RationCode)
	builder.WriteString(", ")
	builder.WriteString("rationName=")
	builder.WriteString(fr.RationName)
	builder.WriteString(", ")
	builder.WriteString("rationAmount=")
	builder.WriteString(fmt.Sprintf("%v", fr.RationAmount))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", fr.Count))
	builder.WriteString(", ")
	builder.WriteString("userName=")
	builder.WriteString(fr.UserName)
	builder.WriteString(", ")
	builder.WriteString("tenantId=")
	builder.WriteString(fmt.Sprintf("%v", fr.TenantId))
	builder.WriteString(", ")
	builder.WriteString("tenantName=")
	builder.WriteString(fr.TenantName)
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(fr.Remarks)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(fmt.Sprintf("%v", fr.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(fmt.Sprintf("%v", fr.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", fr.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// FeedRecords is a parsable slice of FeedRecord.
type FeedRecords []*FeedRecord

func (fr FeedRecords) config(cfg config) {
	for _i := range fr {
		fr[_i].config = cfg
	}
}
