package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SemenFrozenType holds the schema definition for the SemenFrozenType entity.
type SemenFrozenType struct {
	ent.Schema
}

// Fields of the SemenFrozenType.
func (SemenFrozenType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
	}
}

// Edges of the SemenFrozenType.
func (SemenFrozenType) Edges() []ent.Edge {
	return nil
}
