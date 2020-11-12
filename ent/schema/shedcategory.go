package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// ShedCategory holds the schema definition for the ShedCategory entity.
type ShedCategory struct {
	ent.Schema
}

// Fields of the ShedCategory.
func (ShedCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("类别名称"),
	}
}

// Edges of the ShedCategory.
func (ShedCategory) Edges() []ent.Edge {
	return nil
}
