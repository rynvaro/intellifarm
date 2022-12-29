package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CattleGender holds the schema definition for the CattleGender entity.
type CattleGender struct {
	ent.Schema
}

// Fields of the CattleGender.
func (CattleGender) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the CattleGender.
func (CattleGender) Edges() []ent.Edge {
	return nil
}
