package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("earNumber"),
		field.Int("eventTypeId").Optional(),
		field.String("eventTypeName"),
		field.Int("eventSubTypeId").Optional(),
		field.String("eventSubTypeName"),
		field.Int64("tenantId"),
		field.String("tenantName"),
		field.Int64("farmId"),
		field.String("farmName"),
		field.Int64("shedId").Optional(),
		field.String("shedName").Optional(),
		field.Int("times").Comment("胎次").Default(-1).Optional(),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
