package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FrozenSemenInfo holds the schema definition for the FrozenSemenInfo entity.
type FrozenSemenInfo struct {
	ent.Schema
}

// Fields of the FrozenSemenInfo.
func (FrozenSemenInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.String("code").Comment("编码"),
		field.String("type").Comment("类型"),
		field.String("regCode").Comment("国际注册号"),
		field.String("bullNumber").Comment("公牛号"),
		field.Int64("birthday").Comment("出生日期"),
		field.String("from").Comment("来源"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the FrozenSemenInfo.
func (FrozenSemenInfo) Edges() []ent.Edge {
	return nil
}
