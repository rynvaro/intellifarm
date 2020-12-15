package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("租户组织名称"),
		field.Int("enabled").Comment("是否启用"),
		field.String("remarks").Comment("备注"),
		field.Int("deleted").Comment("是否已删除").Default(0),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}
