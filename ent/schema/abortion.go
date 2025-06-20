package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Abortion holds the schema definition for the Abortion entity.
type Abortion struct {
	ent.Schema
}

// Fields of the Abortion.
func (Abortion) Fields() []ent.Field {
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
		field.Int64("pregnantAt"),
		field.Int64("abortionAt"),
		field.Int("abortionTypeId"),
		field.String("abortionTypeName"),
		field.String("userName"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Abortion.
func (Abortion) Edges() []ent.Edge {
	return nil
}
