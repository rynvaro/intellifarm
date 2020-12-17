package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Material holds the schema definition for the Material entity.
type Material struct {
	ent.Schema
}

// Fields of the Material.
func (Material) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
		field.String("seqNumber").Comment("入库单号"),
		field.Int64("date").Comment("入库日期"),
		field.Int("type"),
		field.Int("count"),
		field.Int("category"),
		field.Int("status"),
		field.String("userName"),
		field.Int64("payAt").Comment("结算日期"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Material.
func (Material) Edges() []ent.Edge {
	return nil
}
