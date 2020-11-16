package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// EstrusType holds the schema definition for the EstrusType entity.
type EstrusType struct {
	ent.Schema
}

// Fields of the EstrusType.
func (EstrusType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the EstrusType.
func (EstrusType) Edges() []ent.Edge {
	return nil
}
