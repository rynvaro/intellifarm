// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldFarmId holds the string denoting the farmid field in the database.
	FieldFarmId = "farm_id"
	// FieldFarmName holds the string denoting the farmname field in the database.
	FieldFarmName = "farm_name"
	// FieldPositionId holds the string denoting the positionid field in the database.
	FieldPositionId = "position_id"
	// FieldPositionName holds the string denoting the positionname field in the database.
	FieldPositionName = "position_name"
	// FieldDutyName holds the string denoting the dutyname field in the database.
	FieldDutyName = "duty_name"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldEducation holds the string denoting the education field in the database.
	FieldEducation = "education"
	// FieldMajor holds the string denoting the major field in the database.
	FieldMajor = "major"
	// FieldJobTitle holds the string denoting the jobtitle field in the database.
	FieldJobTitle = "job_title"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldIdCard holds the string denoting the idcard field in the database.
	FieldIdCard = "id_card"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldOnJobState holds the string denoting the onjobstate field in the database.
	FieldOnJobState = "on_job_state"
	// FieldJoinedAt holds the string denoting the joinedat field in the database.
	FieldJoinedAt = "joined_at"
	// FieldTenantId holds the string denoting the tenantid field in the database.
	FieldTenantId = "tenant_id"
	// FieldTenantName holds the string denoting the tenantname field in the database.
	FieldTenantName = "tenant_name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRemarks holds the string denoting the remarks field in the database.
	FieldRemarks = "remarks"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldLevel,
	FieldFarmId,
	FieldFarmName,
	FieldPositionId,
	FieldPositionName,
	FieldDutyName,
	FieldName,
	FieldGender,
	FieldAge,
	FieldEducation,
	FieldMajor,
	FieldJobTitle,
	FieldPhone,
	FieldIdCard,
	FieldAddress,
	FieldOnJobState,
	FieldJoinedAt,
	FieldTenantId,
	FieldTenantName,
	FieldPassword,
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
	// DefaultOnJobState holds the default value on creation for the "onJobState" field.
	DefaultOnJobState int
	// DefaultDeleted holds the default value on creation for the "deleted" field.
	DefaultDeleted int
)
