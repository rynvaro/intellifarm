package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CalveType holds the schema definition for the CalveType entity.
type CalveType struct {
	ent.Schema
}

// Fields of the CalveType.
func (CalveType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the CalveType.
func (CalveType) Edges() []ent.Edge {
	return nil
}
