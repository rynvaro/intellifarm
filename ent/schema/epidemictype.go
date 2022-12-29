package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EpidemicType holds the schema definition for the EpidemicType entity.
type EpidemicType struct {
	ent.Schema
}

// Fields of the EpidemicType.
func (EpidemicType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the EpidemicType.
func (EpidemicType) Edges() []ent.Edge {
	return nil
}
