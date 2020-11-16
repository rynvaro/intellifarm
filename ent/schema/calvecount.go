package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CalveCount holds the schema definition for the CalveCount entity.
type CalveCount struct {
	ent.Schema
}

// Fields of the CalveCount.
func (CalveCount) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
	}
}

// Edges of the CalveCount.
func (CalveCount) Edges() []ent.Edge {
	return nil
}
