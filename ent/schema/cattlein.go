package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleIn holds the schema definition for the CattleIn entity.
type CattleIn struct {
	ent.Schema
}

// Fields of the CattleIn.
func (CattleIn) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("type"),
		field.Int64("date"),
		field.String("from"),
		field.Float32("weight"),
		field.Float32("cost"),
		field.Float32("shippingFee"),
		field.String("transportCertificateNumber"),
		field.String("userName").Comment("运输人"),
		field.String("testCertificateNumber"),
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
