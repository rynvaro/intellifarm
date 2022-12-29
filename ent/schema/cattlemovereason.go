package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CattleMoveReason holds the schema definition for the CattleMoveReason entity.
type CattleMoveReason struct {
	ent.Schema
}

// Fields of the CattleMoveReason.
func (CattleMoveReason) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
		field.String("code").Comment("代码"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int("order").Comment("排序数字"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the CattleMoveReason.
func (CattleMoveReason) Edges() []ent.Edge {
	return nil
}
