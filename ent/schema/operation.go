package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Operation holds the schema definition for the Operation entity.
type Operation struct {
	ent.Schema
}

// Fields of the Operation.
func (Operation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("userId"),
		field.String("userName"),
		field.Int64("tenantId"),
		field.String("tenantName"),
		field.Int64("farmId"),
		field.String("farmName"),
		field.String("api"),
		field.String("ip"),
		field.String("method"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Operation.
func (Operation) Edges() []ent.Edge {
	return nil
}
