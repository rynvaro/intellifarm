package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("职务名称"),
		field.String("code").Comment("代码"),
		field.Int64("tenantId").Comment("租户ID").Default(1),
		field.String("tenantName").Comment("租户组织名称").Default("SYSTEM"),
		field.Int("order").Comment("排序数字"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return nil
}
