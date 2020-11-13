package schema

import "github.com/facebook/ent"

// BreedingType holds the schema definition for the BreedingType entity.
type BreedingType struct {
	ent.Schema
}

// Fields of the BreedingType.
func (BreedingType) Fields() []ent.Field {
	return nil
}

// Edges of the BreedingType.
func (BreedingType) Edges() []ent.Edge {
	return nil
}
