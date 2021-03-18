package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleGrowsData holds the schema definition for the CattleGrowsData entity.
type CattleGrowsData struct {
	ent.Schema
}

// Fields of the CattleGrowsData.
func (CattleGrowsData) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("cattleId").Comment("牛只ID"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称"),
		field.String("name").Comment("名称").Optional(),
		field.String("earNumber").Comment("牛耳号"),
		field.Int64("measuredAt").Comment("测量日期"),
		field.String("measuredBy").Comment("测量人员"),
		field.Int("weight").Comment("体重(kg)"),
		field.Int("bust").Comment("胸围(cm)"),
		field.Int("height").Comment("体高(cm)"),
		field.Int("score").Comment("体况评分"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the CattleGrowsData.
func (CattleGrowsData) Edges() []ent.Edge {
	return nil
}
