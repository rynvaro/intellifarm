package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleJoinedType holds the schema definition for the CattleJoinedType entity.
type CattleJoinedType struct {
	ent.Schema
}

// Fields of the CattleJoinedType.
func (CattleJoinedType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the CattleJoinedType.
func (CattleJoinedType) Edges() []ent.Edge {
	return nil
}
