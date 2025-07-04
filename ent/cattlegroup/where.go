// Code generated by ent, DO NOT EDIT.

package cattlegroup

import (
	"cattleai/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EarNumber applies equality check predicate on the "earNumber" field. It's identical to EarNumberEQ.
func EarNumber(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEarNumber), v))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// ToShed applies equality check predicate on the "toShed" field. It's identical to ToShedEQ.
func ToShed(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToShed), v))
	})
}

// TenantId applies equality check predicate on the "tenantId" field. It's identical to TenantIdEQ.
func TenantId(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// TenantName applies equality check predicate on the "tenantName" field. It's identical to TenantNameEQ.
func TenantName(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantName), v))
	})
}

// Remarks applies equality check predicate on the "remarks" field. It's identical to RemarksEQ.
func Remarks(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemarks), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleted), v))
	})
}

// EarNumberEQ applies the EQ predicate on the "earNumber" field.
func EarNumberEQ(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEarNumber), v))
	})
}

// EarNumberNEQ applies the NEQ predicate on the "earNumber" field.
func EarNumberNEQ(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEarNumber), v))
	})
}

// EarNumberIn applies the In predicate on the "earNumber" field.
func EarNumberIn(vs ...string) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEarNumber), v...))
	})
}

// EarNumberNotIn applies the NotIn predicate on the "earNumber" field.
func EarNumberNotIn(vs ...string) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEarNumber), v...))
	})
}

// EarNumberGT applies the GT predicate on the "earNumber" field.
func EarNumberGT(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEarNumber), v))
	})
}

// EarNumberGTE applies the GTE predicate on the "earNumber" field.
func EarNumberGTE(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEarNumber), v))
	})
}

// EarNumberLT applies the LT predicate on the "earNumber" field.
func EarNumberLT(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEarNumber), v))
	})
}

// EarNumberLTE applies the LTE predicate on the "earNumber" field.
func EarNumberLTE(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEarNumber), v))
	})
}

// EarNumberContains applies the Contains predicate on the "earNumber" field.
func EarNumberContains(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEarNumber), v))
	})
}

// EarNumberHasPrefix applies the HasPrefix predicate on the "earNumber" field.
func EarNumberHasPrefix(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEarNumber), v))
	})
}

// EarNumberHasSuffix applies the HasSuffix predicate on the "earNumber" field.
func EarNumberHasSuffix(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEarNumber), v))
	})
}

// EarNumberEqualFold applies the EqualFold predicate on the "earNumber" field.
func EarNumberEqualFold(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEarNumber), v))
	})
}

// EarNumberContainsFold applies the ContainsFold predicate on the "earNumber" field.
func EarNumberContainsFold(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEarNumber), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...int64) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...int64) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// ToShedEQ applies the EQ predicate on the "toShed" field.
func ToShedEQ(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToShed), v))
	})
}

// ToShedNEQ applies the NEQ predicate on the "toShed" field.
func ToShedNEQ(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToShed), v))
	})
}

// ToShedIn applies the In predicate on the "toShed" field.
func ToShedIn(vs ...string) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldToShed), v...))
	})
}

// ToShedNotIn applies the NotIn predicate on the "toShed" field.
func ToShedNotIn(vs ...string) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldToShed), v...))
	})
}

// ToShedGT applies the GT predicate on the "toShed" field.
func ToShedGT(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToShed), v))
	})
}

// ToShedGTE applies the GTE predicate on the "toShed" field.
func ToShedGTE(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToShed), v))
	})
}

// ToShedLT applies the LT predicate on the "toShed" field.
func ToShedLT(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToShed), v))
	})
}

// ToShedLTE applies the LTE predicate on the "toShed" field.
func ToShedLTE(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToShed), v))
	})
}

// ToShedContains applies the Contains predicate on the "toShed" field.
func ToShedContains(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToShed), v))
	})
}

// ToShedHasPrefix applies the HasPrefix predicate on the "toShed" field.
func ToShedHasPrefix(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToShed), v))
	})
}

// ToShedHasSuffix applies the HasSuffix predicate on the "toShed" field.
func ToShedHasSuffix(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToShed), v))
	})
}

// ToShedEqualFold applies the EqualFold predicate on the "toShed" field.
func ToShedEqualFold(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToShed), v))
	})
}

// ToShedContainsFold applies the ContainsFold predicate on the "toShed" field.
func ToShedContainsFold(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToShed), v))
	})
}

// TenantIdEQ applies the EQ predicate on the "tenantId" field.
func TenantIdEQ(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// TenantIdNEQ applies the NEQ predicate on the "tenantId" field.
func TenantIdNEQ(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantId), v))
	})
}

// TenantIdIn applies the In predicate on the "tenantId" field.
func TenantIdIn(vs ...int64) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTenantId), v...))
	})
}

// TenantIdNotIn applies the NotIn predicate on the "tenantId" field.
func TenantIdNotIn(vs ...int64) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTenantId), v...))
	})
}

// TenantIdGT applies the GT predicate on the "tenantId" field.
func TenantIdGT(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenantId), v))
	})
}

// TenantIdGTE applies the GTE predicate on the "tenantId" field.
func TenantIdGTE(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenantId), v))
	})
}

// TenantIdLT applies the LT predicate on the "tenantId" field.
func TenantIdLT(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenantId), v))
	})
}

// TenantIdLTE applies the LTE predicate on the "tenantId" field.
func TenantIdLTE(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenantId), v))
	})
}

// TenantNameEQ applies the EQ predicate on the "tenantName" field.
func TenantNameEQ(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantName), v))
	})
}

// TenantNameNEQ applies the NEQ predicate on the "tenantName" field.
func TenantNameNEQ(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantName), v))
	})
}

// TenantNameIn applies the In predicate on the "tenantName" field.
func TenantNameIn(vs ...string) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTenantName), v...))
	})
}

// TenantNameNotIn applies the NotIn predicate on the "tenantName" field.
func TenantNameNotIn(vs ...string) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTenantName), v...))
	})
}

// TenantNameGT applies the GT predicate on the "tenantName" field.
func TenantNameGT(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenantName), v))
	})
}

// TenantNameGTE applies the GTE predicate on the "tenantName" field.
func TenantNameGTE(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenantName), v))
	})
}

// TenantNameLT applies the LT predicate on the "tenantName" field.
func TenantNameLT(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenantName), v))
	})
}

// TenantNameLTE applies the LTE predicate on the "tenantName" field.
func TenantNameLTE(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenantName), v))
	})
}

// TenantNameContains applies the Contains predicate on the "tenantName" field.
func TenantNameContains(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTenantName), v))
	})
}

// TenantNameHasPrefix applies the HasPrefix predicate on the "tenantName" field.
func TenantNameHasPrefix(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTenantName), v))
	})
}

// TenantNameHasSuffix applies the HasSuffix predicate on the "tenantName" field.
func TenantNameHasSuffix(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTenantName), v))
	})
}

// TenantNameEqualFold applies the EqualFold predicate on the "tenantName" field.
func TenantNameEqualFold(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTenantName), v))
	})
}

// TenantNameContainsFold applies the ContainsFold predicate on the "tenantName" field.
func TenantNameContainsFold(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTenantName), v))
	})
}

// RemarksEQ applies the EQ predicate on the "remarks" field.
func RemarksEQ(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemarks), v))
	})
}

// RemarksNEQ applies the NEQ predicate on the "remarks" field.
func RemarksNEQ(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemarks), v))
	})
}

// RemarksIn applies the In predicate on the "remarks" field.
func RemarksIn(vs ...string) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemarks), v...))
	})
}

// RemarksNotIn applies the NotIn predicate on the "remarks" field.
func RemarksNotIn(vs ...string) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemarks), v...))
	})
}

// RemarksGT applies the GT predicate on the "remarks" field.
func RemarksGT(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemarks), v))
	})
}

// RemarksGTE applies the GTE predicate on the "remarks" field.
func RemarksGTE(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemarks), v))
	})
}

// RemarksLT applies the LT predicate on the "remarks" field.
func RemarksLT(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemarks), v))
	})
}

// RemarksLTE applies the LTE predicate on the "remarks" field.
func RemarksLTE(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemarks), v))
	})
}

// RemarksContains applies the Contains predicate on the "remarks" field.
func RemarksContains(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemarks), v))
	})
}

// RemarksHasPrefix applies the HasPrefix predicate on the "remarks" field.
func RemarksHasPrefix(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemarks), v))
	})
}

// RemarksHasSuffix applies the HasSuffix predicate on the "remarks" field.
func RemarksHasSuffix(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemarks), v))
	})
}

// RemarksEqualFold applies the EqualFold predicate on the "remarks" field.
func RemarksEqualFold(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemarks), v))
	})
}

// RemarksContainsFold applies the ContainsFold predicate on the "remarks" field.
func RemarksContainsFold(v string) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemarks), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...int64) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...int64) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...int64) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...int64) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v int64) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleted), v))
	})
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleted), v))
	})
}

// DeletedIn applies the In predicate on the "deleted" field.
func DeletedIn(vs ...int) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeleted), v...))
	})
}

// DeletedNotIn applies the NotIn predicate on the "deleted" field.
func DeletedNotIn(vs ...int) predicate.CattleGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeleted), v...))
	})
}

// DeletedGT applies the GT predicate on the "deleted" field.
func DeletedGT(v int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleted), v))
	})
}

// DeletedGTE applies the GTE predicate on the "deleted" field.
func DeletedGTE(v int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleted), v))
	})
}

// DeletedLT applies the LT predicate on the "deleted" field.
func DeletedLT(v int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleted), v))
	})
}

// DeletedLTE applies the LTE predicate on the "deleted" field.
func DeletedLTE(v int) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleted), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CattleGroup) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CattleGroup) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CattleGroup) predicate.CattleGroup {
	return predicate.CattleGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
