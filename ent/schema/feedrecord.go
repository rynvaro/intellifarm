package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FeedRecord holds the schema definition for the FeedRecord entity.
type FeedRecord struct {
	ent.Schema
}

// Fields of the FeedRecord.
func (FeedRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("shedName"),
		field.Int64("date"),
		field.String("rationCode"),
		field.String("rationName"),
		field.Int64("rationAmount").Comment("饲喂量"),
		field.Int64("count").Comment("饲喂头数"),
		field.String("userName"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the FeedRecord.
func (FeedRecord) Edges() []ent.Edge {
	return nil
}
