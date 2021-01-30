package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// HealthCare holds the schema definition for the HealthCare entity.
type HealthCare struct {
	ent.Schema
}

// Fields of the HealthCare.
func (HealthCare) Fields() []ent.Field {
	return []ent.Field{
		field.String("earNumber"),
		field.String("shedName"),
		field.Int64("date").Comment("保健日期"),
		field.String("reason").Comment("保健原因"),
		field.String("method").Comment("保健措施"),
		field.String("vetName").Comment("兽医名称"),
		field.String("hoofArea").Comment("蹄区，逗号分割"),
		field.String("hornMethod").Comment("去角方式"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间时间"),
		field.Int("deleted").Comment("是否已删除"),
		field.String("remarks"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
	}
}

// Edges of the HealthCare.
func (HealthCare) Edges() []ent.Edge {
	return nil
}
