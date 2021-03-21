package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Concentrate holds the schema definition for the Concentrate entity.
type Concentrate struct {
	ent.Schema
}

// Fields of the Concentrate.
func (Concentrate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("精料名称"),
		field.Int64("formulaId").Comment("配方Id"),
		field.String("formulaName").Comment("配方名称"),
		field.String("formulaCode").Comment("配方编号"),
		field.Int64("inventory").Comment("库存"),
		field.String("userName"),
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

// Edges of the Concentrate.
func (Concentrate) Edges() []ent.Edge {
	return nil
}
