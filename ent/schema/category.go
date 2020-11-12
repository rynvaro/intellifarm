package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("类别名称"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
