package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Whereabouts holds the schema definition for the Whereabouts entity.
type Whereabouts struct {
	ent.Schema
}

// Fields of the Whereabouts.
func (Whereabouts) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the Whereabouts.
func (Whereabouts) Edges() []ent.Edge {
	return nil
}
