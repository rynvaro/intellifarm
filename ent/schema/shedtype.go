package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// ShedType holds the schema definition for the ShedType entity.
type ShedType struct {
	ent.Schema
}

// Fields of the ShedType.
func (ShedType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("类型名称"),
	}
}

// Edges of the ShedType.
func (ShedType) Edges() []ent.Edge {
	return nil
}
