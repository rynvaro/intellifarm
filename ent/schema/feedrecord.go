package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
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
