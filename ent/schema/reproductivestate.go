package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// ReproductiveState holds the schema definition for the ReproductiveState entity.
type ReproductiveState struct {
	ent.Schema
}

// Fields of the ReproductiveState.
func (ReproductiveState) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the ReproductiveState.
func (ReproductiveState) Edges() []ent.Edge {
	return nil
}
