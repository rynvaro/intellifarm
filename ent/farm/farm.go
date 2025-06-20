// Code generated by ent, DO NOT EDIT.

package farm

const (
	// Label holds the string label denoting the farm type in the database.
	Label = "farm"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldConstructionDate holds the string denoting the constructiondate field in the database.
	FieldConstructionDate = "construction_date"
	// FieldFeedingScale holds the string denoting the feedingscale field in the database.
	FieldFeedingScale = "feeding_scale"
	// FieldContactUser holds the string denoting the contactuser field in the database.
	FieldContactUser = "contact_user"
	// FieldContactPhone holds the string denoting the contactphone field in the database.
	FieldContactPhone = "contact_phone"
	// FieldContactAddress holds the string denoting the contactaddress field in the database.
	FieldContactAddress = "contact_address"
	// FieldSquare holds the string denoting the square field in the database.
	FieldSquare = "square"
	// FieldShedCount holds the string denoting the shedcount field in the database.
	FieldShedCount = "shed_count"
	// FieldCategoryId holds the string denoting the categoryid field in the database.
	FieldCategoryId = "category_id"
	// FieldCategoryName holds the string denoting the categoryname field in the database.
	FieldCategoryName = "category_name"
	// FieldVarietyId holds the string denoting the varietyid field in the database.
	FieldVarietyId = "variety_id"
	// FieldVarietyName holds the string denoting the varietyname field in the database.
	FieldVarietyName = "variety_name"
	// FieldDistrictCode holds the string denoting the districtcode field in the database.
	FieldDistrictCode = "district_code"
	// FieldDistrictName holds the string denoting the districtname field in the database.
	FieldDistrictName = "district_name"
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
	// Table holds the table name of the farm in the database.
	Table = "farms"
)

// Columns holds all SQL columns for farm fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCode,
	FieldConstructionDate,
	FieldFeedingScale,
	FieldContactUser,
	FieldContactPhone,
	FieldContactAddress,
	FieldSquare,
	FieldShedCount,
	FieldCategoryId,
	FieldCategoryName,
	FieldVarietyId,
	FieldVarietyName,
	FieldDistrictCode,
	FieldDistrictName,
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

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// DefaultDeleted holds the default value on creation for the "deleted" field.
	DefaultDeleted int
)
