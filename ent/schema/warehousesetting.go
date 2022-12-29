package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// WarehouseSetting holds the schema definition for the WarehouseSetting entity.
type WarehouseSetting struct {
	ent.Schema
}

// Fields of the WarehouseSetting.
func (WarehouseSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("type").Comment("类型"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the WarehouseSetting.
func (WarehouseSetting) Edges() []ent.Edge {
	return nil
}
