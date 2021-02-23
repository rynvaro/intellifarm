package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// ShedSetting holds the schema definition for the ShedSetting entity.
type ShedSetting struct {
	ent.Schema
}

// Fields of the ShedSetting.
func (ShedSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("type").Comment("类型"),
		field.Int64("seats").Comment("床位数"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the ShedSetting.
func (ShedSetting) Edges() []ent.Edge {
	return nil
}
