package schema

import "github.com/facebook/ent"

// Cattle holds the schema definition for the Cattle entity.
type Cattle struct {
	ent.Schema
}

// Fields of the Cattle.
func (Cattle) Fields() []ent.Field {
	return nil
}

// Edges of the Cattle.
func (Cattle) Edges() []ent.Edge {
	return nil
}
