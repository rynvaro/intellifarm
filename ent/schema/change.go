package schema

import "entgo.io/ent"

// Change holds the schema definition for the Change entity.
type Change struct {
	ent.Schema
}

// Fields of the Change.
func (Change) Fields() []ent.Field {
	return nil
}

// Edges of the Change.
func (Change) Edges() []ent.Edge {
	return nil
}
