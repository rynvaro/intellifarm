package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TreatmentResult holds the schema definition for the TreatmentResult entity.
type TreatmentResult struct {
	ent.Schema
}

// Fields of the TreatmentResult.
func (TreatmentResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the TreatmentResult.
func (TreatmentResult) Edges() []ent.Edge {
	return nil
}
