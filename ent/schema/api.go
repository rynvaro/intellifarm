package schema

import "github.com/facebook/ent"

// API holds the schema definition for the API entity.
type API struct {
	ent.Schema
}

// Fields of the API.
func (API) Fields() []ent.Field {
	return nil
}

// Edges of the API.
func (API) Edges() []ent.Edge {
	return nil
}
