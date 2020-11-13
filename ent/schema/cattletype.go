package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleType holds the schema definition for the CattleType entity.
type CattleType struct {
	ent.Schema
}

// Fields of the CattleType.
func (CattleType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the CattleType.
func (CattleType) Edges() []ent.Edge {
	return nil
}
