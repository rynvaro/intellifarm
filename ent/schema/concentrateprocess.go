package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// ConcentrateProcess holds the schema definition for the ConcentrateProcess entity.
type ConcentrateProcess struct {
	ent.Schema
}

// Fields of the ConcentrateProcess.
func (ConcentrateProcess) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
		field.Int64("date"),
		field.Int64("count"),
		field.Int64("in").Comment("实际入库量"),
		field.Int64("inventory").Comment("库存"),
		field.String("userName"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the ConcentrateProcess.
func (ConcentrateProcess) Edges() []ent.Edge {
	return nil
}
