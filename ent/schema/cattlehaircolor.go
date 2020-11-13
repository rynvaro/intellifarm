package schema

import "github.com/facebook/ent"

// CattleHairColor holds the schema definition for the CattleHairColor entity.
type CattleHairColor struct {
	ent.Schema
}

// Fields of the CattleHairColor.
func (CattleHairColor) Fields() []ent.Field {
	return nil
}

// Edges of the CattleHairColor.
func (CattleHairColor) Edges() []ent.Edge {
	return nil
}
