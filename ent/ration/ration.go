// Code generated by ent, DO NOT EDIT.

package ration

const (
	// Label holds the string label denoting the ration type in the database.
	Label = "ration"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFormulaId holds the string denoting the formulaid field in the database.
	FieldFormulaId = "formula_id"
	// FieldFormulaName holds the string denoting the formulaname field in the database.
	FieldFormulaName = "formula_name"
	// FieldFormulaCode holds the string denoting the formulacode field in the database.
	FieldFormulaCode = "formula_code"
	// FieldInventory holds the string denoting the inventory field in the database.
	FieldInventory = "inventory"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "user_name"
	// FieldTenantId holds the string denoting the tenantid field in the database.
	FieldTenantId = "tenant_id"
	// FieldTenantName holds the string denoting the tenantname field in the database.
	FieldTenantName = "tenant_name"
	// FieldFarmId holds the string denoting the farmid field in the database.
	FieldFarmId = "farm_id"
	// FieldFarmName holds the string denoting the farmname field in the database.
	FieldFarmName = "farm_name"
	// FieldRemarks holds the string denoting the remarks field in the database.
	FieldRemarks = "remarks"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// Table holds the table name of the ration in the database.
	Table = "rations"
)

// Columns holds all SQL columns for ration fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFormulaId,
	FieldFormulaName,
	FieldFormulaCode,
	FieldInventory,
	FieldUserName,
	FieldTenantId,
	FieldTenantName,
	FieldFarmId,
	FieldFarmName,
	FieldRemarks,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleted,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
