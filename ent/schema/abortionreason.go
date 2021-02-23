package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// AbortionReason holds the schema definition for the AbortionReason entity.
type AbortionReason struct {
	ent.Schema
}

// Fields of the AbortionReason.
func (AbortionReason) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
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

// Edges of the AbortionReason.
func (AbortionReason) Edges() []ent.Edge {
	return nil
}
