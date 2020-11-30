package schema

import "github.com/facebook/ent"

// CattleGroup holds the schema definition for the CattleGroup entity.
type CattleGroup struct {
	ent.Schema
}

// Fields of the CattleGroup.
func (CattleGroup) Fields() []ent.Field {
	return nil
}

// Edges of the CattleGroup.
func (CattleGroup) Edges() []ent.Edge {
	return nil
}
