package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleOut holds the schema definition for the CattleOut entity.
type CattleOut struct {
	ent.Schema
}

// Fields of the CattleOut.
func (CattleOut) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("outType").Comment("出栏类型"),
		field.Int64("date").Comment("日期"),
		field.String("to").Comment("去向"),
		field.Int64("weight"),
		field.Int64("cost").Comment("转出金额"),
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

// Edges of the CattleOut.
func (CattleOut) Edges() []ent.Edge {
	return nil
}
