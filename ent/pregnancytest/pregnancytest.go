// Code generated by ent, DO NOT EDIT.

package pregnancytest

const (
	// Label holds the string label denoting the pregnancytest type in the database.
	Label = "pregnancy_test"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCattleId holds the string denoting the cattleid field in the database.
	FieldCattleId = "cattle_id"
	// FieldTenantId holds the string denoting the tenantid field in the database.
	FieldTenantId = "tenant_id"
	// FieldTenantName holds the string denoting the tenantname field in the database.
	FieldTenantName = "tenant_name"
	// FieldFarmId holds the string denoting the farmid field in the database.
	FieldFarmId = "farm_id"
	// FieldFarmName holds the string denoting the farmname field in the database.
	FieldFarmName = "farm_name"
	// FieldShedId holds the string denoting the shedid field in the database.
	FieldShedId = "shed_id"
	// FieldShedName holds the string denoting the shedname field in the database.
	FieldShedName = "shed_name"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEarNumber holds the string denoting the earnumber field in the database.
	FieldEarNumber = "ear_number"
	// FieldTimes holds the string denoting the times field in the database.
	FieldTimes = "times"
	// FieldBreedingAt holds the string denoting the breedingat field in the database.
	FieldBreedingAt = "breeding_at"
	// FieldTestAt holds the string denoting the testat field in the database.
	FieldTestAt = "test_at"
	// FieldPregnancyTestTypeId holds the string denoting the pregnancytesttypeid field in the database.
	FieldPregnancyTestTypeId = "pregnancy_test_type_id"
	// FieldPregnancyTestTypeName holds the string denoting the pregnancytesttypename field in the database.
	FieldPregnancyTestTypeName = "pregnancy_test_type_name"
	// FieldPregnancyTestMethodId holds the string denoting the pregnancytestmethodid field in the database.
	FieldPregnancyTestMethodId = "pregnancy_test_method_id"
	// FieldPregnancyTestMethodName holds the string denoting the pregnancytestmethodname field in the database.
	FieldPregnancyTestMethodName = "pregnancy_test_method_name"
	// FieldPregnancyTestResultId holds the string denoting the pregnancytestresultid field in the database.
	FieldPregnancyTestResultId = "pregnancy_test_result_id"
	// FieldPregnancyTestResultName holds the string denoting the pregnancytestresultname field in the database.
	FieldPregnancyTestResultName = "pregnancy_test_result_name"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "user_name"
	// FieldRemarks holds the string denoting the remarks field in the database.
	FieldRemarks = "remarks"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// Table holds the table name of the pregnancytest in the database.
	Table = "pregnancy_tests"
)

// Columns holds all SQL columns for pregnancytest fields.
var Columns = []string{
	FieldID,
	FieldCattleId,
	FieldTenantId,
	FieldTenantName,
	FieldFarmId,
	FieldFarmName,
	FieldShedId,
	FieldShedName,
	FieldName,
	FieldEarNumber,
	FieldTimes,
	FieldBreedingAt,
	FieldTestAt,
	FieldPregnancyTestTypeId,
	FieldPregnancyTestTypeName,
	FieldPregnancyTestMethodId,
	FieldPregnancyTestMethodName,
	FieldPregnancyTestResultId,
	FieldPregnancyTestResultName,
	FieldUserName,
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
