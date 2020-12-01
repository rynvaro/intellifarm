package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// API holds the schema definition for the API entity.
type API struct {
	ent.Schema
}

// Fields of the API.
func (API) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("path").Unique(),
		field.Int("level"),
		field.String("hash"),
		field.String("redirect").Optional(),
		field.String("component"),
		field.Bool("isSub"),
		field.Bool("hasSub"),
		field.Bool("single"),
		field.Int64("parentId"),
		field.String("tenantId"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the API.
func (API) Edges() []ent.Edge {
	return nil
}
