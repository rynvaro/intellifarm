package schema

import "github.com/facebook/ent"

// Abortion holds the schema definition for the Abortion entity.
type Abortion struct {
	ent.Schema
}

// Fields of the Abortion.
func (Abortion) Fields() []ent.Field {
	return nil
}

// Edges of the Abortion.
func (Abortion) Edges() []ent.Edge {
	return nil
}
