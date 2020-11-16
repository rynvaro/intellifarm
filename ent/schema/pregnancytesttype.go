package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// PregnancyTestType holds the schema definition for the PregnancyTestType entity.
type PregnancyTestType struct {
	ent.Schema
}

// Fields of the PregnancyTestType.
func (PregnancyTestType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the PregnancyTestType.
func (PregnancyTestType) Edges() []ent.Edge {
	return nil
}
