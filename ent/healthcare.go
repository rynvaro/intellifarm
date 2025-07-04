// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cattleai/ent/healthcare"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// HealthCare is the model entity for the HealthCare schema.
type HealthCare struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 牛只ID
	CattleId int64 `json:"cattleId,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenantId,omitempty"`
	// 租户组织名称
	TenantName string `json:"tenantName,omitempty"`
	// 牛场ID
	FarmId int64 `json:"farmId,omitempty"`
	// 牛场名称
	FarmName string `json:"farmName,omitempty"`
	// 牛舍ID
	ShedId int64 `json:"shedId,omitempty"`
	// 牛舍名称
	ShedName string `json:"shedName,omitempty"`
	// EarNumber holds the value of the "earNumber" field.
	EarNumber string `json:"earNumber,omitempty"`
	// 保健日期
	Date int64 `json:"date,omitempty"`
	// 保健原因
	Reason string `json:"reason,omitempty"`
	// 保健措施
	Method string `json:"method,omitempty"`
	// 兽医名称
	VetName string `json:"vetName,omitempty"`
	// 蹄区，逗号分割
	HoofArea string `json:"hoofArea,omitempty"`
	// 去角方式
	HornMethod string `json:"hornMethod,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"createdAt,omitempty"`
	// 更新时间时间
	UpdatedAt int64 `json:"updatedAt,omitempty"`
	// 是否已删除
	Deleted int `json:"deleted,omitempty"`
	// Remarks holds the value of the "remarks" field.
	Remarks string `json:"remarks,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HealthCare) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case healthcare.FieldID, healthcare.FieldCattleId, healthcare.FieldTenantId, healthcare.FieldFarmId, healthcare.FieldShedId, healthcare.FieldDate, healthcare.FieldCreatedAt, healthcare.FieldUpdatedAt, healthcare.FieldDeleted:
			values[i] = new(sql.NullInt64)
		case healthcare.FieldTenantName, healthcare.FieldFarmName, healthcare.FieldShedName, healthcare.FieldEarNumber, healthcare.FieldReason, healthcare.FieldMethod, healthcare.FieldVetName, healthcare.FieldHoofArea, healthcare.FieldHornMethod, healthcare.FieldRemarks:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type HealthCare", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HealthCare fields.
func (hc *HealthCare) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case healthcare.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			hc.ID = int(value.Int64)
		case healthcare.FieldCattleId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cattleId", values[i])
			} else if value.Valid {
				hc.CattleId = value.Int64
			}
		case healthcare.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				hc.TenantId = value.Int64
			}
		case healthcare.FieldTenantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenantName", values[i])
			} else if value.Valid {
				hc.TenantName = value.String
			}
		case healthcare.FieldFarmId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field farmId", values[i])
			} else if value.Valid {
				hc.FarmId = value.Int64
			}
		case healthcare.FieldFarmName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field farmName", values[i])
			} else if value.Valid {
				hc.FarmName = value.String
			}
		case healthcare.FieldShedId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shedId", values[i])
			} else if value.Valid {
				hc.ShedId = value.Int64
			}
		case healthcare.FieldShedName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shedName", values[i])
			} else if value.Valid {
				hc.ShedName = value.String
			}
		case healthcare.FieldEarNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field earNumber", values[i])
			} else if value.Valid {
				hc.EarNumber = value.String
			}
		case healthcare.FieldDate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				hc.Date = value.Int64
			}
		case healthcare.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				hc.Reason = value.String
			}
		case healthcare.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				hc.Method = value.String
			}
		case healthcare.FieldVetName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vetName", values[i])
			} else if value.Valid {
				hc.VetName = value.String
			}
		case healthcare.FieldHoofArea:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hoofArea", values[i])
			} else if value.Valid {
				hc.HoofArea = value.String
			}
		case healthcare.FieldHornMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hornMethod", values[i])
			} else if value.Valid {
				hc.HornMethod = value.String
			}
		case healthcare.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				hc.CreatedAt = value.Int64
			}
		case healthcare.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				hc.UpdatedAt = value.Int64
			}
		case healthcare.FieldDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				hc.Deleted = int(value.Int64)
			}
		case healthcare.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				hc.Remarks = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this HealthCare.
// Note that you need to call HealthCare.Unwrap() before calling this method if this HealthCare
// was returned from a transaction, and the transaction was committed or rolled back.
func (hc *HealthCare) Update() *HealthCareUpdateOne {
	return (&HealthCareClient{config: hc.config}).UpdateOne(hc)
}

// Unwrap unwraps the HealthCare entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hc *HealthCare) Unwrap() *HealthCare {
	_tx, ok := hc.config.driver.(*txDriver)
	if !ok {
		panic("ent: HealthCare is not a transactional entity")
	}
	hc.config.driver = _tx.drv
	return hc
}

// String implements the fmt.Stringer.
func (hc *HealthCare) String() string {
	var builder strings.Builder
	builder.WriteString("HealthCare(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hc.ID))
	builder.WriteString("cattleId=")
	builder.WriteString(fmt.Sprintf("%v", hc.CattleId))
	builder.WriteString(", ")
	builder.WriteString("tenantId=")
	builder.WriteString(fmt.Sprintf("%v", hc.TenantId))
	builder.WriteString(", ")
	builder.WriteString("tenantName=")
	builder.WriteString(hc.TenantName)
	builder.WriteString(", ")
	builder.WriteString("farmId=")
	builder.WriteString(fmt.Sprintf("%v", hc.FarmId))
	builder.WriteString(", ")
	builder.WriteString("farmName=")
	builder.WriteString(hc.FarmName)
	builder.WriteString(", ")
	builder.WriteString("shedId=")
	builder.WriteString(fmt.Sprintf("%v", hc.ShedId))
	builder.WriteString(", ")
	builder.WriteString("shedName=")
	builder.WriteString(hc.ShedName)
	builder.WriteString(", ")
	builder.WriteString("earNumber=")
	builder.WriteString(hc.EarNumber)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(fmt.Sprintf("%v", hc.Date))
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(hc.Reason)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(hc.Method)
	builder.WriteString(", ")
	builder.WriteString("vetName=")
	builder.WriteString(hc.VetName)
	builder.WriteString(", ")
	builder.WriteString("hoofArea=")
	builder.WriteString(hc.HoofArea)
	builder.WriteString(", ")
	builder.WriteString("hornMethod=")
	builder.WriteString(hc.HornMethod)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(fmt.Sprintf("%v", hc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(fmt.Sprintf("%v", hc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", hc.Deleted))
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(hc.Remarks)
	builder.WriteByte(')')
	return builder.String()
}

// HealthCares is a parsable slice of HealthCare.
type HealthCares []*HealthCare

func (hc HealthCares) config(cfg config) {
	for _i := range hc {
		hc[_i].config = cfg
	}
}
