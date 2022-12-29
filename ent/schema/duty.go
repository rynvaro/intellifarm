package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Duty holds the schema definition for the Duty entity.
type Duty struct {
	ent.Schema
}

// Fields of the Duty.
func (Duty) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("职责名称"),
	}
}

// Edges of the Duty.
func (Duty) Edges() []ent.Edge {
	return nil
}
