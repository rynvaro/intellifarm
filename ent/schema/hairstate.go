package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// HairState holds the schema definition for the HairState entity.
type HairState struct {
	ent.Schema
}

// Fields of the HairState.
func (HairState) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the HairState.
func (HairState) Edges() []ent.Edge {
	return nil
}
