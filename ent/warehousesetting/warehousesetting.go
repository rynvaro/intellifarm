// Code generated by ent, DO NOT EDIT.

package warehousesetting

const (
	// Label holds the string label denoting the warehousesetting type in the database.
	Label = "warehouse_setting"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTenantId holds the string denoting the tenantid field in the database.
	FieldTenantId = "tenant_id"
	// FieldTenantName holds the string denoting the tenantname field in the database.
	FieldTenantName = "tenant_name"
	// FieldRemarks holds the string denoting the remarks field in the database.
	FieldRemarks = "remarks"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// Table holds the table name of the warehousesetting in the database.
	Table = "warehouse_settings"
)

// Columns holds all SQL columns for warehousesetting fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldTenantId,
	FieldTenantName,
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
