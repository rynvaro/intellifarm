package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// ConcentrateFormula holds the schema definition for the ConcentrateFormula entity.
type ConcentrateFormula struct {
	ent.Schema
}

// Fields of the ConcentrateFormula.
func (ConcentrateFormula) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
		field.Int("status").Comment("1 启用 2 停用"),
		field.Int64("cost").Comment("成本"),
		field.String("data").Comment("JSON格式的配比数据"),
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

// Edges of the ConcentrateFormula.
func (ConcentrateFormula) Edges() []ent.Edge {
	return nil
}
