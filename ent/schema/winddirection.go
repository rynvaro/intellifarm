package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// WindDirection holds the schema definition for the WindDirection entity.
type WindDirection struct {
	ent.Schema
}

// Fields of the WindDirection.
func (WindDirection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the WindDirection.
func (WindDirection) Edges() []ent.Edge {
	return nil
}
