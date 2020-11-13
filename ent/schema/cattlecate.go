package schema

import "github.com/facebook/ent"

// CattleCate holds the schema definition for the CattleCate entity.
type CattleCate struct {
	ent.Schema
}

// Fields of the CattleCate.
func (CattleCate) Fields() []ent.Field {
	return nil
}

// Edges of the CattleCate.
func (CattleCate) Edges() []ent.Edge {
	return nil
}
