package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PregnancyTestMethod holds the schema definition for the PregnancyTestMethod entity.
type PregnancyTestMethod struct {
	ent.Schema
}

// Fields of the PregnancyTestMethod.
func (PregnancyTestMethod) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the PregnancyTestMethod.
func (PregnancyTestMethod) Edges() []ent.Edge {
	return nil
}
