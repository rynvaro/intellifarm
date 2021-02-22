package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleGroup holds the schema definition for the CattleGroup entity.
type CattleGroup struct {
	ent.Schema
}

// Fields of the CattleGroup.
func (CattleGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("earNumber"),
		field.Int64("date").Comment("日期"),
		field.String("toShed").Comment("转到的栋舍"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the CattleGroup.
func (CattleGroup) Edges() []ent.Edge {
	return nil
}
