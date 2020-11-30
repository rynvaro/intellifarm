package schema

import "github.com/facebook/ent"

// FeedGroup holds the schema definition for the FeedGroup entity.
type FeedGroup struct {
	ent.Schema
}

// Fields of the FeedGroup.
func (FeedGroup) Fields() []ent.Field {
	return nil
}

// Edges of the FeedGroup.
func (FeedGroup) Edges() []ent.Edge {
	return nil
}
