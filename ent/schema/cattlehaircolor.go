package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleHairColor holds the schema definition for the CattleHairColor entity.
type CattleHairColor struct {
	ent.Schema
}

// Fields of the CattleHairColor.
func (CattleHairColor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the CattleHairColor.
func (CattleHairColor) Edges() []ent.Edge {
	return nil
}
