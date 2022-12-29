package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AbortionType holds the schema definition for the AbortionType entity.
type AbortionType struct {
	ent.Schema
}

// Fields of the AbortionType.
func (AbortionType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the AbortionType.
func (AbortionType) Edges() []ent.Edge {
	return nil
}
