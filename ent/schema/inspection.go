package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Inspection holds the schema definition for the Inspection entity.
type Inspection struct {
	ent.Schema
}

// Fields of the Inspection.
func (Inspection) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("cattleId").Comment("牛只ID"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称"),
		field.String("name").Comment("名称"),
		field.String("earNumber"),
		field.Int64("date").Comment("日期"),
		field.Int("itemId").Comment("检疫项目ID"),
		field.String("itemName").Comment("检疫项目"),
		field.Int("methodId").Comment("检疫方法ID"),
		field.String("methodName").Comment("检疫方法"),
		field.Int64("byId").Comment("检疫单位ID"),
		field.String("byName").Comment("检疫单位"),
		field.Int64("resultId").Comment("检疫结果ID"),
		field.String("resultName").Comment("检疫结果"),
		field.Int("handleId"),
		field.String("handleName"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Inspection.
func (Inspection) Edges() []ent.Edge {
	return nil
}
