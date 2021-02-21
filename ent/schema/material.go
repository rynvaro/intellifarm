package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
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
		field.Int("category"),
		field.String("userName"),
		field.Int64("inventory").Comment("库存量"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
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
