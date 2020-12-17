package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// PositionApi holds the schema definition for the PositionApi entity.
type PositionApi struct {
	ent.Schema
}

// Fields of the PositionApi.
func (PositionApi) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("positionId").Comment("职务ID"),
		field.String("apis").Comment("api ID 列表，逗号分割"),
	}
}

// Edges of the PositionApi.
func (PositionApi) Edges() []ent.Edge {
	return nil
}
