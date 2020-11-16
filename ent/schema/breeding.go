package schema

import "github.com/facebook/ent"

// Breeding holds the schema definition for the Breeding entity.
type Breeding struct {
	ent.Schema
}

// Fields of the Breeding.
func (Breeding) Fields() []ent.Field {
	return nil
}

// Edges of the Breeding.
func (Breeding) Edges() []ent.Edge {
	return nil
}
