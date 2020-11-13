package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("职务名称"),
	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return nil
}
