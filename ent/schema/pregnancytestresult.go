package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PregnancyTestResult holds the schema definition for the PregnancyTestResult entity.
type PregnancyTestResult struct {
	ent.Schema
}

// Fields of the PregnancyTestResult.
func (PregnancyTestResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the PregnancyTestResult.
func (PregnancyTestResult) Edges() []ent.Edge {
	return nil
}
