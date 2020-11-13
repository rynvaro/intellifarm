package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// BreedingType holds the schema definition for the BreedingType entity.
type BreedingType struct {
	ent.Schema
}

// Fields of the BreedingType.
func (BreedingType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the BreedingType.
func (BreedingType) Edges() []ent.Edge {
	return nil
}
