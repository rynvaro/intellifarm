package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Material holds the schema definition for the Material entity.
type Material struct {
	ent.Schema
}

// Fields of the Material.
func (Material) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
		field.Int64("materialId").Comment("基础资料库的物料ID"),
		field.Int("categoryId"),
		field.String("categoryName"),
		field.String("userName"),
		field.Int64("inventory").Comment("库存量").Default(0),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Material.
func (Material) Edges() []ent.Edge {
	return nil
}
