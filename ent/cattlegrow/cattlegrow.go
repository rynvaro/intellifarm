// Code generated by ent, DO NOT EDIT.

package cattlegrow

const (
	// Label holds the string label denoting the cattlegrow type in the database.
	Label = "cattle_grow"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEarNumber holds the string denoting the earnumber field in the database.
	FieldEarNumber = "ear_number"
	// FieldStage holds the string denoting the stage field in the database.
	FieldStage = "stage"
	// FieldDateStart holds the string denoting the datestart field in the database.
	FieldDateStart = "date_start"
	// FieldDateEnd holds the string denoting the dateend field in the database.
	FieldDateEnd = "date_end"
	// FieldWeightStart holds the string denoting the weightstart field in the database.
	FieldWeightStart = "weight_start"
	// FieldWeightEnd holds the string denoting the weightend field in the database.
	FieldWeightEnd = "weight_end"
	// FieldDailyWeight holds the string denoting the dailyweight field in the database.
	FieldDailyWeight = "daily_weight"
	// FieldFeedWeight holds the string denoting the feedweight field in the database.
	FieldFeedWeight = "feed_weight"
	// FieldDailyFeedWeight holds the string denoting the dailyfeedweight field in the database.
	FieldDailyFeedWeight = "daily_feed_weight"
	// FieldConversionRate holds the string denoting the conversionrate field in the database.
	FieldConversionRate = "conversion_rate"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "user_name"
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
	// Table holds the table name of the cattlegrow in the database.
	Table = "cattle_grows"
)

// Columns holds all SQL columns for cattlegrow fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEarNumber,
	FieldStage,
	FieldDateStart,
	FieldDateEnd,
	FieldWeightStart,
	FieldWeightEnd,
	FieldDailyWeight,
	FieldFeedWeight,
	FieldDailyFeedWeight,
	FieldConversionRate,
	FieldUserName,
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
	// EarNumberValidator is a validator for the "earNumber" field. It is called by the builders before save.
	EarNumberValidator func(string) error
	// StageValidator is a validator for the "stage" field. It is called by the builders before save.
	StageValidator func(string) error
)
