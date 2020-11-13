package schema

import "github.com/facebook/ent"

// CattleOwner holds the schema definition for the CattleOwner entity.
type CattleOwner struct {
	ent.Schema
}

// Fields of the CattleOwner.
func (CattleOwner) Fields() []ent.Field {
	return nil
}

// Edges of the CattleOwner.
func (CattleOwner) Edges() []ent.Edge {
	return nil
}
