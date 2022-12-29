package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MaterialTest holds the schema definition for the MaterialTest entity.
type MaterialTest struct {
	ent.Schema
}

// Fields of the MaterialTest.
func (MaterialTest) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
		field.String("seqNumber").Comment("检测单号"),
		field.String("addSeqNumber").Comment("入库单号"),
		field.Int64("date").Comment("检测日期"),
		field.Int("type"),
		field.Int("category"),
		field.Int("materialCategory"),
		field.String("userName"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the MaterialTest.
func (MaterialTest) Edges() []ent.Edge {
	return nil
}
