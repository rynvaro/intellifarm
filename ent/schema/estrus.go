package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Estrus holds the schema definition for the Estrus entity.
type Estrus struct {
	ent.Schema
}

// Fields of the Estrus.
func (Estrus) Fields() []ent.Field {
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
		field.Int64("estrusAt"),
		field.Int("estrusTypeId"),
		field.String("estrusTypeName"),
		field.String("userName"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Estrus.
func (Estrus) Edges() []ent.Edge {
	return nil
}
