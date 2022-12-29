package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CattleOwner holds the schema definition for the CattleOwner entity.
type CattleOwner struct {
	ent.Schema
}

// Fields of the CattleOwner.
func (CattleOwner) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the CattleOwner.
func (CattleOwner) Edges() []ent.Edge {
	return nil
}
