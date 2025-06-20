// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/reproductionparameters"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ReproductionParameters is the model entity for the ReproductionParameters schema.
type ReproductionParameters struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 编码
	Code string `json:"code,omitempty"`
	// 类型
	Value int64 `json:"value,omitempty"`
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
func (*ReproductionParameters) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case reproductionparameters.FieldID, reproductionparameters.FieldValue, reproductionparameters.FieldTenantId, reproductionparameters.FieldCreatedAt, reproductionparameters.FieldUpdatedAt, reproductionparameters.FieldDeleted:
			values[i] = new(sql.NullInt64)
		case reproductionparameters.FieldName, reproductionparameters.FieldCode, reproductionparameters.FieldTenantName, reproductionparameters.FieldRemarks:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ReproductionParameters", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReproductionParameters fields.
func (rp *ReproductionParameters) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reproductionparameters.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rp.ID = int(value.Int64)
		case reproductionparameters.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rp.Name = value.String
			}
		case reproductionparameters.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				rp.Code = value.String
			}
		case reproductionparameters.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				rp.Value = value.Int64
			}
		case reproductionparameters.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				rp.TenantId = value.Int64
			}
		case reproductionparameters.FieldTenantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenantName", values[i])
			} else if value.Valid {
				rp.TenantName = value.String
			}
		case reproductionparameters.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				rp.Remarks = value.String
			}
		case reproductionparameters.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				rp.CreatedAt = value.Int64
			}
		case reproductionparameters.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				rp.UpdatedAt = value.Int64
			}
		case reproductionparameters.FieldDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				rp.Deleted = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ReproductionParameters.
// Note that you need to call ReproductionParameters.Unwrap() before calling this method if this ReproductionParameters
// was returned from a transaction, and the transaction was committed or rolled back.
func (rp *ReproductionParameters) Update() *ReproductionParametersUpdateOne {
	return (&ReproductionParametersClient{config: rp.config}).UpdateOne(rp)
}

// Unwrap unwraps the ReproductionParameters entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rp *ReproductionParameters) Unwrap() *ReproductionParameters {
	_tx, ok := rp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReproductionParameters is not a transactional entity")
	}
	rp.config.driver = _tx.drv
	return rp
}

// String implements the fmt.Stringer.
func (rp *ReproductionParameters) String() string {
	var builder strings.Builder
	builder.WriteString("ReproductionParameters(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rp.ID))
	builder.WriteString("name=")
	builder.WriteString(rp.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(rp.Code)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", rp.Value))
	builder.WriteString(", ")
	builder.WriteString("tenantId=")
	builder.WriteString(fmt.Sprintf("%v", rp.TenantId))
	builder.WriteString(", ")
	builder.WriteString("tenantName=")
	builder.WriteString(rp.TenantName)
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(rp.Remarks)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(fmt.Sprintf("%v", rp.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(fmt.Sprintf("%v", rp.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", rp.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// ReproductionParametersSlice is a parsable slice of ReproductionParameters.
type ReproductionParametersSlice []*ReproductionParameters

func (rp ReproductionParametersSlice) config(cfg config) {
	for _i := range rp {
		rp[_i].config = cfg
	}
}
