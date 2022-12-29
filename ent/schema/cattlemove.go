package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CattleMove holds the schema definition for the CattleMove entity.
type CattleMove struct {
	ent.Schema
}

// Fields of the CattleMove.
func (CattleMove) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("cattleId").Comment("牛只ID"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称"),
		field.String("earNumber"),
		field.Int64("date").Comment("日期"),
		field.Int64("fromShedId").Comment("转出的牛舍ID"),
		field.String("fromShed").Comment("转出的栋舍"),
		field.Int64("toShedId").Comment("转到的牛舍ID"),
		field.String("toShed").Comment("转到的栋舍"),
		field.String("userName"),
		field.Int64("reasonId"),
		field.String("reasonName"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the CattleMove.
func (CattleMove) Edges() []ent.Edge {
	return nil
}
