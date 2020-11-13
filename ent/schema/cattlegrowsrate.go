package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleGrowsRate holds the schema definition for the CattleGrowsRate entity.
type CattleGrowsRate struct {
	ent.Schema
}

// Fields of the CattleGrowsRate.
func (CattleGrowsRate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.String("earNumber").Comment("牛耳号"),
		field.String("shedName").Comment("所在牛舍"),
		field.Int64("ratedAt").Comment("评定日期"),
		field.String("ratedBy").Comment("评定人员"),
		field.Int("rate").Comment("评分"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the CattleGrowsRate.
func (CattleGrowsRate) Edges() []ent.Edge {
	return nil
}
