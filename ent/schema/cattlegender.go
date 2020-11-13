package schema

import "github.com/facebook/ent"

// CattleGender holds the schema definition for the CattleGender entity.
type CattleGender struct {
	ent.Schema
}

// Fields of the CattleGender.
func (CattleGender) Fields() []ent.Field {
	return nil
}

// Edges of the CattleGender.
func (CattleGender) Edges() []ent.Edge {
	return nil
}
