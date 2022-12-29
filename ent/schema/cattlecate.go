package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CattleCate holds the schema definition for the CattleCate entity.
type CattleCate struct {
	ent.Schema
}

// Fields of the CattleCate.
func (CattleCate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
		field.String("genderIds").NotEmpty().Comment("性别ID"),
	}
}

// Edges of the CattleCate.
func (CattleCate) Edges() []ent.Edge {
	return nil
}
