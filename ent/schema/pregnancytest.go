package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// PregnancyTest holds the schema definition for the PregnancyTest entity.
type PregnancyTest struct {
	ent.Schema
}

// Fields of the PregnancyTest.
func (PregnancyTest) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("cattleId").Comment("牛只ID"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称"),
		field.String("name").Optional(),
		field.String("earNumber"),
		field.Int("times").Comment("胎次"),
		field.Int64("breedingAt"),
		field.Int64("testAt"),
		field.Int("pregnancyTestTypeId"),
		field.String("pregnancyTestTypeName"),
		field.Int("pregnancyTestMethodId"),
		field.String("pregnancyTestMethodName"),
		field.Int("pregnancyTestResultId"),
		field.String("pregnancyTestResultName"),
		field.String("userName"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the PregnancyTest.
func (PregnancyTest) Edges() []ent.Edge {
	return nil
}
