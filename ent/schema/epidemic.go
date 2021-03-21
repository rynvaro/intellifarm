package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Epidemic holds the schema definition for the Epidemic entity.
type Epidemic struct {
	ent.Schema
}

// Fields of the Epidemic.
func (Epidemic) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("cattleId").Comment("牛只ID"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称"),
		field.String("name").Comment("疫病名称"),
		field.String("earNumber"),
		field.Int64("onset").Comment("发病日期"),
		field.Int("epidemicTypeId"),
		field.String("epidemicTypeName"),
		field.String("isolatedShedName").Comment("隔离栋舍"),
		field.String("diagedBy").Comment("诊断兽医"),
		field.Int("treatmentResultId"),
		field.String("treatmentResultName"),
		field.String("treatmentState"),
		field.Int64("treatmentAt").Comment("治疗结束日期"),
		field.String("whereabout").Comment("去向"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Epidemic.
func (Epidemic) Edges() []ent.Edge {
	return nil
}
