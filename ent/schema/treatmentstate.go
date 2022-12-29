package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TreatmentState holds the schema definition for the TreatmentState entity.
type TreatmentState struct {
	ent.Schema
}

// Fields of the TreatmentState.
func (TreatmentState) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the TreatmentState.
func (TreatmentState) Edges() []ent.Edge {
	return nil
}
