package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// BreathRate holds the schema definition for the BreathRate entity.
type BreathRate struct {
	ent.Schema
}

// Fields of the BreathRate.
func (BreathRate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the BreathRate.
func (BreathRate) Edges() []ent.Edge {
	return nil
}
