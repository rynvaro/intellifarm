package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CattleIn holds the schema definition for the CattleIn entity.
type CattleIn struct {
	ent.Schema
}

// Fields of the CattleIn.
func (CattleIn) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("joinType").Comment("入群类型"),
		field.Int64("date").Comment("日期"),
		field.String("from").Comment("来源"),
		field.Int64("weight"),
		field.Int64("cost").Comment("购入金额"),
		field.Int64("shippingFee").Comment("运输费用"),
		field.String("shippingCode").Comment("运输证号"),
		field.String("userName"),
		field.String("checkCode").Comment("检疫证号"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the CattleIn.
func (CattleIn) Edges() []ent.Edge {
	return nil
}
