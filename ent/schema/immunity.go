package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Immunity holds the schema definition for the Immunity entity.
type Immunity struct {
	ent.Schema
}

// Fields of the Immunity.
func (Immunity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("earNumber"),
		field.String("shedName"),
		field.Int64("date").Comment("日期"),
		field.Int("itemId").Comment("检疫项目ID"),
		field.String("itemName").Comment("检疫项目"),
		field.String("userName").Comment("责任兽医"),
		field.String("drug").Comment("免疫用药"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Immunity.
func (Immunity) Edges() []ent.Edge {
	return nil
}
