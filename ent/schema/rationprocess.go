package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// RationProcess holds the schema definition for the RationProcess entity.
type RationProcess struct {
	ent.Schema
}

// Fields of the RationProcess.
func (RationProcess) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("rationId").Comment("日粮ID"),
		field.String("name").Optional(),
		field.Int64("date").Comment("加工日期"),
		field.Int64("count").Comment("加工数量"),
		field.Int64("in").Comment("实际入库量"),
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

// Edges of the RationProcess.
func (RationProcess) Edges() []ent.Edge {
	return nil
}
