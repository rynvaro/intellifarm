package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("earNumber"),
		field.String("eventType"),
		field.String("eventName"),
		field.Int64("tenantId"),
		field.String("tenantName"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
