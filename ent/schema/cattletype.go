package schema

import "github.com/facebook/ent"

// CattleType holds the schema definition for the CattleType entity.
type CattleType struct {
	ent.Schema
}

// Fields of the CattleType.
func (CattleType) Fields() []ent.Field {
	return nil
}

// Edges of the CattleType.
func (CattleType) Edges() []ent.Edge {
	return nil
}
